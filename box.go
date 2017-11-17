package box_drawing

import (
	"strings"
    "math"
)

type Boxer interface {
	Height() int
	Weight() int
	String() string
}

type Box struct {
	height, weight uint
	lines          []string
}

func NewBox(r rune, h uint, w uint) *Box {
	lines := make([]string, h)
	for i := uint(0); i < h; i++ {
		lines[i] = strings.Repeat(string(r), int(w))
	}
	return &Box{h, w, lines}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func (b *Box) Weight() int {
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

func (b *Box) Merge(n *Box, x uint, y uint) {
	lines := n.Lines()
	mh := int(math.Min(float64(n.height+x), float64(b.Height())))
    mw := int(math.Min(float64(n.weight+y), float64(b.Weight())))
	for i := int(x); i < mh; i++ {
		bl := []byte(b.lines[i])
		for j := int(y); j < mw; j++ {
			bl[j] = lines[i-int(x)][j-int(y)]
		}
		b.lines[i] = string(bl)
	}
}

func (b *Box) String() string {
	return strings.Join(b.lines, "\n")
}

func (b *Box) Lines() []string {
	return b.lines
}
