package dashboard

import (
	"charm.land/bubbles/v2/textinput"
	"github.com/yorukot/termuru/internal/settings"
)

type Model struct {
	cursor    int
	offset    int
	searchBar textinput.Model
}

func NewModel() Model {

	return Model{
		cursor:    0,
		offset:    0,
		searchBar: newSearchBarTextInput(),
	}
}

func newSearchBarTextInput() textinput.Model {
	ti := textinput.New()
	ti.Prompt = ""
	ti.Placeholder = "(" + settings.Hotkey.Search[0] + ") Type something"
	ti.Blur()
	return ti
}

func (m *Model) SyncItemCount(itemCount, termHeight int) {
	if itemCount == 0 {
		m.cursor = 0
		m.offset = 0
		return
	}

	capacity := listCapacity(termHeight)
	m.clamp(itemCount, capacity)
	m.syncOffset(itemCount, capacity)
}
