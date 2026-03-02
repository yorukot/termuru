package settings

// Hotkeys defines key bindings for all keyboard actions in the app.
type Hotkeys struct {
	Quit       []string
	Select     []string
	CursorUp   []string
	CursorDown []string
}

// Hotkey is the global key map used by all keyboard handlers.
var Hotkey = Hotkeys{
	Quit:       []string{"q", "ctrl+c"},
	Select:     []string{"enter"},
	CursorUp:   []string{"up", "k"},
	CursorDown: []string{"down", "j"},
}
