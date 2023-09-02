package main

import (
	"log"

	"github.com/MajotraderLucky/MarketRepository/initlog"
	"github.com/MajotraderLucky/MarketRepository/positionlog"
	"github.com/MajotraderLucky/Utils/logger"
)

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

	log.Println("Btc bot started...")

	initlog.Init()

	positionlog.Hello()
}
