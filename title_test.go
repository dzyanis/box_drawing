package box_drawing

import (
	"testing"
)

func TestNewTitle(t *testing.T) {
	title := NewTitle("")

	if title == nil {
		t.Error("Unexpected result")
	}
}

func TestTitleString(t *testing.T) {
	title := NewTitle("some text")

	if title.String() != "Some Text" {
		t.Error("Unexpected result")
	}
}

func TestTitleHeightAndWidth(t *testing.T) {
	title := NewTitle("some text\n\thello")

	if title.Box().Height() != 1 {
		t.Error("Unexpected result")
	}

	if title.Box().Width() != 15 {
		t.Errorf("Unexpected result: `%s` len:%d", title.String(), title.Box().Width())
	}
}
