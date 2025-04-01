package main

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	New    key.Binding
	Edit   key.Binding
	Delete key.Binding
	Up     key.Binding
	Down   key.Binding
	Left   key.Binding
	Right  key.Binding
	Enter  key.Binding
	Help   key.Binding
	Quit   key.Binding
	Back   key.Binding
}

// Returns keybindings for the mini help view.
// Both ShortHelp and FullHelp are Part of the key.Map interface.
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

// Returns keybindings for the expanded help view.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{ // 1st Col.
			k.Up,
			k.Down,
			k.Left,
			k.Right,
		},
		{ // 2nd Col.
			k.Help,
			k.Quit,
		},
	}
}

var keys = keyMap{
	New: key.NewBinding(
		key.WithKeys("n"),
		key.WithHelp("n", "New"),
	),

	Edit: key.NewBinding(
		key.WithKeys("e"),
		key.WithHelp("e", "Edit"),
	),

	Delete: key.NewBinding(
		key.WithKeys("backspace"),
		key.WithHelp("backspace", "Delete"),
	),

	Up: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑", "Move Up"),
	),

	Down: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("↓", "Move Down"),
	),

	Left: key.NewBinding(
		key.WithKeys("left", "a"),
		key.WithHelp("←/a", "Move Right"),
	),

	Right: key.NewBinding(
		key.WithKeys("right", "d"),
		key.WithHelp("→/d", "Move Right"),
	),

	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "enter"),
	),

	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "Help"),
	),

	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q/ctrl+c", "Quit"),
	),

	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "Back"),
	),
}
