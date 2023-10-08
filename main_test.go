package main

import (
	"context"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/MajotraderLucky/MarketRepository/initlog"
	"github.com/MajotraderLucky/MarketRepository/klinesdata"
	"github.com/MajotraderLucky/MarketRepository/positionlog"
	"github.com/MajotraderLucky/Utils/logger"
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

// Testing WritePositionsToFile with empty data and verifying file creation behavior in case of an error.
func TestWritePositionsToFile(t *testing.T) {
	data := positionlog.AutoGeneratedPos{}

	err := positionlog.WritePositionsToFile(data)
	if err == nil {
		t.Fatal("Expected an error when calling WritePositionsToFile with empty data, got no error")
	}

	data = positionlog.AutoGeneratedPos{
		Positions: []struct {
			Isolated               bool   `json:"isolated"`
			Leverage               string `json:"leverage"`
			InitialMargin          string `json:"initialMargin"`
			MaintMargin            string `json:"maintMargin"`
			OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
			PositionInitialMargin  string `json:"positionInitialMargin"`
			Symbol                 string `json:"symbol"`
			UnrealizedProfit       string `json:"unrealizedProfit"`
			EntryPrice             string `json:"entryPrice"`
			MaxNotional            string `json:"maxNotional"`
			PositionSide           string `json:"positionSide"`
			PositionAmt            string `json:"positionAmt"`
			Notional               string `json:"notional"`
			IsolatedWallet         string `json:"isolatedWallet"`
			UpdateTime             int64  `json:"updateTime"`
		}{
			{
				Symbol: "ETHUSDT",
			},
		},
	}

	err = positionlog.WritePositionsToFile(data)
	if err == nil {
		t.Fatal("Expected an error when calling WritePositionsToFile with data without BTCUSDT position, got no error")
	}
}

// Create a mock for futuresClient with a predictable response
type MockFuturesClient struct {
	mock.Mock
}

func (m *MockFuturesClient) NewKlinesService() klinesdata.KlinesService {
	args := m.Called()
	return args.Get(0).(klinesdata.KlinesService)
}

type MockKlinesService struct {
	mock.Mock
}

func (m *MockKlinesService) Symbol(symbol string) klinesdata.KlinesService {
	m.Called(symbol)
	return m
}

func (m *MockKlinesService) Interval(interval string) klinesdata.KlinesService {
	m.Called(interval)
	return m
}

func (m *MockKlinesService) Do(ctx context.Context, opts ...klinesdata.RequestOption) ([]*klinesdata.Kline, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*klinesdata.Kline), args.Error(1)
}

func TestFindMinMaxInfo2(t *testing.T) {
	mockClient := new(MockFuturesClient)

	expectedKlines := []*klinesdata.Kline{
		{
			High: "12000",
			Low:  "9000",
		},
	}

	mockService := new(MockKlinesService)

	// Mock function calls
	mockClient.On("NewKlinesService").Return(mockService)
	mockService.On("Symbol", "BTCUSDT").Return(mockService)
	mockService.On("Interval", "15m").Return(mockService)
	mockService.On("Do", mock.Anything).Return(expectedKlines, nil)

	max, min, err := klinesdata.FindMinMaxInfoTest(mockClient)

	// Verification assertions
	assert.NoError(t, err)
	assert.Equal(t, float64(12000), max)
	assert.Equal(t, float64(9000), min)

	// Check that the functions were called
	mockClient.AssertCalled(t, "NewKlinesService")
	mockService.AssertCalled(t, "Symbol", "BTCUSDT")
	mockService.AssertCalled(t, "Interval", "15m")
	mockService.AssertCalled(t, "Do", mock.Anything)
}
