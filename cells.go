package box_drawing

type Celler interface {
	Height() uint
	Width() uint
	Draw() string
}

type Cell struct {
	content     Contenter
	border      *Border
	innerHeight *uint
	innerWidth  *uint
}

func NewCellEmpty() *Cell {
	return &Cell{
		border:  NewBorderEmpty(),
	}
}

func NewCell(args ...interface{}) (*Cell, error) {
	c := NewCellEmpty()
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

func (c *Cell) Add(b Contenter) error {
	if c.innerHeight != nil && b.Box().Height()+c.contentHeight() > *c.innerHeight {
		return ErrRowNotFitHeight
	}

	if c.innerWidth != nil && b.Box().Width()+c.contentWidth() > *c.innerWidth {
		return ErrRowNotFitWidth
	}

	c.content = b
	return nil
}

//func (c *Cell) firstLine() string {
//	return fmt.Sprintf(
//		"%s%s%s",
//		c.border.Element(AngleLeftTop),
//		strings.Repeat(string(c.border.Element(SideTop)), c.content.Width()),
//		c.border.Element(AngleTopRight),
//	)
//}
//
//func (c *Cell) lastLine() string {
//	return fmt.Sprintf(
//		"%s%s%s",
//		c.border.Element(AngleBottomLeft),
//		strings.Repeat(string(c.border.Element(SideBottom)), c.content.Width()),
//		c.border.Element(AngleRightBottom),
//	)
//}

func (c *Cell) Box() *Box {
	b := NewBoxByRune(' ', c.Height(), c.Width())

	xm := c.contentWidth()-1
	ym := c.contentHeight()-1

	b.SetPoint(c.border.Rune(AngleLeftTop), 0, 0)
	if c.content != nil {
		b.SetRow(c.border.Rune(SideTop), 1, 0, xm)
	}
	b.SetPoint(c.border.Rune(AngleTopRight),xm, 0)

	b.SetColumn(c.border.Rune(SideLeft), 0, 1, c.contentHeight()-c.border.Height(0))
	b.SetColumn(c.border.Rune(SideRight), xm, 1, c.contentHeight()-c.border.Height(0))

	b.SetPoint(c.border.Rune(AngleBottomLeft), 0, ym)
	if c.content != nil {
		b.SetRow(c.border.Rune(SideBottom), 1, ym, xm)
	}
	b.SetPoint(c.border.Rune(AngleRightBottom), xm, ym)

	b.Merge(c.content.Box(), 1, 1)

	return b
}

func (c *Cell) SetBorder(border *Border) {
	c.border = border
}

func (c *Cell) Height() uint {
	if c.innerHeight != nil {
		return *c.innerHeight + c.border.Height(0)
	}

	height := c.contentHeight()
	return height
}

func (c *Cell) contentHeight() uint {
	var h uint
	if c.content != nil {
		h += c.content.Box().Height()
	}
	h += c.border.Height(0)
	return h
}

func (c *Cell) SetInnerHeight(height uint) {
	c.innerHeight = &height
}

func (c *Cell) Width() uint {
	if c.innerWidth != nil {
		return *c.innerWidth + c.border.Width(0)
	}

	width := c.contentWidth()
	return width
}

func (c *Cell) contentWidth() uint {
	var w uint
	if c.content != nil {
		w += c.content.Box().Width()
	}
	w += c.border.Width(0)
	return w
}

func (c *Cell) SetInnerWidth(width uint) {
	c.innerWidth = &width
}
