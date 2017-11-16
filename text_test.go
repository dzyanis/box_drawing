package box_drawing

import (
	"testing"
)

func TestNewText(t *testing.T) {
	text := NewText("")
	if text == nil {
		t.Error("Unexpected result")
	}
}

func TestTextString(t *testing.T) {
	text := NewText("some text")
	if text.String() != "some text" {
		t.Error("Unexpected result")
	}
}

func TestTextHeight(t *testing.T) {
	text := NewText("a\nab\nabc")
	if text.Height() != 3 {
		t.Errorf("Unexpected result: %d", text.Height())
	}

	text = NewText("")
	if text.Height() != 1 {
		t.Errorf("Unexpected result: %d", text.Height())
	}
}

func TestTextWidth(t *testing.T) {
	text := NewText("a\nab\nabc")
	if text.Width() != 3 {
		t.Errorf("Unexpected result: %d", text.Width())
	}
}

func TestTextLines(t *testing.T) {
	text := NewText("a\nab\nabc")
	lines := text.Lines()

	if lines[0] != "a  " {
		t.Errorf("Unexpected result: `%s`", lines[0])
	}

	if lines[1] != "ab " {
		t.Errorf("Unexpected result: `%s`", lines[1])
	}

	if lines[2] != "abc" {
		t.Errorf("Unexpected result: `%s`", lines[2])
	}
}

func TestTextLinesWithTab(t *testing.T) {
	text := NewText("a\nab\na\tc")
	lines := text.Lines()

	if lines[0] != "a     " {
		t.Errorf("Unexpected result: `%s`", lines[0])
	}

	if lines[1] != "ab    " {
		t.Errorf("Unexpected result: `%s`", lines[1])
	}

	if lines[2] != "a    c" {
		t.Errorf("Unexpected result: `%s`", lines[2])
	}
}