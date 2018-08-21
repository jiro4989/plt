package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"

	flags "github.com/jessevdk/go-flags"
	"github.com/jiro4989/plt/graph"
	options "github.com/jiro4989/plt/options"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// アプリのバージョン情報
var Version string

func main() {
	var opts options.Options
	opts.Version = func() {
		fmt.Println(Version)
		os.Exit(0)
	}

	args, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(0)
	}

	if len(args) < 1 {
		fmt.Println("Need arguments. args=", args)
		os.Exit(1)
	}

	if err := plt(args, opts); err != nil {
		panic(err)
	}
}

func plt(args []string, opts options.Options) error {
	l := len(args)
	if l < 1 {
		return errors.New("引数が不足しています。")
	}

	// 引数が一つの場合は標準入力からデータ読み取り
	// 引数が２つ以上のときは、ファイル読み取り
	var r *os.File
	ofn := args[0]
	if l < 2 {
		r = os.Stdin
	} else {
		// TODO 入力データは複数指定できるようにする
		var err error
		r, err = os.Open(args[1])
		if err != nil {
			return err
		}
		defer r.Close()
	}

	return writeGraph(r, opts, ofn)
}

func writeGraph(r *os.File, opts options.Options, ofn string) error {
	p, err := plot.New()
	if err != nil {
		return err
	}

	lines, err := readLines(r)
	if err != nil {
		return err
	}
	graphs, err := graph.GetGraphs(lines, opts)
	if err != nil {
		return err
	}

	p.Title.Text = "Plotutil example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	for _, v := range graphs {
		t := v.Title
		d := v.Data
		if err := plotutil.AddLinePoints(p, t, d); err != nil {
			return err
		}
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, ofn); err != nil {
		panic(err)
	}
	return nil
}

func readLines(r *os.File) ([]string, error) {
	lines := make([]string, 0)
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := sc.Text()
		line = strings.TrimSpace(line)
		lines = append(lines, line)
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

// randomPoints returns some random x, y points.
func randomPoints(n int) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i := range pts {
		if i == 0 {
			pts[i].X = rand.Float64()
		} else {
			pts[i].X = pts[i-1].X + rand.Float64()
		}
		pts[i].Y = pts[i].X + 10*rand.Float64()
	}
	return pts
}
