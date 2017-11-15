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

func (t Title) Draw() string {
	return string(t)
}

func (t Title) String() string {
	return t.Draw()
}

func (t Title) Height() int {
	return 1
}

func (t Title) Width() int {
	return utf8.RuneCountInString(string(t))
}
