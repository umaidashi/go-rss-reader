package gui

import (
	"log"

	"grr/app/domain/model"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type m struct {
	model.Model
	quitting bool
	err      error
}

var quitKeys = key.NewBinding(
	key.WithKeys("q", "esc", "ctrl+c"),
	key.WithHelp("", "press q to quit"),
)

func Start() {
	p := tea.NewProgram(m{})
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func (m m) Init() tea.Cmd {
	return nil
}

func (m m) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, quitKeys) {
			m.quitting = true
			return m, tea.Quit
		}
		return m, nil
	case error:
		m.err = msg
		return m, nil
	}

	return m, nil
}

func (m m) View() string {
	return "Hello, World!"
}
