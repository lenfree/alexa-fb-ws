alexa-fb-skill
==============

A web service for Alexa Facebook skill integration.

Usage:
------

```
$ cp .env.example .env
$ curl -o ngrok.zip https://bin.equinox.io/c/4VmDzA7iaHb/ngrok-stable-darwin-amd64.zip \
    && unzip ngrok.zip
$ go run main.go
$ ./ngrok http 3000
```

Ngrok is a secure tunnel to localhost. This can be replaced by other reverse proxy
as long as it can offload SSL cert as required for Alexa integration. When using
ngrok, it would return a http and https endpoint and you use this https endpoint to
configure Alexa's endpoint and SSL certification with "My development endpoint is
a sub-domain of a domain that has a wildcard certificate from a certificate authority".


Alexa skills:
-------------

# Alexa, ask myfacebook if I have any new messages?
# Alexa, ask myfacebook do I have any unread messages?