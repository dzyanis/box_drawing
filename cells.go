package box_drawing

import (
	"fmt"
	"strings"
)

type Contenter interface {
	Height() int
	Width() int
	String() string
	Lines() []string
}

type Celler interface {
	Height() int
	Width() int
	Draw() string
}

type Cell struct {
	content Contenter
	border  *Border
	height  *int
	width   *int
}

func NewCell(args ...interface{}) (*Cell, error) {
	c := &Cell{}
	for _, v := range args {
		switch v.(type) {
		case Contenter:
			err := c.Add(v.(Contenter))
			if err != nil {
				return c, err
			}
		case *Border:
			c.SetBorder(v.(*Border))
		}
	}
	return c, nil
}

func (c *Cell) Add(cont Contenter) error {
	if c.height != nil && cont.Height()+c.contentHeight() > *c.height {
		return ErrRowNotFitHeight
	}

	if c.width != nil && cont.Width()+c.contentWidth() > *c.width {
		return ErrRowNotFitWidth
	}

	c.content = cont
	return nil
}

func (c *Cell) firstLine() string {
	return fmt.Sprintf(
		"%c%s%s",
		string(c.border.Element(AngleLeftTop)),
		strings.Repeat(string(c.border.Element(SideTop)), c.content.Width()),
		string(c.border.Element(AngleTopRight)),
	)
}

func (c *Cell) lastLine() string {
	return fmt.Sprintf(
		"%c%s%s",
		string(c.border.Element(AngleBottomLeft)),
		strings.Repeat(string(c.border.Element(SideBottom)), c.content.Width()),
		string(c.border.Element(AngleRightBottom)),
	)
}

func (c *Cell) Draw() string {
	box := c.firstLine() + "\n"

	lines := c.content.Lines()
	for h := 0; h < c.contentHeight(); h++ {
		box += string(c.border.Element(SideLeft))
		for w := 0; w < c.contentWidth(); w++ {
			box += string(lines[h][w])
		}
		box += string(c.border.Element(SideRight)) + "\n"
	}

	box += c.lastLine()
	return box
}

func (c Cell) SetBorder(border *Border) {
	c.border = border
}

func (c *Cell) Height() int {
	if c.height != nil {
		return *c.height
	}

	height := c.contentHeight()
	return height
}

func (c *Cell) contentHeight() int {
	if c.content != nil {
		return c.content.Height()
	}
	return 0
}

func (c *Cell) SetHeight(height int) {
	c.height = &height
}

func (c *Cell) Width() int {
	if c.width != nil {
		return *c.width
	}

	width := c.contentWidth()
	return width
}

func (c *Cell) contentWidth() int {
	if c.content != nil {
		return c.content.Width()
	}
	return 0
}

func (c *Cell) SetWidth(width int) {
	c.width = &width
}
