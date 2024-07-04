package generate

import (
	"fmt"
	"os"
	
	tea "github.com/charmbracelet/bubbletea"
	
	"github.com/daarlabs/hrx/internal/component"
)

func Run() {
	p := tea.NewProgram(
		generateModel{
			step:          1,
			filenameInput: component.Input(),
			filepathInput: component.Input(),
		},
	)
	if _, err := p.Run(); err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
}
