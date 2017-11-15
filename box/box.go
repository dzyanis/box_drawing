package box

import (
	"fmt"
	"strings"
)

type Boxer interface {
	Height() int
	Width() int
	Draw() string
}

type Box struct {
	parent Boxer
	sides  Lines
}

type Decorator func(b *Box)

func NewBox(args ...interface{}) *Box {
	b := &Box{}
	for _, v := range args {
		switch v.(type) {
		case Boxer:
			b.parent = v.(Boxer)
		case Lines:
			b.sides = v.(Lines)
		case Decorator:
			v.(Decorator)(b)
		}
	}
	return b
}

func (b *Box) firstLine() string {
	return fmt.Sprintf("%c%s%s", b.sides[AngleLeftTop], strings.Repeat(string(b.sides[SideTop]), b.parent.Width()), string(b.sides[AngleTopRight]))
}

func (b *Box) lastLine() string {
	return fmt.Sprintf("%c%s%c", b.sides[AngleBottomLeft], strings.Repeat(string(b.sides[SideBottom]), b.parent.Width()), b.sides[AngleRightBottom])
}

func (b *Box) Draw() string {
	box := b.firstLine() + "\n"

	if b.parent != nil {
		content := b.parent.Draw()
		for _, line := range strings.Split(content, "\n") {
			box += fmt.Sprintf("%c%s%c\n", b.sides[SideLeft], line, b.sides[SideRight])
		}
	}

	box += b.lastLine()
	return box
}

func (b *Box) String() string {
	return b.Draw()
}

func (b *Box) Height() int {
	l := 0
	if b.parent != nil {
		l = b.parent.Height()
	}
	return l + 2
}

func (b *Box) Width() int {
	l := 0
	if b.parent != nil {
		l = b.parent.Width()
	}
	return l + 2
}
