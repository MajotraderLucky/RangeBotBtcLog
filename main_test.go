package main

import (
	"os"
	"testing"

	"github.com/MajotraderLucky/MarketRepository/initlog"
	"github.com/MajotraderLucky/Utils/logger"
	"github.com/stretchr/testify/assert"
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
