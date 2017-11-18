package box_drawing

import (
	"fmt"
	"strings"
)

type Celler interface {
	Height() int
	Width() int
	Draw() string
}

type Cell struct {
	content     *Box
	border      *Border
	innerHeight *int
	innerWidth  *int
}

func NewCellEmpty() *Cell {
	return &Cell{
		content: &Box{},
		border:  NewBorderEmpty(),
	}
}

func NewCell(args ...interface{}) (*Cell, error) {
	c := NewCellEmpty()
	for _, v := range args {
		switch v.(type) {
		case *Box:
			err := c.Add(v.(*Box))
			if err != nil {
				return c, err
			}
		case *Border:
			c.SetBorder(v.(*Border))
		}
	}
	return c, nil
}

func (c *Cell) Add(b *Box) error {
	if c.innerHeight != nil && b.Height()+c.contentHeight() > *c.innerHeight {
		return ErrRowNotFitHeight
	}

	if c.innerWidth != nil && b.Width()+c.contentWidth() > *c.innerWidth {
		return ErrRowNotFitWidth
	}

	c.content = b
	return nil
}

func (c *Cell) firstLine() string {
	return fmt.Sprintf(
		"%s%s%s",
		c.border.Element(AngleLeftTop),
		strings.Repeat(string(c.border.Element(SideTop)), c.content.Width()),
		c.border.Element(AngleTopRight),
	)
}

func (c *Cell) lastLine() string {
	return fmt.Sprintf(
		"%s%s%s",
		c.border.Element(AngleBottomLeft),
		strings.Repeat(string(c.border.Element(SideBottom)), c.content.Width()),
		c.border.Element(AngleRightBottom),
	)
}

func (c *Cell) Draw() string {
	box := c.firstLine() + "\n"

	for _, line := range c.content.Lines() {
		box += c.border.Element(SideLeft)
		box += string(line)
		box += c.border.Element(SideRight) + "\n"
	}

	box += c.lastLine()

	nb := NewBoxByRune(' ', uint(c.Height()), uint(c.Width()))
	nb.Merge(c.content, uint(c.border.Height(0))/2, uint(c.border.Width(0))/2)

	return box
}

func (c *Cell) SetBorder(border *Border) {
	c.border = border
}

func (c *Cell) Height() int {
	if c.innerHeight != nil {
		return *c.innerHeight
	}

	height := c.contentHeight()
	return height
}

func (c *Cell) contentHeight() int {
	h := 0
	if c.content != nil {
		h += c.content.Height()
	}
	h += c.border.Height(0)
	return h
}

func (c *Cell) SetInnerHeight(height int) {
	c.innerHeight = &height
}

func (c *Cell) Width() int {
	if c.innerWidth != nil {
		return *c.innerWidth
	}

	width := c.contentWidth()
	return width
}

func (c *Cell) contentWidth() int {
	w := 0
	if c.content != nil {
		w += c.content.Width()
	}
	w += c.border.Width(0)
	return w
}

func (c *Cell) SetInnerWidth(width int) {
	c.innerWidth = &width
}
