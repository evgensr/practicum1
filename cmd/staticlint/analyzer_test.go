package main

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

// TestMyAnalyzer the ErrCheckAnalyzer analyzer under test from folder analysistest.TestData()
func TestMyAnalyzer(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), OsExitCheck, "./...")
}
