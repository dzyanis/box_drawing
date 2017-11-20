package box_drawing

import "errors"

var (
	ErrRowNotFitHeight = errors.New("Row doesn't fit by the innerHeight")
	ErrRowNotFitWidth  = errors.New("Row doesn't fit by the innerWidth")
)

type Tabler interface {
	Height() uint
	Width() uint
}

type Table struct {
	rows   []Rower
	height *uint
	width  *uint
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

func (t *Table) Height() uint {
	if t.height != nil {
		return *t.height
	}

	height := t.rowsHeight()
	return height
}

func (t *Table) rowsHeight() uint {
	var height uint
	for _, r := range t.rows {
		height += r.Height()
	}
	return height
}

func (t *Table) SetHeight(height uint) {
	t.height = &height
}

func (t *Table) Width() uint {
	if t.width != nil {
		return *t.width
	}

	width := t.rowsWidth()
	return width
}

func (t *Table) rowsWidth() uint {
	var width uint
	for _, r := range t.rows {
		if width < r.Width() {
			width = r.Width()
		}
	}
	return width
}

func (t *Table) SetWidth(width uint) {
	t.width = &width
}

func (t *Table) Draw() string {
	return ""
}
