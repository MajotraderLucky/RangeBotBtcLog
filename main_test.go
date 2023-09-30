package main

import (
	"testing"

	"github.com/MajotraderLucky/MarketRepository/initlog"
	"github.com/stretchr/testify/assert"
)

func TestCheckFilesExist(t *testing.T) {
	// Perform the CheckFilesExist() function test
	result := initlog.CheckFilesExist()

	// Check that the function returns true
	assert.True(t, result)
}
