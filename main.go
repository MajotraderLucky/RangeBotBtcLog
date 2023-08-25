package main

import (
	"log"

	"github.com/MajotraderLucky/Utils/logger"
	"github.com/joho/godotenv"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No.env file found.")
	}
	log.Println(".env file loaded.")
}

func main() {
	logger := logger.Logger{}
	err := logger.CreateLogsDir()
	if err != nil {
		log.Fatal(err)
	}
	err = logger.OpenLogFile()
	if err != nil {
		log.Fatal(err)
	}
	logger.SetLogger()
	logger.LogLine()

	Init()
}
