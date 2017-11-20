package box_drawing

import (
	"strings"
)

type Rower interface {
	Height() uint
	Width() uint
}

type Row struct {
	cells       []Celler
	border      *Border
	innerHeight *uint
	innerWidth  *uint
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

func (r *Row) Height() uint {
	if r.innerHeight != nil {
		return *r.innerHeight
	}

	height := r.cellsHeight()
	return height
}

func (r *Row) cellsHeight() uint {
	var height uint
	for _, r := range r.cells {
		if height < r.Width() {
			height = r.Width()
		}
	}
	return height
}

func (r *Row) SetInnerHeight(height uint) {
	r.innerHeight = &height
}

func (r *Row) Width() uint {
	if r.innerWidth != nil {
		return *r.innerWidth
	}

	width := r.cellsWidth()
	return width
}

func (r *Row) cellsWidth() uint {
	var width uint
	for _, r := range r.cells {
		if width < r.Width() {
			width = r.Width()
		}
	}
	return width
}

func (r *Row) SetInnerWidth(width uint) {
	r.innerWidth = &width
}

func (r *Row) SetBorder(border *Border) {
	r.border = border
}

func (r *Row) firstLine() string {
	s := string(r.border.Rune(AngleLeftTop))
	cnt := len(r.cells)
	for _, cell := range r.cells {
		s += strings.Repeat(string(r.border.Rune(SideTop)), int(cell.Width()))
		if cnt < 1 {
			s += string(r.border.Rune(AngleTopRight))
		} else {
			s += string(r.border.Rune(SeparatorTop))
		}
		cnt--
	}
	return s
}

func (r *Row) lastLine() string {
	s := string(r.border.Rune(AngleBottomLeft))
	cnt := len(r.cells)
	for _, cell := range r.cells {
		s += strings.Repeat(string(r.border.Rune(SideBottom)), int(cell.Width()))
		if cnt < 1 {
			s += string(r.border.Rune(AngleRightBottom))
		} else {
			s += string(r.border.Rune(SeparatorBottom))
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
