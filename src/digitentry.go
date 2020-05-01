// digitentry.go
// 数字一文字だけを受け付けるエントリ

package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

type digitEntry struct {
	widget.Entry
	num int
}

func newDigitEntry(num int) (e *digitEntry) {
	e = &digitEntry{
		num: num,
	}
	e.ExtendBaseWidget(e)
	e.setLabel(num)
	//e.SetText(fmt.Sprintf("%d", num))
	return
}

func (e *digitEntry) setLabel(num int) {
	label := ""
	if num > 0 && num <= 9 {
		label = fmt.Sprintf("%d", num)
	} else {
		num = 0
	}
	e.num = num
	e.SetText(label)
}

func (e *digitEntry) TypedRune(r rune) {
	if conf.Debug {
		fmt.Println("TypedRune =", r, string(r))
	}
	// 入力文字が 0 ～ 9 の場合のみ受け付ける
	if r >= '0' && r <= '9' {
		e.setLabel(int(r - '0'))
	}
}

func (e *digitEntry) TypedKey(key *fyne.KeyEvent) {
	if conf.Debug {
		fmt.Println("Typedkey =", key, key.Name)
	}

	// スペースキーとバックスペースキーが入力されたときは
	// エントリを空にする
	switch key.Name {
	case "Space":
		e.setLabel(0)
	case "BackSpace":
		e.setLabel(0)
	default:
		e.Entry.TypedKey(key)
	}
}

/*
func (e *digitEntry) CreateRenderer() fyne.WidgetRenderer {
	return e.Entry.CreateRenderer()
}

func (e *digitEntry) Disable() {
	e.Entry.Disable()
}

func (e *digitEntry) DoubleTapped(ev *fyne.PointEvent) {
	e.Entry.DoubleTapped(ev)
}

func (e *digitEntry) DragEnd() {
	e.Entry.DragEnd()
}

func (e *digitEntry) Dragged(d *fyne.DragEvent) {
	e.Entry.Dragged(d)
}

func (e *digitEntry) Enable() {
	e.Entry.Enable()
}

func (e *digitEntry) ExtendBaseWidget(wid fyne.Widget) {
	//e.Entry.ExtendBaseWidget(wid)
}

func (e *digitEntry) FocusGained() {
	e.Entry.FocusGained()
}

func (e *digitEntry) FocusLost() {
	e.Entry.FocusLost()
}

func (e *digitEntry) Focused() bool {
	return e.Entry.Focused()
}

func (e *digitEntry) KeyDown(key *fyne.KeyEvent) {
	e.Entry.KeyDown(key)
}

func (e *digitEntry) KeyUp(key *fyne.KeyEvent) {
	e.Entry.KeyUp(key)
}

func (e *digitEntry) MinSize() fyne.Size {
	return e.Entry.MinSize()
}

func (e *digitEntry) MouseDown(m *desktop.MouseEvent) {
	e.Entry.MouseDown(m)
}

func (e *digitEntry) MouseUp(m *desktop.MouseEvent) {
	e.Entry.MouseUp(m)
}

func (e *digitEntry) SelectedText() string {
	return e.Entry.SelectedText()
}

func (e *digitEntry) SetPlaceHolder(text string) {
	e.Entry.SetPlaceHolder(text)
}

func (e *digitEntry) SetReadOnly(ro bool) {
	e.Entry.SetReadOnly(ro)
}

func (e *digitEntry) SetText(text string) {
	e.Entry.SetText(text)
}

func (e *digitEntry) Tapped(ev *fyne.PointEvent) {
	e.Entry.Tapped(ev)
}

func (e *digitEntry) TappedSecondary(pe *fyne.PointEvent) {
	e.Entry.TappedSecondary(pe)
}

func (e *digitEntry) TypedShortcut(shortcut fyne.Shortcut) {
	e.Entry.TypedShortcut(shortcut)
}
*/
