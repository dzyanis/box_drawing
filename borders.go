package box_drawing

type BorderElement int

const (
	SideTop BorderElement = iota
	AngleTopRight
	SideRight
	AngleRightBottom
	SideBottom
	AngleBottomLeft
	SideLeft
	AngleLeftTop

	SeparatorTop
	SeparatorVertical
	SeparatorBottom

	SeparatorRight
	SeparatorLeft
	SeparatorCross
	SeparatorHorizontal
)

type Element *rune

func RuneToElement(r rune) Element {
	return Element(&r)
}

type BorderElements [SeparatorHorizontal + 1]*rune

type Border struct {
	elements BorderElements
	height   int
	width    int
}

func NewBorder(elements BorderElements) *Border {
	return &Border{
		elements: elements,
	}
}

func NewBorderBoxDrawing() *Border {
	elements := BorderElements{
		SideTop:             RuneToElement('═'),
		AngleTopRight:       RuneToElement('╗'),
		SideRight:           RuneToElement('║'),
		AngleRightBottom:    RuneToElement('╝'),
		SideBottom:          RuneToElement('═'),
		AngleBottomLeft:     RuneToElement('╚'),
		SideLeft:            RuneToElement('║'),
		AngleLeftTop:        RuneToElement('╔'),
		SeparatorTop:        RuneToElement('╦'),
		SeparatorVertical:   RuneToElement('║'),
		SeparatorBottom:     RuneToElement('╩'),
		SeparatorRight:      RuneToElement('╣'),
		SeparatorLeft:       RuneToElement('╠'),
		SeparatorCross:      RuneToElement('╬'),
		SeparatorHorizontal: RuneToElement('═'),
	}
	return NewBorder(elements)
}

func (b Border) Element(element BorderElement) string {
	if b.elements[element] != nil {
		return string(*b.elements[element])
	}
	return ""
}
