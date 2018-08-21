package graph

import (
	options "github.com/jiro4989/plt/internal/options"
	"gonum.org/v1/plot/plotter"
)

type Graph struct {
	Title string
	Data  plotter.XYs
}

func GetGraphs(lines []string, opts options.Options) ([]Graph, error) {
	return nil, nil
}
