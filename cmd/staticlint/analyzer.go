package main

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

// OsExitCheck - the variable of type &analysis.Analyzer for creating custom
// analyzer.
var OsExitCheck = &analysis.Analyzer{
	Name: "noExit",
	Doc:  "check if os.Exit don't use",
	Run:  run,
}

// parseMainFunction - checks the fun for the existence of Exit
func parseMainFunction(pass *analysis.Pass, fun *ast.FuncDecl) {
	ast.Inspect(fun, func(n ast.Node) bool {
		if c, ok := n.(*ast.CallExpr); ok {
			if s, ok := c.Fun.(*ast.SelectorExpr); ok {
				if fmt.Sprint(s.X) == "os" && s.Sel.Name == "Exit" {
					pass.Reportf(s.X.Pos(), "os.Exit in main()")
				}
			}
		}
		return true
	})
}

// run contains the main code of checking. Accepts the pointer to struct
// analysis.Pass in parameter.
// The function tracks usage os.Exit.
func run(pass *analysis.Pass) (interface{}, error) {
	if pass.Pkg.Name() != "main" {
		return nil, nil
	}
	for _, file := range pass.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			if f, ok := node.(*ast.FuncDecl); ok {
				if f.Name.Name == "main" {
					parseMainFunction(pass, f)
				}
			}
			return true
		})
	}
	return nil, nil
}
