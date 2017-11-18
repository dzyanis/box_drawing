package box_drawing

import (
	"strings"
	"unicode/utf8"
)

type Title string

func NewTitle(text string) *Title {
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\t", " ", -1)
	text = strings.Title(text)
	t := Title(text)
	return &t
}

func (t Title) String() string {
	return string(t)
}

func (t Title) Box() *Box {
	w := utf8.RuneCountInString(string(t))

	line := make([]rune, w)
	copy(line, []rune(string(t)))

	return NewBox([]string{string(line)})
}
