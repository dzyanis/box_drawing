package box_drawing

import (
	"strings"
	"unicode/utf8"
)

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
			t.width = utf8.RuneCountInString(line)
		}
	}

	return t
}

func (t Text) Lines() []string {
	s := string(t.cnt)
	lines := make([]string, t.Height())

	for _, src := range strings.Split(s, "\n") {
		line := make([]rune, t.Width())
		copy(line, []rune(src))
		lines = append(lines, string(line))
	}

	return lines
}

func (t Text) String() string {
	return t.cnt
}

func (t Text) Height() int {
	return t.height
}

func (t Text) Width() int {
	return t.width
}
