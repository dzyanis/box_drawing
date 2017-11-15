package main

import (
	"fmt"
	bd "github.com/dzyanis/box_drawing"
)

func main() {
	title := bd.NewTitle("test")
	fmt.Println(title)

	b := bd.NewBox(title, bd.BoxLines{'═', '╗', '║', '╝', '═', '╚', '║', '╔'})
	fmt.Println(b)

	s := bd.NewSequence(
		bd.NewText("cell 1"),
		bd.NewText("cell 2"),
		bd.NewText("cell 3"),
		bd.SequenceLines{'═', '╗', '║', '╝', '═', '╚', '║', '╔', '╦', '║', '╩'},
	)
	fmt.Println(s)
}
