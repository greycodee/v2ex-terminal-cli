package topic

import tea "github.com/charmbracelet/bubbletea"
import "github.com/charmbracelet/lipgloss"


var tabStyle = lipgloss.Border{
	Top:         "*",
	Bottom:      "",
	Left:        "*",
	Right:       "|",
	TopLeft:     "",
	TopRight:    "",
	BottomRight: "",
	BottomLeft:  "",
}

type Model struct {

}

func (m Model) Init() tea.Cmd{
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd)  {
	return m, nil
}

func (m Model) View() string {
	return lipgloss.NewStyle().Width(40).Height(20).Border(tabStyle,false).Align(lipgloss.Center).BorderTop(true).BorderRight(true).Render("topic view")
}

