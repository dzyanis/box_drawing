package box

import "strings"

type Text struct {
	cnt    string
	height int
	width  int
}

func NewText(text string) *Text {
	t := &Text{cnt: text}

	for _, line := range strings.Split(text, "\n") {
		if t.width < len(line) {
			t.height += 1
			t.width = len(line)
		}
	}

	return t
}

func (t Text) Draw() string {
	return t.cnt
}

func (t Text) Height() int {
	return t.height
}

func (t Text) Width() int {
	return t.width
}
