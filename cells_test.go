package box_drawing

import (
	"testing"
)

type MockContenter struct {
	H, W int
	R    string
}

func (m *MockContenter) Height() int {
	return m.H
}

func (m *MockContenter) Width() int {
	return m.W
}

func (m *MockContenter) String() string {
	return m.R
}

func (m *MockContenter) Lines() []string {
	return []string{m.R}
}

func TestNewCell(t *testing.T) {
	tab, err := NewCell()

	if tab == nil || err != nil {
		t.Error("Unexpected result")
	}
}

func TestCellAdd(t *testing.T) {
	tab, _ := NewCell()
	tab.SetInnerWidth(10)
	tab.SetInnerHeight(10)

	if tab.Add(NewBoxByRune(' ', 9, 9)) != nil {
		t.Error("Unexpected result")
	}

	if tab.Add(NewBoxByRune(' ', 9, 1)) == nil {
		t.Error("Unexpected result")
	}

	if tab.Add(NewBoxByRune(' ', 1, 9)) == nil {
		t.Error("Unexpected result")
	}
}

func TestCellWidth(t *testing.T) {
	tab, _ := NewCell()
	tab.SetInnerWidth(10)

	if tab.Width() != 10 {
		t.Error("Unexpected result")
	}
}

func TestCellWidthWithBorder(t *testing.T) {
	tab, _ := NewCell(
		NewBorderBoxDrawing(),
		NewBoxByRune(' ', 10, 10),
	)

	if tab.Width() != 12 {
		t.Errorf("Unexpected result: %d", tab.Width())
	}
}

func TestCellHeight(t *testing.T) {
	tab, _ := NewCell()
	tab.SetInnerHeight(10)

	if tab.Height() != 10 {
		t.Error("Unexpected result")
	}
}
