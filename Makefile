.PHONY: run
run:
	go run main.go

.PHONY: proxy
proxy:
	./ngrok http 3000