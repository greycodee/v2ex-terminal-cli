package tui

import "github.com/rivo/tview"

func (t *TUI) topicList() {
	tp := tview.NewList()
	t.Topic = tp
	t.Topic.SetBorder(true).SetTitle("主题")
	t.Topic.SetSelectedFunc(t.handleSelect)
}

func (t *TUI) handleSelect(index int, mainText string, secondaryText string, shortcut rune)  {
	t.content.SetText(mainText)
}