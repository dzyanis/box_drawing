package box_drawing

import "errors"

var (
	ErrRowNotFitHeight = errors.New("Row doesn't fit by the height")
	ErrRowNotFitWidth  = errors.New("Row doesn't fit by the width")
)

type Tabler interface {
	Height() int
	Width() int
}

type Table struct {
	rows   []Rower
	height *int
	width  *int
}

type TableDecorator func(tab *Table) error

func NewDecorator(f TableDecorator) TableDecorator {
	return f
}

func NewTable(args ...interface{}) (*Table, error) {
	t := &Table{}
	for _, v := range args {
		switch v.(type) {
		case Rower:
			err := t.Add(v.(Rower))
			if err != nil {
				return t, err
			}
		case TableDecorator:
			err := v.(TableDecorator)(t)
			if err != nil {
				return t, err
			}
		}
	}
	return t, nil
}

func (t *Table) Add(row Rower) error {
	if t.height != nil && row.Height()+t.rowsHeight() > *t.height {
		return ErrRowNotFitHeight
	}

	if t.width != nil && row.Width()+t.rowsWidth() > *t.width {
		return ErrRowNotFitWidth
	}

	t.rows = append(t.rows, row)
	return nil
}

func (t *Table) Height() int {
	if t.height != nil {
		return *t.height
	}

	height := t.rowsHeight()
	return height
}

func (t *Table) rowsHeight() int {
	height := 0
	for _, r := range t.rows {
		height += r.Height()
	}
	return height
}

func (t *Table) SetHeight(height int) {
	t.height = &height
}

func (t *Table) Width() int {
	if t.width != nil {
		return *t.width
	}

	width := t.rowsWidth()
	return width
}

func (t *Table) rowsWidth() int {
	width := 0
	for _, r := range t.rows {
		if width < r.Width() {
			width = r.Width()
		}
	}
	return width
}

func (t *Table) SetWidth(width int) {
	t.width = &width
}

func (t *Table) Draw() string {
	return ""
}
