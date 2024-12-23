// Package main is based on the tutorial from https://github.com/charmbracelet/bubbletea?tab=readme-ov-file#tutorial
package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const url = "https://charm.sh/"

// struct model hold's the application's state
type model struct {
	status int
	err    error
}

type statusMsg int

type errMsg struct{ err error }

func (e errMsg) Error() string { return e.err.Error() }

func checkServer() tea.Msg {
	c := &http.Client{Timeout: 10 * time.Second}
	res, err := c.Get(url)
	if err != nil {
		return errMsg{err}
	}

	return statusMsg(res.StatusCode)
}

// creates a new model
func newModel() model {
	return model{}
}

// can return a command that could perform I/O
func (m model) Init() tea.Cmd {
	return checkServer
}

// called when things happen
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case statusMsg:
		m.status = int(msg)
		return m, tea.Quit

	case errMsg:
		m.err = msg
		return m, tea.Quit

	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}
	}

	// return the model without a command
	return m, nil
}

// renders the UI
func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("\nError: %v\n\n", m.err)
	}

	s := fmt.Sprintf("Checking %s ... ", url)

	if m.status > 0 {
		s += fmt.Sprintf("%d %s!", m.status, http.StatusText(m.status))
	}

	return "\n" + s + "\n\n"
}

// entry point
func main() {
	p := tea.NewProgram(newModel())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
