package detectfmt

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "detectfmt",
	Doc:      "Check if there is unnecessary fmt",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	// pass.ResultOf[inspect.Analyzer] will be set if we've added inspect.Analyzer to Requires.
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{ // filter needed nodes: visit only them
		(*ast.SelectorExpr)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		selectorExpr, ok := node.(*ast.SelectorExpr)
		if !ok {
			return
		}

		xIdent, ok := selectorExpr.X.(*ast.Ident)
		if !ok {
			return
		}

		// Only process if expression is from "fmt" package
		if xIdent.Name != "fmt" {
			return
		}

		if report := detectPrintCode(selectorExpr.Sel.Name); report != "" {
			pass.Reportf(selectorExpr.Pos(), report)
		}
	})

	return nil, nil
}

func detectPrintCode(selName string) string {
	result := ""
	if strings.HasPrefix(selName, "Print") {
		result = "Print code detected!"
		return result
	}
	return result
}
