// Package main is based on the tutorial from https://github.com/charmbracelet/bubbletea?tab=readme-ov-file#tutorial
package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// struct model hold's the application's state
type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

// creates a new model
func newModel() model {
	return model{
		// grocery list
		choices: []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},

		// this map uses slice indicies as keys
		selected: make(map[int]struct{}),
	}
}

// can return a command that could perform I/O
func (m model) Init() tea.Cmd {
	// nil means "no I/O right now, please"
	return nil
}

// called when things happen
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Is it a key press?
	case tea.KeyMsg:
		// Which key?
		switch msg.String() {
		// Quit
		case "ctrl+c", "q":
			return m, tea.Quit

			// cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

			// cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

			// toggle selected state
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	// return the model without a command
	return m, nil
}

// renders the UI
func (m model) View() string {
	// view header
	s := "What should we buy at the market?\n\n"

	// view body

	// iterate over choices
	for i, choice := range m.choices {
		// is the cursor pointing at this choice?
		cursor := " " // no cursor visible
		if m.cursor == i {
			cursor = ">" // show cursor glyph
		}

		// is choice selected?
		checked := " " // don't show selected glyph
		if _, ok := m.selected[i]; ok {
			checked = "x" // show selected glyph
		}

		// render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// view footer
	s += "\nPress q to quit.\n"

	// return the ui
	return s
}

// entry point
func main() {
	p := tea.NewProgram(newModel())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
