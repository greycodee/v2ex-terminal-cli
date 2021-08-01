package main

import (
	"fmt"
	"github.com/greycodee/v2ex-terminal-cli/views/tab"
	"github.com/greycodee/v2ex-terminal-cli/views/topic"
	"os"
)
import "github.com/greycodee/v2ex-terminal-cli/views"
import tea "github.com/charmbracelet/bubbletea"

func main(){
	fmt.Println("hello world!")

	p := tea.NewProgram(views.MainView{
		Tab:   tab.Model{},
		Topic: topic.Model{}},
		tea.WithAltScreen(),
	)
	if err := p.Start(); err != nil {
		fmt.Println("could not run program:", err)
		os.Exit(1)
	}
}
