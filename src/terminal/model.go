package terminal

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/table"
)

type Model struct {
	Table table.Model
	Help  help.Model
}
