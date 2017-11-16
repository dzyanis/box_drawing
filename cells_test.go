package box_drawing

import (
    "testing"
)

type MockContenter struct {
    H, W int
    R string
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
    tab.SetWidth(10)
    tab.SetHeight(10)

    if tab.Add(&MockContenter{9, 9, ""}) != nil {
        t.Error("Unexpected result")
    }

    if tab.Add(&MockContenter{9, 0, ""}) == nil {
        t.Error("Unexpected result")
    }

    if tab.Add(&MockContenter{0, 9, ""}) == nil {
        t.Error("Unexpected result")
    }
}