package main

import (
	"github.com/alextanhongpin/nosetter"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(nosetter.Analyzer) }
