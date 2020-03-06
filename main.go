package main

import (
	"net/http"
	"os"

	// "log"

	"github.com/HexSeal/FactCheckBot/bot"
	"github.com/HexSeal/FactCheckBot/factcheck"
	"github.com/HexSeal/FactCheckBot/selenium"
	_ "github.com/joho/godotenv/autoload"
	// "github.com/tebeka/selenium"
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

// runCheck runs the Selenium Webdriver
func runCheck() {
	factcheck.SnopesCheck()
}

func main() {
	port := ":" + os.Getenv("PORT")
	go http.ListenAndServe(port, nil)
	// slackIt()
	// runCheck()
	selenium.ChromeTest()

	// Env check
	// fmt.Println("Start:")
	// envWorks := os.Getenv("envWorks")
	// checkEnv(envWorks)
}
