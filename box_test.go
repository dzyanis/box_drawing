package box_drawing

import (
	"testing"
)

func TestNewBox(t *testing.T) {
	box := NewBox('0', 1, 1)
	if box.String() != "0" {
		t.Errorf("Unexpected result: %#v", box.String())
	}

    box = NewBox('0', 3, 1)
    if box.String() != "0\n0\n0" {
        t.Errorf("Unexpected result: %#v", box.String())
    }

    box = NewBox('0', 1, 3)
    if box.String() != "000" {
        t.Errorf("Unexpected result: %#v", box.String())
    }

    box = NewBox('0', 0, 3)
    if box.String() != "" {
        t.Errorf("Unexpected result: %#v", box.String())
    }

    box = NewBox(' ', 0, 0)
    if box.String() != "" {
        t.Errorf("Unexpected result: %#v", box.String())
    }
}

func TestBoxHeight(t *testing.T) {
    box := NewBox('0', 0, 1)
    if box.Height() != 0 {
        t.Errorf("Unexpected result: %#v", box.Height())
    }

    box = NewBox('0', 42, 1)
    if box.Height() != 42 {
        t.Errorf("Unexpected result: %#v", box.Height())
    }

    box = NewBox('0', 42, 0)
    if box.Height() != 0 {
        t.Errorf("Unexpected result: %#v", box.Height())
    }
}

func TestBoxWeight(t *testing.T) {
    box := NewBox('0', 0, 0)
    if box.Weight() != 0 {
        t.Errorf("Unexpected result: %#v", box.Weight())
    }

    box = NewBox('0', 1, 42)
    if box.Weight() != 42 {
        t.Errorf("Unexpected result: %#v", box.Weight())
    }

    box = NewBox('0', 0, 1)
    if box.Weight() != 0 {
        t.Errorf("Unexpected result: %#v", box.Weight())
    }
}

func TestBoxMerge(t *testing.T) {
    b1 := NewBox('0', 3, 3)
    b1.Merge(NewBox('1', 1, 1), 1, 1)
    if b1.String() != "000\n010\n000" {
        t.Errorf("Unexpected result: %#v", b1.String())
    }

    b1.Merge(NewBox('1', 1, 1), 0, 0)
    if b1.String() != "100\n010\n000" {
        t.Errorf("Unexpected result: %#v", b1.String())
    }

    b1.Merge(NewBox('1', 3, 3), 2, 2)
    if b1.String() != "100\n010\n001" {
        t.Errorf("Unexpected result: %#v", b1.String())
    }

    b1.Merge(NewBox('1', 3, 3), 0, 0)
    if b1.String() != "111\n111\n111" {
        t.Errorf("Unexpected result: %#v", b1.String())
    }

    b1.Merge(NewBox('0', 0, 0), 0, 0)
    if b1.String() != "111\n111\n111" {
        t.Errorf("Unexpected result: %#v", b1.String())
    }

    b1.Merge(NewBox('0', 5, 5), 0, 0)
    if b1.String() != "000\n000\n000" {
        t.Errorf("Unexpected result: %#v", b1.String())
    }
}

