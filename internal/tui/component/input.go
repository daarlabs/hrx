package component

import "github.com/charmbracelet/bubbles/textinput"

func Input() textinput.Model {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 255
	ti.Width = 30
	return ti
}
