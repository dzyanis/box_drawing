package box_drawing

type Contenter interface {
	Box() *Box
}

type EmptyContent struct{}

func (e EmptyContent) Height() int     { return 0 }
func (e EmptyContent) Width() int      { return 0 }
func (e EmptyContent) String() string  { return "" }
func (e EmptyContent) Lines() []string { return []string{} }
