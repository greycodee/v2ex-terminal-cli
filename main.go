package main

import (
	"fmt"
	"github.com/greycodee/v2ex-terminal-cli/tui"
)

func main() {
	ui := tui.TUI{}

	err := ui.Start()
	if err != nil {
		fmt.Println(err)
		return
	}
}

