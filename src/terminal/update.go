package terminal

import (
	tea "github.com/charmbracelet/bubbletea"
)

type StartExecution struct{}
type ExecutionFinished struct{}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case StartExecution:
        // Start the function execution
        m.ExecutionTime.Start()
        m.IsRunning = true

    case ExecutionFinished:
        // Stop the stopwatch
        m.ExecutionTime.Stop()
        m.IsRunning = false

	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.Table.Focused() {
				m.Table.Blur()
			} else {
				m.Table.Focus()
			}
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
	//m.Table, cmd = m.Table.Update(msg)
	return m, cmd
}
