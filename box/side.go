package box

type Lines [8]rune

const (
	SideTop int = iota
	AngleTopRight
	SideRight
	AngleRightBottom
	SideBottom
	AngleBottomLeft
	SideLeft
	AngleLeftTop
)
