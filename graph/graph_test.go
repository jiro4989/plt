package graph

import (
	"testing"

	options "github.com/jiro4989/plt/internal/options"
	"github.com/stretchr/testify/assert"
	"gonum.org/v1/plot/plotter"
)

type TestGetGraphsData struct {
	in  []string
	out []Graph
	err error
}

func TestGetGraphs(t *testing.T) {
	f := func(n ...int) plotter.XYs {
		p := make(plotter.XYs, len(n))
		for i, v := range n {
			p[i].X = float64(i)
			p[i].Y = float64(v)
		}
		return p
	}
	opts := options.Options{
		Separator:   "\t",
		Direction:   "vertical",
		NoHeader:    false,
		NoRowHeader: false,
	}
	datas := []TestGetGraphsData{
		TestGetGraphsData{
			in: []string{
				"value",
				"1",
				"3",
				"5",
				"7",
				"9",
			},
			out: []Graph{
				Graph{
					Title: "value",
					Data:  f(1, 3, 5, 7, 9),
				},
			},
			err: nil,
		},
	}
	for _, v := range datas {
		actual, err := GetGraphs(v.in, opts)
		assert.Equal(t, v.out, actual)
		assert.Equal(t, v.err, err)
	}
}
