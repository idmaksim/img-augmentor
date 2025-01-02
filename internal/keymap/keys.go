package keymap

import "github.com/charmbracelet/bubbles/key"

var Keys = Keymap{
	Up: key.NewBinding(
		key.WithKeys("up", "w"),
		key.WithHelp("↑/w", "move up"),
	),

	Down: key.NewBinding(
		key.WithKeys("down", "s"),
		key.WithHelp("↓/s", "move down"),
	),

	Enter: key.NewBinding(
		key.WithKeys("enter", " "),
		key.WithHelp("enter/space", "select"),
	),

	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "q"),
		key.WithHelp("ctrl+c/q", "quit"),
	),
}
