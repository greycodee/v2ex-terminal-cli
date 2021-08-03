package main

import (
	"fmt"
	"github.com/greycodee/v2ex-terminal-cli/tui"
)

func main()  {
	ui := tui.TUI{}

	err := ui.Start()
	if err != nil {
		fmt.Println(err)
		return
	}


	//list, err := api.TabTopicList(urls.TabCity)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//marshal, err := json.Marshal(list)
	//if err != nil {
	//	return
	//}
	//fmt.Println(string(marshal))




}
