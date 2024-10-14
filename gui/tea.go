package gui

import (
	"log"

	"grr/app/domain/model"

	tea "github.com/charmbracelet/bubbletea"
)

type m struct {
	model.Model
}

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
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m m) View() string {
	return "Hello, World!"
}
