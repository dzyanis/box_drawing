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

func (t Title) Lines() []string {
	line := make([]rune, t.Width())
	copy(line, []rune(string(t)))

	return []string{string(line)}
}

func (t Title) String() string {
	return string(t)
}

func (t Title) Height() int {
	return 1
}

func (t Title) Width() int {
	return utf8.RuneCountInString(string(t))
}
