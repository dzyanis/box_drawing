package box_drawing

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
	child Boxer
	sides BoxLines
}

type BoxDecorator func(b *Box)

func NewBox(args ...interface{}) *Box {
	b := &Box{}
	for _, v := range args {
		switch v.(type) {
		case Boxer:
			b.child = v.(Boxer)
		case BoxLines:
			b.sides = v.(BoxLines)
		case BoxDecorator:
			v.(BoxDecorator)(b)
		}
	}
	return b
}

func (b *Box) firstLine() string {
	return fmt.Sprintf("%c%s%s", b.sides[AngleLeftTop], strings.Repeat(string(b.sides[SideTop]), b.child.Width()), string(b.sides[AngleTopRight]))
}

func (b *Box) lastLine() string {
	return fmt.Sprintf("%c%s%c", b.sides[AngleBottomLeft], strings.Repeat(string(b.sides[SideBottom]), b.child.Width()), b.sides[AngleRightBottom])
}

func (b *Box) Draw() string {
	box := b.firstLine() + "\n"

	if b.child != nil {
		content := b.child.Draw()
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
	if b.child != nil {
		l = b.child.Height()
	}
	return l + 2
}

func (b *Box) Width() int {
	l := 0
	if b.child != nil {
		l = b.child.Width()
	}
	return l + 2
}
