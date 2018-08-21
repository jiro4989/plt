package option

// options オプション引数
type Options struct {
	Version     func() `short:"v" long:"version" description:"バージョン情報"`
	Type        string `short:"t" long:"type" description:"出力グラフタイプ" default:"line"`
	Separator   string `short:"s" long:"separator" description:"入力データのセパレータ" default:"\t"`
	Width       int    `long:"width" description:"出力画像横幅(px)" default:"1280"`
	Height      int    `long:"height" description:"出力画像縦幅(px)" default:"720"`
	Direction   string `short:"d" long:"direction" description:"データの増加方向" default:"vertical"`
	NoHeader    bool   `long:"noheader" description:"列ヘッダなしフラグ"`
	NoRowHeader bool   `long:"norowheader" description:"行ヘッダなしフラグ"`
}
