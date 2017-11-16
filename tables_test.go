package box_drawing

import (
	"testing"
)

type MockRow struct {
	height, width int
}

func (m *MockRow) Height() int {
	return m.height
}

func (m *MockRow) Width() int {
	return m.width
}

func TestNewTableEmpty(t *testing.T) {
	tab, err := NewTable()

	if tab == nil || err != nil || tab.Draw() != "" {
		t.Error("Unexpected result")
	}
}

func TestAdd(t *testing.T) {
	tab, _ := NewTable()
	tab.SetHeight(10)
	tab.SetWidth(10)

	if tab.Add(&MockRow{9, 9}) != nil {
		t.Error("Unexpected result")
	}

	if tab.Add(&MockRow{9, 0}) == nil {
		t.Error("Unexpected result")
	}

	if tab.Add(&MockRow{0, 9}) == nil {
		t.Error("Unexpected result")
	}
}

func TestHeight(t *testing.T) {
	tab, _ := NewTable()
	height := 100
	tab.SetHeight(height)
	if tab.Height() != height {
		t.Error("Unexpected result")
	}

	height = 1000
	if tab.Height() == height {
		t.Error("Unexpected result")
	}
}

func TestWidth(t *testing.T) {
	tab, _ := NewTable()
	width := 100
	tab.SetWidth(width)
	if tab.Width() != width {
		t.Error("Unexpected result")
	}

	width = 1000
	if tab.Width() == width {
		t.Error("Unexpected result")
	}
}

func TestRowsHeightAndWidth(t *testing.T) {
	tab, _ := NewTable(&MockRow{10, 10}, &MockRow{20, 20})
	if tab.Height() != 30 {
		t.Error("Unexpected result")
	}

	if tab.Width() != 20 {
		t.Error("Unexpected result")
	}
}

func TestNewDecorator(t *testing.T) {
	tab, _ := NewTable(NewDecorator(func(tab *Table) error {
		tab.SetHeight(100)
		return nil
	}))

	if tab.Height() != 100 {
		t.Error("Unexpected result")
	}
}
