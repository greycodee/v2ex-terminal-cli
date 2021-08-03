package tui

import (
	"github.com/greycodee/v2ex-terminal-cli/api"
	"github.com/rivo/tview"
	"github.com/charmbracelet/glamour"
)

func (t *TUI) topicList() {
	tp := tview.NewList()
	t.Topic = tp
	t.Topic.SetBorder(true).SetTitle("主题")
	t.Topic.SetSelectedFunc(t.handleSelect)
}

func (t *TUI) handleSelect(index int, mainText string, secondaryText string, shortcut rune)  {
	topicInfo,err := api.GetTopicInfo(secondaryText)
	if err != nil {
		return
	}
	md,err := glamour.Render(topicInfo.Content,"dark")
	if err != nil {
		return
	}
	t.content.SetText(md)
}