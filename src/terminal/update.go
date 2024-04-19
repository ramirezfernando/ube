package terminal

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type CompletionResult struct {
	Table table.Model
	Help  help.Model
	Err   error
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case CompletionResult:
		if msg.Err != nil {
			return m, tea.Quit
		}
		m.ElapsedTime.Stop()
		m.IsRunning = false
		m.Table = msg.Table
		m.Help = msg.Help
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			m.ElapsedTime.Stop()
			return m, tea.Quit
		case "u", "up":
			m.Table.MoveUp(1)
		case "d", "down":
			m.Table.MoveDown(1)
		}
	}

	if m.IsRunning {
		m.ElapsedTime, cmd = m.ElapsedTime.Update(msg)
		m.Table, _ = m.Table.Update(msg)
	}

	return m, cmd
}
