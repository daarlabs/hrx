package generate

import (
	"fmt"
	
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	
	"github.com/daarlabs/hrx/internal/config"
)

type generateModel struct {
	step          int
	filenameInput textinput.Model
	filepathInput textinput.Model
}

func (m generateModel) Init() tea.Cmd {
	return nil
}

func (m generateModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.step == 1 {
				config.Config.Generate.Name = m.filenameInput.Value()
			}
			if m.step == 2 {
				config.Config.Generate.Path = m.filepathInput.Value()
			}
			if m.step == 2 {
				return m, tea.Quit
			}
			m.step++
			return m, cmd
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	}
	if m.step == 1 {
		m.filenameInput, cmd = m.filenameInput.Update(msg)
	}
	if m.step == 2 {
		m.filepathInput, cmd = m.filepathInput.Update(msg)
	}
	return m, cmd
}

func (m generateModel) View() string {
	if m.step == 1 {
		return fmt.Sprintf(
			"How do you want be the file named?\n\n%s",
			m.filenameInput.View(),
		) + "\n"
	}
	if m.step == 2 {
		return fmt.Sprintf(
			"Where do you want to create the file?\n\n%s",
			m.filepathInput.View(),
		) + "\n"
	}
	return ""
}
