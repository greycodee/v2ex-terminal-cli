package tui

import (
	"github.com/rivo/tview"
)

type TUI struct {
	node    tview.Box
	Topic   *tview.List
	title   tview.Box
	content *tview.TextView
	comment tview.Box

	mainView tview.Flex
}



func (t *TUI) Start() error {

	t.topicList()

	textView := tview.NewTextView()
	textView.SetTitle("内容").SetBorder(true)
	t.content = textView

	nodeAndTopic := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(t.dropDown(),3,1,false).
		AddItem(t.Topic,0,2,false)

	mainV := tview.NewFlex().
		AddItem(nodeAndTopic,0,1,false).
		AddItem(t.content,0,2,false)

	//t.mainView = mainV
	return tview.NewApplication().SetRoot(mainV,true).EnableMouse(true).Run()
}
