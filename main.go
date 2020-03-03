package main

import (
	"net/http"
	"os"
	// "log"

	"github.com/HexSeal/FactCheckBot/bot"
	"github.com/HexSeal/FactCheckBot/factCheck"
	_ "github.com/joho/godotenv/autoload"
)

// slackIt initializes the slackbot
func slackIt() {
	botToken := os.Getenv("BOT_OAUTH_ACCESS_TOKEN")
	slackClient := bot.CreateSlackClient(botToken)
	bot.RespondToEvents(slackClient)
}

// func checkEnv(envWorks string) {
// 	fmt.Println(envWorks)
// }

func main() {
	port := ":" + os.Getenv("PORT")
	go http.ListenAndServe(port, nil)
	// slackIt()
	factCheck.SnopesCheck()

	// Env check
	// fmt.Println("Start:")
	// envWorks := os.Getenv("envWorks")
	// checkEnv(envWorks)
}
