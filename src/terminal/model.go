package terminal

import (
	"github.com/charmbracelet/bubbles/table"
	"time"
)

type Model struct {
	Table          table.Model
	FilesRead      int
	ExecutionTime  time.Duration
}