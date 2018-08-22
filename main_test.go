package main

import (
	"testing"

	options "github.com/jiro4989/plt/internal/options"
	"github.com/stretchr/testify/assert"
)

type TestPltData struct {
	args []string
	opts options.Options
	err  error
}

func TestPlt(t *testing.T) {
	testdatas := []TestPltData{
		TestPltData{
			args: []string{"testdata/in/value_only.tsv", "testdata/out/value_only.png"},
			opts: options.Options{},
			err:  nil,
		},
	}
	for _, v := range testdatas {
		err := plt(v.args, v.opts)
		assert.Equal(t, v.err, err)
	}

	testdatas = []TestPltData{
		TestPltData{
			// 引数未指定の場合
			args: []string{},
			opts: options.Options{},
		},
		TestPltData{
			// 存在しない入力ファイル
			args: []string{"hogefugatmp", "testdata/out/hogefugatmp.png"},
			opts: options.Options{},
		},
		TestPltData{
			// 出力先ディレクトリが存在しない
			args: []string{"testdata/in/value_only.tsv", "testdata/hogefugatmp/value_only.png"},
			opts: options.Options{},
		},
	}
	for _, v := range testdatas {
		err := plt(v.args, v.opts)
		assert.Error(t, err)
	}
}

func TestWriteGraph(t *testing.T) {
	// TODO
}

func TestReadLines(t *testing.T) {
	// TODO
}
