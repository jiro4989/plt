package graph

import (
	"fmt"
	"strconv"
	"strings"

	options "github.com/jiro4989/plt/internal/options"
	"gonum.org/v1/plot/plotter"
)

type Graph struct {
	Title string
	Data  plotter.XYs
}

// reshape は行列を入れ替えます。
// ex:
//   in:
//     1,2,3
//     4,5,6
//   out:
//     1,4
//     2,5
//     3,6
func reshape(s [][]string) [][]string {
	if len(s) <= 0 {
		return [][]string{}
	}

	rl := len(s)
	cl := len(s[0])
	matrix := make([][]string, cl)
	for i := range matrix {
		matrix[i] = make([]string, rl)
	}

	for i, r := range s {
		for j, c := range r {
			matrix[j][i] = c
		}
	}
	return matrix
}

// splitLines は配列文字列を指定の文字列で更に分割した二次元配列を返す。
// セパレータがから文字列の場合はそのまま文字列を二次元配列で返す。
func splitLines(lines []string, sep string) [][]string {
	if sep == "" {
		return [][]string{lines}
	}

	splited := make([][]string, len(lines))
	for i, line := range lines {
		ss := strings.Split(line, sep)
		ns := make([]string, len(ss))
		for i, v := range ss {
			ns[i] = strings.TrimSpace(v)
		}
		splited[i] = ns
	}
	return reshape(splited)
}

func GetGraphs(lines []string, opts options.Options) ([]Graph, error) {
	matrix := splitLines(lines, opts.Separator)
	gs := make([]Graph, len(matrix))
	for x, ls := range matrix {
		var title string
		p := make(plotter.XYs, len(ls)-1)
		for i, v := range ls {
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
		g := Graph{
			Title: title,
			Data:  p,
		}
		gs[x] = g
	}

	return gs, nil
}
