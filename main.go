package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type status int

func (s status) getNext() status {
	if s == done {
		return todo
	} else {
		return s + 1
	}
}

func (s status) getPrev() status {
	if s == todo {
		return done
	} else {
		return s - 1
	}
}

const margin = 4

var board *Board

const (
	todo status = iota
	inProgress
	done
)

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close() // Close the file right before the Main function finishes execution

	board = NewBoard()
	board.initLists()
	p := tea.NewProgram(board)
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
