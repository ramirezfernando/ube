package terminal

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/help"
	"fmt"
)

type ClocCompleted struct {
    Table table.Model
    Help  help.Model
    Err   error
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case ClocCompleted:
		fmt.Println("ClockCompleted message received")
		if msg.Err != nil {
			return m, tea.Quit
		}
		m.ExecutionTime.Stop()
		m.IsRunning = false
		m.Table = msg.Table
		m.Help = msg.Help
		return m, cmd
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			m.ExecutionTime.Stop()
			return m, tea.Quit
		case "u", "up":
			m.Table.MoveUp(1)
		case "d", "down":
			m.Table.MoveDown(1)
		}
	}
	m.ExecutionTime, cmd = m.ExecutionTime.Update(msg)
	m.Table, _ = m.Table.Update(msg)
	return m, cmd
}
