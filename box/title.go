package box

import (
    "strings"
    "unicode/utf8"
)

type Title string

func NewTitle(text string) *Title {
	text = strings.Title(text)
	t := Title(text)
	return &t
}

func (t Title) Draw() string {
	return string(t)
}

func (t Title) Height() int {
	return 1
}

func (t Title) Width() int {
	return utf8.RuneCountInString(string(t))
}
