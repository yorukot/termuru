package settings

// Hotkeys defines key bindings for all keyboard actions in the app.
type Hotkeys struct {
	Quit       []string
	Confirm    []string
	CursorUp   []string
	CursorDown []string
	Search     []string
	Cancel     []string
}

// Hotkey is the global key map used by all keyboard handlers.
var Hotkey = Hotkeys{
	Quit:       []string{"q", "ctrl+c"},
	Confirm:    []string{"enter"},
	CursorUp:   []string{"up", "k"},
	CursorDown: []string{"down", "j"},
	Search:     []string{"/"},
	Cancel:     []string{"esc", "ctrl+c"},
}
