package bot

import (
	"os"
	
	"github.com/slack-go/slack"
	"log"
)

/*
	CreateSlackClient initiates the slack socket connection and real-time messaging(RTM) client library,
	and returns the client
*/
func CreateSlackClient(apiKey string) *slack.RTM {
	token := os.Getenv("BOT_OAUTH_ACCESS_TOKEN")
	api := slack.New(
		token,
		slack.OptionDebug(true),
		slack.OptionLog(log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)),
	)

	rtm := api.NewRTM()
	go rtm.ManageConnection()
	return rtm
}
