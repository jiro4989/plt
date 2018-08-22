package graph

import (
	"fmt"
	"strconv"

	options "github.com/jiro4989/plt/internal/options"
	"gonum.org/v1/plot/plotter"
)

type Graph struct {
	Title string
	Data  plotter.XYs
}

func GetGraphs(lines []string, opts options.Options) ([]Graph, error) {
	var title string
	p := make(plotter.XYs, len(lines)-1)
	for i, v := range lines {
		if i == 0 {
			title = v
		} else {
			n, err := strconv.ParseFloat(v, 64)
			if err != nil {
				fmt.Println("error")
				continue
			}
			p[i-1].X = float64(i - 1)
			p[i-1].Y = n
		}
	}

	g := []Graph{
		Graph{
			Title: title,
			Data:  p,
		},
	}
	return g, nil
}
