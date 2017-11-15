package box_drawing

const (
	SideTop int = iota
	AngleTopRight
	SideRight
	AngleRightBottom
	SideBottom
	AngleBottomLeft
	SideLeft
	AngleLeftTop
	SpliterTop
	Spliter
	SpliterBottom
)

type BoxLines [8]rune

type SequenceLines [11]rune
