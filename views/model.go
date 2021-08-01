package views

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/greycodee/v2ex-terminal-cli/views/tab"
	"github.com/greycodee/v2ex-terminal-cli/views/topic"
)

type MainView struct {
	Tab tab.Model
	Topic topic.Model
}

func (m MainView) Init() tea.Cmd{
	return nil
}

func (m MainView) Update(msg tea.Msg) (tea.Model, tea.Cmd)  {
	return m, nil
}

func (m MainView) View() string {
	return lipgloss.JoinHorizontal(lipgloss.Top,m.Tab.View(),m.Topic.View())
}
