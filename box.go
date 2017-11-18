package box_drawing

import (
	"strings"
)

type Boxer interface {
	Height() int
	Weight() int
	String() string
}

type Box struct {
	height, weight uint
	lines          [][]rune
}

func NewBoxByRune(r rune, h uint, w uint) *Box {
	lines := make([][]rune, h)
	for i := uint(0); i < h; i++ {
		l := strings.Repeat(string(r), int(w))
		lines[i] = []rune(l)
	}
	return &Box{h, w, lines}
}

func NewBox(sl []string) *Box {
	w := 0
	h := len(sl)
	lines := make([][]rune, h)
	for i, line := range sl {
		cnt := len(line)
		if w < cnt {
			w = cnt
		}
		lines[i] = []rune(line)
	}
	return &Box{uint(h), uint(w), lines}
}

func (b *Box) Width() int {
	if b.height > 0 {
		return int(b.weight)
	}
	return 0
}

func (b *Box) Height() int {
	if b.weight > 0 {
		return int(b.height)
	}
	return 0
}

func (b *Box) min(l uint, r uint) uint {
	if l < r {
		return l
	}
	return r
}

func (b *Box) Merge(n *Box, x uint, y uint) {
	lines := n.Lines()
	mh := b.min(n.height+x, uint(b.Height()))
	mw := b.min(n.weight+y, uint(b.Width()))
	for i := x; i < mh; i++ {
		for j := y; j < mw; j++ {
			b.lines[i][j] = lines[i-x][j-y]
		}
	}
}

func (b *Box) SetPoint(r rune, x uint, y uint) {
	if x < uint(b.Width()) && y < uint(b.Height()) {
		b.lines[y][x] = r
	}
}

func (b *Box) SetRow(r rune, x uint, y uint, n uint) {
	if x < uint(b.Width()) && y < uint(b.Height()) {
		mw := b.min(x+n, uint(b.Width()))
		for i := x; i < mw; i++ {
			b.lines[y][i] = r
		}
	}
}

func (b *Box) SetColumn(r rune, x uint, y uint, n uint) {
	if x < uint(b.Width()) && y < uint(b.Height()) {
		mh := b.min(y+n, uint(b.Height()))
		for j := x; j < mh; j++ {
			b.lines[j][x] = r
		}
	}
}

func (b *Box) String() string {
	lines := make([]string, b.Height())
	for i, line := range b.lines {
		lines[i] = string(line)
	}
	return strings.Join(lines, "\n")
}

func (b *Box) Lines() [][]rune {
	return b.lines
}
