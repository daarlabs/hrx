package log

import (
	"fmt"
	
	"github.com/charmbracelet/lipgloss"
)

func Info(msg string) {
	fmt.Println(
		lipgloss.NewStyle().Foreground(lipgloss.Color("#409CFF")).Render(msg),
	)
}
