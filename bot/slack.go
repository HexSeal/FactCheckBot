package bot

import (
	"os"

	"fmt"
	"log"
	"strings"

	"github.com/HexSeal/FactCheckBot/factcheck"
	"github.com/slack-go/slack"
)

/*
	CreateSlackClient initiates the slack socket connection and real-time messaging(RTM) client library,
	and returns the client
*/
func CreateSlackClient(apiKey string) *slack.RTM {
	api := slack.New(
		apiKey,
		slack.OptionDebug(true),
		slack.OptionLog(log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)),
	)

	rtm := api.NewRTM()
	go rtm.ManageConnection()
	return rtm
}

/*
	RespondToEvents waits for messages on the Slack client's incomingEvents channel,
	and sends a response when it detects the bot has been tagged in a message with @<botTag>.
*/
func RespondToEvents(slackClient *slack.RTM) {
	for msg := range slackClient.IncomingEvents {
		// Log all events
		fmt.Println("Event Received: ", msg.Type)
		// Switch on the incoming event type
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			// The bot's prefix (@topofreddit)
			botTagString := fmt.Sprintf("<@%s> ", slackClient.GetInfo().User.ID)
			if !strings.Contains(ev.Msg.Text, botTagString) {
				continue
			}
			// Get rid of the prefix
			message := strings.Replace(ev.Msg.Text, botTagString, "", -1)
			splitMessage := strings.Fields(message)

			// If they do not specify a command just @ the bot, send them the help menu.
			if message == "" {
				sendHelp(slackClient, ev.Channel)
			}

			// Basic command handler
			switch strings.ToLower(splitMessage[0]) {
			case "help":
				sendHelp(slackClient, ev.Channel)
			case "echo":
				echoMessage(slackClient, strings.Join(splitMessage[:], " "), ev.Channel)
			case "check":
				checkQuery(slackClient, strings.Join(splitMessage[:1], " "), ev.Channel)
			}
		case *slack.PresenceChangeEvent:
			fmt.Printf("Presence Change: %v\n", ev)

		case *slack.LatencyReport:
			fmt.Printf("Current latency: %v\n", ev.Value)

		case *slack.DesktopNotificationEvent:
			fmt.Printf("Desktop Notification: %v\n", ev)

		case *slack.RTMError:
			fmt.Printf("Error: %s\n", ev.Error())

		case *slack.InvalidAuthEvent:
			fmt.Printf("Invalid credentials")
			return
		default:
		}
	}
}

const helpMessage = "No"

// sendHelp is a working help message, for reference.
func sendHelp(slackClient *slack.RTM, slackChannel string) {
	slackClient.SendMessage(slackClient.NewOutgoingMessage(helpMessage, slackChannel))
}

// echoMessage will just echo anything after the echo keyword.
func echoMessage(slackClient *slack.RTM, message, slackChannel string) {
	splitMessage := strings.Fields(strings.ToLower(message))

	slackClient.SendMessage(slackClient.NewOutgoingMessage(strings.Join(splitMessage[1:], " "), slackChannel))
}

func checkQuery(slackClient *slack.RTM, message, slackChannel string) {
	formattedQuery := strings.ToLower(message)
	factcheck.ChromeCheck(formattedQuery)

	slackClient.SendMessage(slackClient.NewOutgoingMessage(answer, slackChannel))
}
