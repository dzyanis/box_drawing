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

func NewBorderEmpty() *Border {
	return &Border{elements: BorderElements{}}
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

func (b Border) Element(element BorderElement) Element {
	if b.elements[element] != nil {
		return b.elements[element]
	}
	return nil
}

func (b Border) Rune(element BorderElement) rune {
	if b.elements[element] != nil {
		return rune(*b.elements[element])
	}
	return ' '
}

func (b Border) WidthTop(separators int) uint {
	var w uint
	if separators > 1 && b.elements[SeparatorVertical] != nil {
		w += uint(separators)
	}
	// todo: Probably an element can be more then 1
	if b.elements[AngleTopRight] != nil {
		w += 1
	}
	if b.elements[AngleLeftTop] != nil {
		w += 1
	}
	return w
}

func (b Border) WidthBottom(separators int) uint {
	var w uint
	if separators > 1 && b.elements[SeparatorVertical] != nil {
		w += uint(separators)
	}
	// todo: Probably an element can be more then 1
	if b.elements[AngleBottomLeft] != nil {
		w += 1
	}
	if b.elements[AngleRightBottom] != nil {
		w += 1
	}
	return w
}

func (b Border) Width(separators int) uint {
	var w uint
	if separators > 1 && b.elements[SeparatorVertical] != nil {
		w += uint(separators)
	}
	// todo: Probably an element can be more then 1
	if b.elements[SideRight] != nil {
		w += 1
	}
	if b.elements[SideLeft] != nil {
		w += 1
	}
	return w
}

func (b Border) Height(separators int) uint {
	var h uint
	if separators > 1 && b.elements[SeparatorHorizontal] != nil {
		h += uint(separators)
	}
	if b.elements[SideTop] != nil {
		h += 1
	}
	if b.elements[SideBottom] != nil {
		h += 1
	}
	return h
}
