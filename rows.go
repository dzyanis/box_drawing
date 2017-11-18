package box_drawing

import (
	"strings"
)

type Rower interface {
	Height() int
	Width() int
}

type Row struct {
	cells       []Celler
	border      *Border
	innerHeight *int
	innerWidth  *int
}

func NewRow(args ...interface{}) (*Row, error) {
	r := &Row{}
	for _, v := range args {
		switch v.(type) {
		case Celler:
			err := r.Add(v.(Celler))
			if err != nil {
				return r, nil
			}
		case *Border:
			r.SetBorder(v.(*Border))
		}
	}

	return r, nil
}

func (r *Row) Add(cell Celler) error {
	if r.innerHeight != nil && cell.Height()+r.cellsHeight() > *r.innerHeight {
		return ErrRowNotFitHeight
	}

	if r.innerWidth != nil && cell.Width()+r.cellsWidth() > *r.innerWidth {
		return ErrRowNotFitWidth
	}

	r.cells = append(r.cells, cell)
	return nil
}

func (r *Row) Height() int {
	if r.innerHeight != nil {
		return *r.innerHeight
	}

	height := r.cellsHeight()
	return height
}

func (r *Row) cellsHeight() int {
	height := 0
	for _, r := range r.cells {
		if height < r.Width() {
			height = r.Width()
		}
	}
	return height
}

func (r *Row) SetInnerHeight(height int) {
	r.innerHeight = &height
}

func (r *Row) Width() int {
	if r.innerWidth != nil {
		return *r.innerWidth
	}

	width := r.cellsWidth()
	return width
}

func (r *Row) cellsWidth() int {
	width := 0
	for _, r := range r.cells {
		if width < r.Width() {
			width = r.Width()
		}
	}
	return width
}

func (r *Row) SetInnerWidth(width int) {
	r.innerWidth = &width
}

func (r *Row) SetBorder(border *Border) {
	r.border = border
}

func (r *Row) firstLine() string {
	s := r.border.Element(AngleLeftTop)
	cnt := len(r.cells)
	for _, cell := range r.cells {
		s += strings.Repeat(r.border.Element(SideTop), cell.Width())
		if cnt < 1 {
			s += r.border.Element(AngleTopRight)
		} else {
			s += r.border.Element(SeparatorTop)
		}
		cnt--
	}
	return s
}

func (r *Row) lastLine() string {
	s := r.border.Element(AngleBottomLeft)
	cnt := len(r.cells)
	for _, cell := range r.cells {
		s += strings.Repeat(r.border.Element(SideBottom), cell.Width())
		if cnt < 1 {
			s += r.border.Element(AngleRightBottom)
		} else {
			s += r.border.Element(SeparatorBottom)
		}
		cnt--
	}
	return s
}

//func (r *Row) Draw() int {
//	text := r.firstLine()
//	for _, cell := range r.cells {
//		for _, line := range cell.Boxes() {
//			text += r.border.Element(SideLeft)
//			text += line
//			text += r.border.Element(SideRight)
//		}
//	}
//	text += r.lastLine()
//	return width
//}
