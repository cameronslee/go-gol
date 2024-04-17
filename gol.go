// gol
/*
 *
 *
 */

package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

var ROWS int = 32
var COLS int = 32

type model struct {
	grid    [32][32]int
	cursorX int
	cursorY int
}

func initialModel() model {
	var grid [32][32]int
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			grid[i][j] = 0
		}
	}
	return model{
		grid:    grid,
		cursorX: 0,
		cursorY: 0,
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
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

		case "up", "k":
			if m.cursorY > 0 {
				m.cursorY--
			}

		case "down", "j":
			if m.cursorY < ROWS-1 {
				m.cursorY++
			}

		case "left", "h":
			if m.cursorX > 0 {
				m.cursorX--
			}

		case "right", "l":
			if m.cursorX < COLS-1 {
				m.cursorX++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":

		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	// The header
	s := "GOL\nMove the cursor to simulate life\n\n"

	for j := 0; j < COLS; j++ {
		for i := 0; i < ROWS; i++ {
			cursor := " "
			cell := " "
			if m.cursorX == i && m.cursorY == j {
				cursor = "λ"
				s += fmt.Sprintf("[%s]", cursor)
			} else {
				if m.grid[i][j] == 1 { // 0: dead, 1: alive
					cell = "*"
				}
				s += fmt.Sprintf("[%s]", cell)
			}
		}
		s += "\n"
	}

	// TODO controls for moving around and placing a new cell
	/*
		// Iterate over our choices
		for i, choice := range m.choices {

			// Is the cursor pointing at this choice?
			cursor := " " // no cursor
			if m.cursor == i {
				cursor = ">" // cursor!
			}

			// Is this choice selected?
			checked := " " // not selected
			if _, ok := m.selected[i]; ok {
				checked = "x" // selected!
			}

			// Render the row
			s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
		}
	*/

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
