package terminal

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

func (m Model) View() string {
	return baseStyle.Render(m.Table.View()) + fmt.Sprintf(`
    Files Read: %d
    Execution Time: %v
    `, m.FilesRead, m.ExecutionTime)
}