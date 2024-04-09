package terminal

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/lipgloss"
)

type keyMap struct {
	Up   key.Binding
	Down key.Binding
	Quit key.Binding
}

// N/A
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{}
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.Quit}
}

var keys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "u"),
		key.WithHelp("↑/u", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "d"),
		key.WithHelp("↓/d", "move down"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

func (m Model) View() string {
	out := "Elapsed: " + m.ExecutionTime.View()

	if !m.IsRunning {
		out += "\n" + baseStyle.Render(m.Table.View())
		out += "\n" + m.Help.View(keys)
	}

	return out
}
