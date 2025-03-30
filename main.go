package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type status int

const divisor = 4

const (
	todo status = iota
	inProgress
	done
)

/* Styling */
var (
	columnStyle = lipgloss.NewStyle().
			Padding(1, 2)
	focusedStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62"))
	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))
)

// Custom Item
type Task struct {
	status      status
	title       string
	description string
}

// List.item interface implementation.
func (t Task) FilterValue() string {
	return t.title
}

func (t Task) Title() string {
	return t.title
}

func (t Task) Description() string {
	return t.description
}

// Main Model
type Model struct {
	loaded   bool // To make sure that list is not accessed before it's Initialised.
	focused  status
	lists    []list.Model
	err      error
	quitting bool
}

func New() *Model {
	return &Model{}
}

// TODO: Go To Next List
func (m *Model) Next() {
	if m.focused == done {
		m.focused = todo
	} else {
		m.focused++
	}
}

// TODO: Go To Previous List
func (m *Model) Prev() {
	if m.focused == todo {
		m.focused = done
	} else {
		m.focused--
	}
}

// On Startup we get a Window Size message, so we need Width and Height as Well.
func (m *Model) initLists(width, height int) {
	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), width/divisor, height/2)
	defaultList.SetShowHelp(false)
	m.lists = []list.Model{defaultList, defaultList, defaultList}

	// Init To Do
	m.lists[todo].Title = "To Do"
	m.lists[todo].SetItems([]list.Item{
		Task{
			status:      todo,
			title:       "Buy Milk",
			description: "Strawberry Milk",
		},
		Task{
			status:      todo,
			title:       "Eat Sushi",
			description: "Miso Soup, and Rice",
		},
		Task{
			status:      todo,
			title:       "Fold Laundry",
			description: "or Wear Wrinkly Tees",
		},
	})

	// Init in Progress
	m.lists[inProgress].Title = "In Progress"
	m.lists[inProgress].SetItems([]list.Item{
		Task{
			status:      inProgress,
			title:       "Go Proj",
			description: "CLI Kanban",
		},
	})

	// Init To Do
	m.lists[done].Title = "Done"
	m.lists[done].SetItems([]list.Item{
		Task{
			status:      done,
			title:       "Portfolio",
			description: "React, JS, Tailwind",
		},
	})

}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.loaded {
			columnStyle.Width(msg.Width / divisor)
			focusedStyle.Width(msg.Width / divisor)

			columnStyle.Height(msg.Height - divisor)
			focusedStyle.Height(msg.Height - divisor)

			m.initLists(msg.Width, msg.Height)
			m.loaded = true
		}

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit

		case "left", "j":
			m.Prev()

		case "right", "l":
			m.Next()
		}
	}

	var cmd tea.Cmd
	m.lists[m.focused], cmd = m.lists[m.focused].Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if m.quitting {
		return ""
	}

	if m.loaded {
		todoView := m.lists[todo].View()
		inProgressView := m.lists[inProgress].View()
		doneView := m.lists[done].View()

		switch m.focused {
		case inProgress:
			return lipgloss.JoinHorizontal(
				lipgloss.Left,
				// .View() to get their String Rep.
				columnStyle.Render(todoView),
				focusedStyle.Render(inProgressView),
				columnStyle.Render(doneView),
			)

		case done:
			return lipgloss.JoinHorizontal(
				lipgloss.Left,
				// .View() to get their String Rep.
				columnStyle.Render(todoView),
				columnStyle.Render(inProgressView),
				focusedStyle.Render(doneView),
			)

		default:
			return lipgloss.JoinHorizontal(
				lipgloss.Left,
				// .View() to get their String Rep.
				focusedStyle.Render(todoView),
				columnStyle.Render(inProgressView),
				columnStyle.Render(doneView),
			)
		}

	} else {
		return "loading..."
	}

}

func main() {
	m := New()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
