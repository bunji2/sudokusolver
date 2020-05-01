// sudoku.go
// 数独の問題エディタの GUI

package main

import (
	"fmt"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

const (
	title = "sudoku"
)

// SudokuEditor は数独エディタの構造体型
type SudokuEditor struct {
	window fyne.Window
	width  int
	height int
	//cells  []*widget.Entry
	cells    []*digitEntry
	initNums []int
}

// NewSudokuEditor は数独用エディタの GUI を作成する関数
//
// width は横のセル数、height は縦のセル数、initCells は各セルの初期値(0～9)。
func NewSudokuEditor(app fyne.App, width, height int, initNums []int) SudokuEditor {
	// cells は各セルの初期値を持つ数字エントリ群を格納する配列
	cells := []*digitEntry{}
	for i := 0; i < width*height; i++ {
		cells = append(cells, newDigitEntry(initNums[i]))
	}
	return SudokuEditor{
		window:   app.NewWindow(title),
		width:    width,
		height:   height,
		cells:    cells,
		initNums: initNums,
	}
}

// clear は全セルをクリアする関数
func (s SudokuEditor) clear() {
	for i := 0; i < s.width*s.height; i++ {
		s.cells[i].setLabel(0)
	}
}

// reset は全セルを書き換えるする関数
func (s SudokuEditor) reset() {
	if conf.Debug {
		fmt.Println("reset")
	}
	for i := 0; i < s.width*s.height; i++ {
		s.cells[i].setLabel(s.initNums[i])
	}
}

// rewrite は全セルを書き換える関数
func (s SudokuEditor) rewrite(nums []int) {
	if conf.Debug {
		fmt.Println("rewrite")
	}
	for i := 0; i < s.width*s.height; i++ {
		s.cells[i].setLabel(nums[i])
	}
}

// resolve はセルの値を解決する関数
func (s SudokuEditor) resolve() {
	if conf.Debug {
		fmt.Println("resolve")
	}
	// 各セルの値を取り出して数字の配列を作成する
	values := make([]int, s.width*s.height)
	for i := 0; i < s.width*s.height; i++ {
		values[i] = s.cells[i].num
	}
	if conf.Debug {
		fmt.Println(values)
	}
	// ここで SMT Solver を呼び出す
	values, err := smtsolver(values)
	if err != nil {
		dialog.ShowError(err, s.window)
		return
	}
	// 結果を画面に反映
	s.rewrite(values)
}

// copy はセルの値をクリップボードにコピーする関数
func (s SudokuEditor) copy() {
	if conf.Debug {
		fmt.Println("copy")
	}
	// 各セルの値を取り出して数字の配列を作成する
	values := make([]string, s.width*s.height)
	for i := 0; i < s.width*s.height; i++ {
		values[i] = fmt.Sprintf("%d", s.cells[i].num)
	}
	if conf.Debug {
		fmt.Println(values)
	}
	s.window.Clipboard().SetContent(strings.Join(values, ","))
}

// loadUI は数独エディタの GUI を作成する関数
func (s SudokuEditor) loadUI() (r fyne.CanvasObject) {
	objects := make([]fyne.CanvasObject, len(s.cells))
	for i := 0; i < len(s.cells); i++ {
		objects[i] = s.cells[i]
	}
	cells := fyne.NewContainerWithLayout(layout.NewGridLayoutWithColumns(s.width), objects...)
	clearButton := widget.NewButton("clear", s.clear)
	resetButton := widget.NewButton("reset", s.reset)
	copyButton := widget.NewButton("copy", s.copy)
	resolveButton := widget.NewButton("resolve", s.resolve)
	ctrls := fyne.NewContainerWithLayout(layout.NewGridLayout(4), clearButton, resetButton, copyButton, resolveButton)
	r = widget.NewVBox(cells, ctrls)
	return
}

// Show は数独エディタの GUI を表示する関数
func (s SudokuEditor) Show() {
	s.window.SetContent(s.loadUI())
	s.window.SetIcon(resourceIconPng)
	s.window.Resize(fyne.NewSize(480, 320))
	s.window.SetFixedSize(true)
	s.window.ShowAndRun()

}
