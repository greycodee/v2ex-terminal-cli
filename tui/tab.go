package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/greycodee/v2ex-terminal-cli/api"
	"github.com/greycodee/v2ex-terminal-cli/urls"
	"github.com/rivo/tview"
)

func (t *TUI) dropDown() *tview.DropDown {
	d := tview.NewDropDown().
		SetLabel("选择Tab: ").
		SetOptions(TabList, func(text string, index int) {
			//t.Topic.SetText(urls.TabLinkList[index]["link"])
			list, err := api.TabTopicList(urls.TabUrl(urls.TabLinkList[index]["link"]))
			if err != nil {
				return
			}
			t.Topic.Clear()
			for _,v := range list {
				t.Topic.AddItem(v.Title,v.Member,0, nil).ShowSecondaryText(false)
			}

		})
	d.SetBorder(true).SetTitle("节点").SetBackgroundColor(tcell.ColorGray)

	return d
}

