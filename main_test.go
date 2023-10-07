package main

import (
	"log"
	"os"
	"strings"
	"testing"

	"github.com/MajotraderLucky/MarketRepository/initlog"
	"github.com/MajotraderLucky/Utils/logger"
	"github.com/adshao/go-binance/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCheckFilesExist(t *testing.T) {
	// Perform the CheckFilesExist() function test
	result := initlog.CheckFilesExist()

	// Check that the function returns true
	assert.True(t, result)
}

func TestCreateLogsDir(t *testing.T) {
	logger := logger.Logger{}
	err := logger.CreateLogsDir()
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// Check that the "logs" directory was created
	_, err = os.Stat("logs")
	if os.IsNotExist(err) {
		t.Error("Expected 'logs' directory to be created, but it doesn't exist")
	}
}

func TestOpenLogFile(t *testing.T) {
	logger := logger.Logger{}
	err := logger.OpenLogFile()
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// Check that the log file was created
	_, err = os.Stat("logs/log.txt")
	if os.IsNotExist(err) {
		t.Error("Expected 'logs/log.txt' file to be created, but it doesn't exist")
	}
}

func TestSetLogger(t *testing.T) {
	// Create a temporary file for testing
	file, err := os.Create("test.log")
	if err != nil {
		t.Fatalf("Failed to create log file: %v", err)
	}
	defer file.Close()

	logger := logger.Logger{}
	logger.SetLogger()

	// Redirect log output to the specified file
	log.SetOutput(file)

	// Check that the log output is indeed redirected to the specified file
	log.Println("This is a test log message")

	// Reed the contents of the log file
	contents, err := os.ReadFile("test.log")
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	// Check that the log message is present in the file
	if !strings.Contains(string(contents), "This is a test log message") {
		t.Errorf("Expected log message not found in log file")
	}
}
