package main

import (
	"github.com/fawwazaf/detectfmt"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(detectfmt.Analyzer)
}
