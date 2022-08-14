// staticlint is a tool for static analysis of Go programs.
// Usage: staticlint [-flag] [package]
// Run 'staticlint help' for more detail,
// or 'staticlint help name' for details and flags of a specific analyzer.
// for testing all './staticlint ./...'
package main

import (
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/shadow"
	"golang.org/x/tools/go/analysis/passes/shift" // импортируем дополнительный анализатор
	"golang.org/x/tools/go/analysis/passes/structtag"
	"honnef.co/go/tools/staticcheck"
)

func main() {

	myChecks := []*analysis.Analyzer{
		OsExitCheck,
		printf.Analyzer,
		shadow.Analyzer,
		structtag.Analyzer,
		shift.Analyzer,
	}

	additionalChecks := map[string]bool{
		"S1003":  true, // Replace call to strings.Index with strings.Contains
		"ST1008": true, // A function’s error value should be its last return value
		"ST1017": true, // Don’t use Yoda conditions
	}

	for _, v := range staticcheck.Analyzers {
		_, ok := additionalChecks[v.Analyzer.Name]
		if strings.Contains(v.Analyzer.Name, "SA") || ok {
			myChecks = append(myChecks, v.Analyzer)
		}
	}

	multichecker.Main(
		myChecks...,
	)
}
