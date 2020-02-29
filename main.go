package main

import (
	"net/http"
	"os"
	// "log"
	
	_ "github.com/joho/godotenv/autoload"
	"github.com/HexSeal/FactCheckBot/bot"
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
	slackIt()

	// fmt.Println("Start:")
	// envWorks := os.Getenv("envWorks")
	// checkEnv(envWorks)
}