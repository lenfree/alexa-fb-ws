package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	fb "github.com/huandu/facebook"
	"github.com/joho/godotenv"
	alexa "github.com/mikeflynn/go-alexa/skillserver"
)

// Messages contains my inbox messages
type Messages struct {
	ID    string `json:"id"`
	Inbox struct {
		Data []struct {
			ID string `json:"id"`
			To struct {
				Data []struct {
					ID   string `json:"id"`
					Name string `json:"name"`
				} `json:"data"`
			} `json:"to"`
			Unread      int    `json:"unread"`
			Unseen      int    `json:"unseen"`
			UpdatedTime string `json:"updated_time"`
		} `json:"data"`
		Paging struct {
			Next     string `json:"next"`
			Previous string `json:"previous"`
		} `json:"paging"`
		Summary struct {
			UnreadCount int    `json:"unread_count"`
			UnseenCount int    `json:"unseen_count"`
			UpdatedTime string `json:"updated_time"`
		} `json:"summary"`
	} `json:"inbox"`
}

var (
	APP_ID       string
	APP_SECRET   string
	API_TOKEN    string
	ALEXA_APP_ID string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	APP_ID = os.Getenv("APP_ID")
	APP_SECRET = os.Getenv("APP_SECRET")
	API_TOKEN = os.Getenv("API_TOKEN")
	ALEXA_APP_ID = os.Getenv("ALEXA_APP_ID")
	log.Printf("APPID %s\n", APP_ID)
	log.Printf("APP_SECRET %s\n", APP_SECRET)
	log.Printf("Token %s\n", API_TOKEN)
	log.Printf("ALEXAID: %s\n", ALEXA_APP_ID)
}

func main() {
	Applications := map[string]interface{}{
		"/echo/myfacebook": alexa.EchoApplication{ // Route
			AppID:    ALEXA_APP_ID, // Echo App ID from Amazon Dashboard
			OnIntent: EchoIntentHandler,
			OnLaunch: EchoIntentHandler,
		},
	}
	alexa.Run(Applications, "3000")
}

func EchoIntentHandler(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	s := new()
	total, unreadMsgFrom := unreadMsg(s)
	var speechText string
	if total > 0 {
		speechText = "You have " + strconv.Itoa(total) + " of unread messages. From " + strings.Join(unreadMsgFrom, "... ")
	} else {
		speechText = "You have 0 unread message."
	}
	echoResp.OutputSpeech(speechText).Card("Facebook", "Unread Messages.")
}

func new() *fb.Session {
	app := fb.New(APP_ID, APP_SECRET)
	s := app.Session(API_TOKEN)
	s.Version = "v2.3"

	return s
}

func unreadMsg(s *fb.Session) (int, []string) {
	res, err := s.Get("/me", fb.Params{
		"fields": "inbox{from,message,subject,updated_time,to,unread,unseen,id}",
	})
	if err != nil {
		log.Printf("error: %s\n", err.Error())
	}

	var m Messages
	res.Decode(&m)
	var total int
	var unReadMsgs []string
	log.Printf("read: %+#v", unReadMsgs)
	for _, msg := range m.Inbox.Data {
		if msg.Unread > 0 {
			from := msg.To.Data[0].Name
			unReadMsgs = append(unReadMsgs, from)
			total++

		}
	}
	log.Printf("You have %s unread messages\n", strconv.Itoa(total))
	log.Println(strings.Join(unReadMsgs, "... "))

	return total, unReadMsgs
}
