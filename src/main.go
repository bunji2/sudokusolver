// 数独エディタ

package main

import (
	"fyne.io/fyne/app"
)

var conf Config

func main() {
	ap := app.New()

	var err error

	// 設定ファイルの読み込み
	conf, err = LoadConfig()
	if err != nil {
		return
		//dialog.ShowError(err, w)
	}

	// 数独エディタの作成
	s := NewSudokuEditor(ap, conf.Width, conf.Height, conf.InitCells)
	s.Show()
}
