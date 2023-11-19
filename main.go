package main

import (
	"log"

	"github.com/MajotraderLucky/MarketRepository/initlog"
	"github.com/MajotraderLucky/MarketRepository/klinesdata"
	"github.com/MajotraderLucky/MarketRepository/orderinfolog"
	"github.com/MajotraderLucky/MarketRepository/positionlog"
	"github.com/MajotraderLucky/MarketRepository/tradinglog"
	"github.com/MajotraderLucky/MarketRepository/transactions"
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

	// If all files are found, start the program
	if initlog.CheckFilesExist() {
		log.Println("----------All files found, starting program----------")
	}

	log.Println("Btc bot started...")

	initlog.Init()

	autoGeneratedPos := positionlog.GetAutoGeneratedPos()
	hasOpenPos := positionlog.IsOpenPositions(autoGeneratedPos)
	log.Println("Has open positions: ", hasOpenPos)
	positionlog.GetOpenPositionVolume(autoGeneratedPos)

	logger.LogLine()

	positionlog.WritePositionsToFile(autoGeneratedPos)

	logger.LogLine()

	klinesdata.GetDebthData()
	klinesdata.KlinesInfo()
	klinesdata.FindMinMaxInfo()
	klinesdata.GetFibonacciLevelsReturns()
	klinesdata.FindPriceCorridor()
	klinesdata.IsCorridorHigher(5)

	orderinfolog.Hello()
	orderinfolog.GetOpenOrdersInfoJson()
	logger.LogLine()

	tradinglog.GetFiboLevelStartTrade()

	logger.LogLine()
	transactions.Hello()

	floLevels, err := klinesdata.GetFibonacciLevelsReturns()
	if err != nil {
		log.Fatalf("Error getting Fibonacci levels: %v", err)
	}

	intLevels, err := klinesdata.ConvertFibonacciLevelsToInts(floLevels)
	if err != nil {
		log.Fatalf("Error converting Fibonacci levels to integers: %v", err)
	}

	log.Println("Fibonacci levels to ints", intLevels)

	strLevels, err := klinesdata.ConvertIntsToStrings(intLevels)
	if err != nil {
		log.Fatalf("Error converting ints to strings: %v", err)
	}
	log.Println("Ints to strings", strLevels)

	if tradinglog.IsStartTradeLevel382Met() {
		transactions.CreateLimitOrder("0.003", strLevels[1])
	}

	logger.CleanLogCountLines(200)
}
