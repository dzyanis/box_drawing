package box_drawing

type Rower interface {
	Height() int
	Width() int
}

type Row struct {
	cells  []Celler
	border *Border
	height *int
	width  *int
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
	if r.height != nil && cell.Height()+r.cellsHeight() > *r.height {
		return ErrRowNotFitHeight
	}

	if r.width != nil && cell.Width()+r.cellsWidth() > *r.width {
		return ErrRowNotFitWidth
	}

	r.cells = append(r.cells, cell)
	return nil
}

func (r *Row) Height() int {
	if r.height != nil {
		return *r.height
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

func (r *Row) SetHeight(height int) {
	r.height = &height
}

func (r *Row) Width() int {
	if r.width != nil {
		return *r.width
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

func (r *Row) SetWidth(width int) {
	r.width = &width
}

func (r *Row) SetBorder(border *Border) {
	r.border = border
}
