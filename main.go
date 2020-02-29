package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	// "log"
	"os"
)



func checkEnv(envWorks string) {
	fmt.Println(envWorks)
}

func main() {
	fmt.Println("Start:")
	envWorks := os.Getenv("envWorks")
	checkEnv(envWorks)
}