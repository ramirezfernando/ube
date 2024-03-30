package terminal

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/stopwatch"
	"github.com/charmbracelet/bubbles/table"
)

type Model struct {
	ExecutionTime stopwatch.Model
	IsRunning     bool
	Table         table.Model
	Help          help.Model
}
