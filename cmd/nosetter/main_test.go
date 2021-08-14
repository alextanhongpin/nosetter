package main_test

import (
	"testing"

	"github.com/alextanhongpin/nosetter"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestBasic(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, nosetter.Analyzer)
}
