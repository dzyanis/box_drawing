package box_drawing

import (
	"strings"
	"unicode/utf8"
)

type Text string

func NewText(text string) Text {
	text = strings.Replace(text, "\t", "    ", -1)
	return Text(text)
}

func (t Text) String() string {
	return string(t)
}

func (t Text) Box() *Box {
	var h, w int
	lines := strings.Split(string(t), "\n")
	for _, line := range lines {
		cnt := utf8.RuneCountInString(line)
		if w < cnt {
			w = cnt
		}
		h++
	}
	//lines := make([]string, h)
	for h, src := range lines {
		line := []rune(strings.Repeat(" ", w))
		copy(line, []rune(src))
		lines[h] = string(line)
	}
	return NewBox(lines)
}
