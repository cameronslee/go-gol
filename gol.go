// gol
/* Any live cell with fewer than two live neighbors dies, as if by underpopulation.
 *
 * Any live cell with two or three live neighbors lives on to the next generation.
 *
 * Any live cell with more than three live neighbors dies, as if by overpopulation.
 *
 * Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
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
		cursorX: ROWS / 2,
		cursorY: COLS / 2,
	}
}

func in_bounds(i, j int) bool {
	return i >= 0 && j >= 0 && i < ROWS && j < COLS
}

// Compute transitions (count neighbors and update)
func gol_run(m model) {
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
				m.grid[m.cursorX][m.cursorY] = 1
			}

		case "down", "j":
			if m.cursorY < ROWS-1 {
				m.cursorY++
				m.grid[m.cursorX][m.cursorY] = 1
			}

		case "left", "h":
			if m.cursorX > 0 {
				m.cursorX--
				m.grid[m.cursorX][m.cursorY] = 1
			}

		case "right", "l":
			if m.cursorX < COLS-1 {
				m.cursorX++
				m.grid[m.cursorX][m.cursorY] = 1
			}

		// simulate
		case "s":
			for j := 0; j < COLS; j++ {
				for i := 0; i < ROWS; i++ {
					// count neighbors
					neighbors := 0
					if in_bounds(i-1, j-1) && m.grid[i-1][j-1] == 1 {
						neighbors++
					}
					if in_bounds(i, j-1) && m.grid[i][j-1] == 1 {
						neighbors++
					}
					if in_bounds(i+1, j-1) && m.grid[i+1][j-1] == 1 {
						neighbors++
					}
					if in_bounds(i-1, j) && m.grid[i-1][j] == 1 {
						neighbors++
					}
					if in_bounds(i+1, j) && m.grid[i+1][j] == 1 {
						neighbors++
					}
					if in_bounds(i-1, j+1) && m.grid[i-1][j+1] == 1 {
						neighbors++
					}
					if in_bounds(i, j+1) && m.grid[i][j+1] == 1 {
						neighbors++
					}
					if in_bounds(i+1, j+1) && m.grid[i+1][j+1] == 1 {
						neighbors++
					}

					// alive transition
					if m.grid[i][j] == 1 {
						if neighbors < 2 || neighbors > 3 {
							m.grid[i][j] = 0
						} else {
							// live on
						}
						// dead transition
					} else {
						if neighbors == 3 {
							m.grid[i][j] = 1
						}
					}
				}
			}

		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	return m, nil
}

func (m model) View() string {
	// The header
	s := "GOL\nMove the cursor to add life, press s to simulate life\n\n"

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
