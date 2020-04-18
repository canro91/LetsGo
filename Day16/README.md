# Notes

## How-to setup Slack

1. Create a workspace [here](https://slack.com/create)
2. Once the workspace is created, create a new bot [here](https://my.slack.com/services/new/bot). Name it `dollybot`
3. Grab API Token
4. Post a message into a channel to test. From Postman:
```
POST https://slack.com/api/chat.postMessage 
Authorization: Bearer <xoxb-API-Token>
Content-type: application/json;charset=utf-8
{"channel":"#general","text":"Hello, Slack!"}
```
5. `$ export SLACK_API_TOKEN=<xoxb-API-Token>`
5. Run `go run main.go`
6. Invite `@dollybot` into the channel
7. In Slack, do `@dollybot ping`
8. Voilà

### See:

* [How to build a basic slackbot: a beginner’s guide](https://www.freecodecamp.org/news/how-to-build-a-basic-slackbot-a-beginners-guide-6b40507db5c5/)
* [Joke teller bot](https://www.youtube.com/watch?v=nyyXTIL3Hkw)
* [Slacker](https://github.com/shomali11/slacker)
* [GoQuery](https://github.com/PuerkitoBio/goquery)
