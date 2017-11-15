package main

import (
	"fmt"
	"github.com/dzyanis/box_drawing/box"
)

func main() {
	title := box.NewTitle("їбланити")
	fmt.Println(title.Draw())

	bs := box.NewBox(title, box.Lines{'*', '*', '*', '*', '*', '*', '*', '*'})
	fmt.Println(bs)

	bd := box.NewBox(bs, box.Lines{'═', '╗', '║', '╝', '═', '╚', '║', '╔'})
	fmt.Println(bd)

	bpm := box.NewBox(bd, box.Lines{'-', '+', '|', '+', '-', '+', '|', '+'})
	fmt.Println(bpm)

	bb := box.NewBox(bpm, box.Lines{'▀', '▜', '▐', '▟', '▄', '▙', '▌', '▛'})
	fmt.Println(bb)
}
