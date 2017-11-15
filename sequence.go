package box_drawing

import (
	"strings"
)

type Sequence struct {
	children []Boxer
	sides    SequenceLines
}

type SequenceDecorator func(b *Sequence)

func NewSequence(args ...interface{}) *Sequence {
	s := &Sequence{}
	for _, v := range args {
		switch v.(type) {
		case Boxer:
			s.children = append(s.children, v.(Boxer))
		case SequenceLines:
			s.sides = v.(SequenceLines)
		case SequenceDecorator:
			v.(SequenceDecorator)(s)
		}
	}
	return s
}

func (b *Sequence) firstLine() string {
	line := string(b.sides[AngleLeftTop])
	cnt := len(b.children)
	for i := 0; i < cnt; i++ {
		line += strings.Repeat(string(b.sides[SideTop]), b.children[i].Width())

		if i+1 >= cnt {
			line += string(b.sides[AngleTopRight])
		} else {
			line += string(b.sides[SpliterTop])
		}
	}
	return line
}

func (b *Sequence) lastLine() string {
	line := string(b.sides[AngleBottomLeft])
	cnt := len(b.children)
	for i := 0; i < cnt; i++ {
		line += strings.Repeat(string(b.sides[SideBottom]), b.children[i].Width())

		if i+1 >= cnt {
			line += string(b.sides[AngleRightBottom])
		} else {
			line += string(b.sides[SpliterBottom])
		}
	}
	return line
}

func (b *Sequence) Draw() string {
	box := b.firstLine() + "\n"

	cnt := len(b.children)
	for i := 0; i < cnt; i++ {
		content := b.children[i].Draw()
		if i <= 0 {
			box += string(b.sides[SideLeft])
		}

		for _, line := range strings.Split(content, "\n") {
			box += line
		}

		if i+1 >= cnt {
			box += string(b.sides[SideRight])
		} else {
			box += string(b.sides[Spliter])
		}
	}

	box += "\n" + b.lastLine()
	return box
}

func (b *Sequence) String() string {
	return b.Draw()
}

func (b *Sequence) internalHeight() int {
	l := 0
	for _, ch := range b.children {
		l += ch.Height() + 1
	}
	return l + 2
}

func (b *Sequence) Height() int {
	l := 1
	for _, ch := range b.children {
		l += ch.Height() + 1
	}
	return l
}

func (b *Sequence) internalWidth() int {
	l := 0
	for _, ch := range b.children {
		l += ch.Width()
	}
	return l + 2
}

func (b *Sequence) Width() int {
	l := 0
	for _, ch := range b.children {
		l += ch.Width()
	}
	return l + 2
}
