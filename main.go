package main

import (
	"log"

	"github.com/joho/godotenv"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No.env file found.")
	}
	log.Println(".env file loaded.")
}

func main() {
	Init()
}
