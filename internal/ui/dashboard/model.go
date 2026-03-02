package dashboard

import (
	"slices"

	tea "charm.land/bubbletea/v2"

	"github.com/yorukot/termuru/internal/settings"
)

type Model struct {
	cursor int
	offset int
}

func NewModel() Model {
	return Model{cursor: 0, offset: 0}
}

func (m Model) Update(msg tea.Msg, itemCount, termWidth, termHeight int) (Model, tea.Cmd) {
	if itemCount == 0 {
		m.cursor = 0
		m.offset = 0
		return m, nil
	}

	capacity := listCapacity(termHeight)
	m.clamp(itemCount, capacity)

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		key := msg.String()
		switch {
		case slices.Contains(settings.Hotkey.CursorUp, key):
			if m.cursor > 0 {
				m.cursor--
			}
		case slices.Contains(settings.Hotkey.CursorDown, key):
			if m.cursor < itemCount-1 {
				m.cursor++
			}
		}
	case tea.MouseWheelMsg:
		mouse := msg.Mouse()
		switch mouse.Button {
		case tea.MouseWheelUp:
			if m.cursor > 0 {
				m.cursor--
			}
		case tea.MouseWheelDown:
			if m.cursor < itemCount-1 {
				m.cursor++
			}
		}
	case tea.MouseClickMsg:
		mouse := msg.Mouse()
		if mouse.Button == tea.MouseLeft {
			if idx, ok := m.itemIndexAt(mouse.X, mouse.Y, itemCount, termWidth, termHeight); ok {
				m.cursor = idx
			}
		}
	}

	m.syncOffset(itemCount, capacity)
	return m, nil
}

func (m Model) Cursor() int {
	return m.cursor
}

func (m Model) Window(itemCount, termHeight int) (start, end int) {
	if itemCount <= 0 {
		return 0, 0
	}

	capacity := listCapacity(termHeight)
	if m.offset > itemCount-1 {
		m.offset = itemCount - 1
	}
	if m.offset < 0 {
		m.offset = 0
	}

	start = m.offset
	end = start + capacity
	end = min(end, itemCount)

	return start, end
}

func (m Model) itemIndexAt(x, y, itemCount, termWidth, termHeight int) (int, bool) {
	if itemCount <= 0 {
		return 0, false
	}
	if x < 0 || y < 0 || x >= termWidth || y >= termHeight {
		return 0, false
	}

	start, end := m.Window(itemCount, termHeight)
	listRow := y - listStartY
	if listRow < 0 {
		return 0, false
	}

	idx := start + listRow
	if idx < start || idx >= end {
		return 0, false
	}

	return idx, true
}

func listCapacity(termHeight int) int {
	if termHeight <= 0 {
		return 10
	}

	// Reserve lines for top/bottom padding, border, title, and help text.
	capacity := termHeight - 8
	if capacity < 1 {
		return 1
	}
	return capacity
}

const listStartY = 4

func (m *Model) clamp(itemCount, capacity int) {
	if m.cursor < 0 {
		m.cursor = 0
	}
	if m.cursor > itemCount-1 {
		m.cursor = itemCount - 1
	}

	maxOffset := itemCount - capacity
	maxOffset = max(maxOffset, 0)

	if m.offset < 0 {
		m.offset = 0
	}
	if m.offset > maxOffset {
		m.offset = maxOffset
	}
}

func (m *Model) syncOffset(itemCount, capacity int) {
	if m.cursor < m.offset {
		m.offset = m.cursor
	}
	if m.cursor >= m.offset+capacity {
		m.offset = m.cursor - capacity + 1
	}

	maxOffset := itemCount - capacity
	maxOffset = max(maxOffset, 0)

	if m.offset > maxOffset {
		m.offset = maxOffset
	}
	if m.offset < 0 {
		m.offset = 0
	}
}
