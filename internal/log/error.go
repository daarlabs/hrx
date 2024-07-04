package log

import (
	"fmt"
	
	"github.com/charmbracelet/lipgloss"
)

func Error(err error) {
	fmt.Println(
		lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0048")).Render(err.Error()),
	)
}
