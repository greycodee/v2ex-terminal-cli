package tui

import "github.com/rivo/tview"

type TUI struct {
	node tview.DropDown
	topic tview.List
	title tview.Box
	content tview.TextView
	comment tview.Pages

	mainView tview.Flex
}

func (t *TUI) Start() error {

	nodeAndTopic := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(t.node.SetBorder(true).SetTitle("选择节点"),5,1,false).
		AddItem(t.topic.SetBorder(true).SetTitle("选择主题"),0,2,false)

	mainView := tview.NewFlex().
		AddItem(nodeAndTopic,0,1,false).
		AddItem(t.content.SetBorder(true).SetTitle("内容"),0,2,false)

	return tview.NewApplication().SetRoot(mainView,true).EnableMouse(true).Run()
}

