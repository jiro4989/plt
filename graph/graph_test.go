package graph

import (
	"testing"

	options "github.com/jiro4989/plt/internal/options"
	"github.com/stretchr/testify/assert"
	"gonum.org/v1/plot/plotter"
)

type TestReshapeData struct {
	in  [][]string
	out [][]string
}

func TestReshape(t *testing.T) {
	testdatas := []TestReshapeData{
		TestReshapeData{
			in: [][]string{
				{"1", "2", "3"},
				{"4", "5", "6"},
				{"7", "8", "9"},
			},
			out: [][]string{
				{"1", "4", "7"},
				{"2", "5", "8"},
				{"3", "6", "9"},
			},
		},
		TestReshapeData{
			in: [][]string{
				{"1", "2"},
				{"3", "4"},
			},
			out: [][]string{
				{"1", "3"},
				{"2", "4"},
			},
		},
		TestReshapeData{
			in: [][]string{
				{"1", "2"},
			},
			out: [][]string{
				{"1"},
				{"2"},
			},
		},
		TestReshapeData{
			in:  [][]string{{"1"}},
			out: [][]string{{"1"}},
		},
		TestReshapeData{
			in:  [][]string{},
			out: [][]string{},
		},
	}
	for _, v := range testdatas {
		assert.Equal(t, v.out, reshape(v.in))
	}
}

type TestSplitLinesData struct {
	lines []string
	sep   string
	out   [][]string
}

func TestSplitLines(t *testing.T) {
	testdatas := []TestSplitLinesData{
		TestSplitLinesData{
			lines: []string{"value", "1"},
			sep:   "\t",
			out:   [][]string{[]string{"value", "1"}},
		},
		TestSplitLinesData{
			lines: []string{"value1\tvalue2", "1\t2"},
			sep:   "\t",
			out: [][]string{
				[]string{"value1", "1"},
				[]string{"value2", "2"},
			},
		},
		TestSplitLinesData{
			lines: []string{"value1,value2", "1,2"},
			sep:   ",",
			out: [][]string{
				[]string{"value1", "1"},
				[]string{"value2", "2"},
			},
		},
		TestSplitLinesData{
			lines: []string{"\t  value1   ,   value2  \t ", "1 \t  ,\t    2"},
			sep:   ",",
			out: [][]string{
				[]string{"value1", "1"},
				[]string{"value2", "2"},
			},
		},
		TestSplitLinesData{
			lines: []string{"value1,value2", "1,2"},
			sep:   "",
			out: [][]string{
				[]string{"value1,value2", "1,2"},
			},
		},
	}
	for _, v := range testdatas {
		assert.Equal(t, v.out, splitLines(v.lines, v.sep))
	}
}

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
		TestGetGraphsData{
			in: []string{
				"value1\tvalue2",
				"1\t3",
				"3\t5",
				"5\t7",
				"7\t9",
				"9\t11",
			},
			out: []Graph{
				Graph{
					Title: "value1",
					Data:  f(1, 3, 5, 7, 9),
				},
				Graph{
					Title: "value2",
					Data:  f(3, 5, 7, 9, 11),
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
