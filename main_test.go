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
}

func TestWriteGraph(t *testing.T) {
	// TODO
}

func TestReadLines(t *testing.T) {
	// TODO
}
