package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)



func checkEnv(envWorks string) {
	fmt.Println(envWorks)
}

func main() {
	fmt.Println("Start:")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	envWorks := os.Getenv("envWorks")
	checkEnv(envWorks)
}