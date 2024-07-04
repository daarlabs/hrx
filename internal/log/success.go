package log

import (
	"fmt"
	
	"github.com/charmbracelet/lipgloss"
)

func Success(msg string) {
	fmt.Println(
		lipgloss.NewStyle().Foreground(lipgloss.Color("#04B575")).Render(msg),
	)
}
