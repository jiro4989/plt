package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	flags "github.com/jessevdk/go-flags"
	"github.com/jiro4989/plt/graph"
	options "github.com/jiro4989/plt/internal/options"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

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
	var (
		r   *os.File
		ifn string // 入力ファイル
		ofn string // 出力ファイル名
	)
	if l < 2 {
		ofn = args[0]
		r = os.Stdin
	} else {
		ifn = args[0]
		ofn = args[1]
		// TODO 入力データは複数指定できるようにする
		var err error
		r, err = os.Open(ifn)
		if err != nil {
			return err
		}
		defer r.Close()
	}
	lines, err := readLines(r)
	if err != nil {
		return err
	}

	return writeGraph(lines, opts, ofn)
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

// writeGraph は入力文字列からグラフ画像を出力する。
func writeGraph(lines []string, opts options.Options, ofn string) error {
	p, err := plot.New()
	if err != nil {
		return err
	}

	// グラフ出力のためのプロットデータを算出する
	graphs, err := graph.GetGraphs(lines, opts)
	if err != nil {
		return err
	}

	p.Title.Text = opts.Title
	p.X.Label.Text = opts.XLabel
	p.Y.Label.Text = opts.YLabel

	for _, v := range graphs {
		t := v.Title
		d := v.Data
		if err := plotutil.AddLinePoints(p, t, d); err != nil {
			return err
		}
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, ofn); err != nil {
		return err
	}
	return nil
}
