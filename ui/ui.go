package ui

import (
	"CLITOOL/reminders"
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	Choices  []string
	Cursor   int
	Selected map[int]struct{}

	started map[int]bool
}

func initialModel() model {
	return model{
		Choices:  []string{"Drink water", "Go for a walk", "Change body posture"},
		Selected: make(map[int]struct{}),
		started:  make(map[int]bool),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.Cursor < len(m.Choices)-1 {
				m.Cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := m.Selected[m.Cursor]
			if ok {
				delete(m.Selected, m.Cursor)
			} else {
				m.Selected[m.Cursor] = struct{}{}
			}
		}

		for index := range m.Selected {
			if m.started[index] {
				continue
			}
			m.started[index] = true

			choice := m.Choices[index]
			switch choice {
			case "Drink water":
				go reminders.Start(reminders.Reminder{
					Name:     "Hydration",
					Interval: 15 * time.Minute,
					Message:  "Drink some water 💧",
				})
			case "Go for a walk":
				go reminders.Start(reminders.Reminder{
					Name:     "Walk",
					Interval: 30 * time.Minute,
					Message:  "It's time for a walk 🚶",
				})
			case "Change body posture":
				go reminders.Start(reminders.Reminder{
					Name:     "Posture",
					Interval: 10 * time.Minute,
					Message:  "Change your posture 🧍",
				})
			}
		}
	}
	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil

}

func (m model) View() string {
	s := "Choose your reminders:\n\n"

	for i, choice := range m.Choices {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.Selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress space or enter to select | q to Quit.\n"
	return s
}

func Run() (model, error) {
	p := tea.NewProgram(initialModel())
	m, err := p.Run()
	if err != nil {
		return model{}, err
	}
	return m.(model), nil
}
