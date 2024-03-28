package terminal

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/stopwatch"
)

type Model struct {
	ExecutionTime     stopwatch.Model
    IsRunning         bool
    Table             table.Model
    Help              help.Model
}
