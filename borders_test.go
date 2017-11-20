package box_drawing

import (
	"testing"
)

func TestNewBorder(t *testing.T) {
	bord := NewBorderEmpty()

	if bord == nil {
		t.Error("Unexpected result")
	}
}

func TestNewBorderBoxDrawing(t *testing.T) {
	bord := NewBorderBoxDrawing()
	if bord.Rune(SideTop) != '═' {
		t.Error("Unexpected result")
	}
}

func TestBorderWidth(t *testing.T) {
	bord := NewBorderBoxDrawing()
	if bord.Width(0) != 2 {
		t.Error("Unexpected result")
	}
	if bord.Width(40) != 42 {
		t.Error("Unexpected result")
	}

	bord = NewBorderEmpty()
	if bord.Width(0) != 0 {
		t.Errorf("Unexpected result")
	}
	if bord.Width(40) != 0 {
		t.Errorf("Unexpected result")
	}
}

func TestBorderHeight(t *testing.T) {
	bord := NewBorderBoxDrawing()
	if bord.Height(0) != 2 {
		t.Error("Unexpected result")
	}
	if bord.Height(40) != 42 {
		t.Error("Unexpected result")
	}

	bord = NewBorderEmpty()
	if bord.Height(0) != 0 {
		t.Errorf("Unexpected result")
	}
	if bord.Height(40) != 0 {
		t.Errorf("Unexpected result")
	}
}
