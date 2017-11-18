package box_drawing

import (
	"testing"
)

func TestNewText(t *testing.T) {
	text := NewText("")
	if text.String() != "" {
		t.Errorf("Unexpected result: %#v", text.String())
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
	if text.Box().Height() != 3 {
		t.Errorf("Unexpected result: %d", text.Box().Height())
	}

	text = NewText("")
	if text.Box().Height() != 0 {
		t.Errorf("Unexpected result: %d", text.Box().Height())
	}
}

func TestTextWidth(t *testing.T) {
	text := NewText("a\nab\nabc")
	if text.Box().Width() != 3 {
		t.Errorf("Unexpected result: %d", text.Box().Width())
	}
}

func TestTextLines(t *testing.T) {
	text := NewText("a\nab\nabc")
	lines := text.Box().Lines()

	if string(lines[0]) != "a  " {
		t.Errorf("Unexpected result: `%s`", lines[0])
	}

	if string(lines[1]) != "ab " {
		t.Errorf("Unexpected result: `%s`", lines[1])
	}

	if string(lines[2]) != "abc" {
		t.Errorf("Unexpected result: `%s`", lines[2])
	}
}

func TestTextLinesWithTab(t *testing.T) {
	text := NewText("a\nab\na\tc")
	lines := text.Box().Lines()

	if string(lines[0]) != "a     " {
		t.Errorf("Unexpected result: `%s`", lines[0])
	}

	if string(lines[1]) != "ab    " {
		t.Errorf("Unexpected result: `%s`", lines[1])
	}

	if string(lines[2]) != "a    c" {
		t.Errorf("Unexpected result: `%s`", lines[2])
	}
}
