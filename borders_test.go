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
	if bord.Element(SideTop) != "═" {
		t.Error("Unexpected result")
	}
}
