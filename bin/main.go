package main

import (
	"fmt"
	bd "github.com/dzyanis/box_drawing"
)

func main() {
	title := bd.NewTitle("test")
	fmt.Println(title)

	a1, _ := bd.NewCell(bd.NewTitle("A1"), bd.NewBorderBoxDrawing())
	fmt.Println(a1.Draw())
	a2, _ := bd.NewCell(bd.NewTitle("A2"), bd.NewBorderBoxDrawing())
	fmt.Println(a2.Draw())
	a3, _ := bd.NewCell(bd.NewTitle("A3"), bd.NewBorderBoxDrawing())
	fmt.Println(a2.Draw())
	row1, _ := bd.NewRow(a1, a2, a3)
	fmt.Println(row1)
	//cell2, _ := bd.NewRow(bd.NewCell(bd.NewTitle("B1")), bd.NewCell(bd.NewTitle("B2")), bd.NewCell(bd.NewTitle("B3")))
	//cell3, _ := bd.NewRow(bd.NewCell(bd.NewTitle("C1")), bd.NewCell(bd.NewTitle("C2")), bd.NewCell(bd.NewTitle("C3")))
	//tbl, _ := bd.NewTable(cell1, cell2, cell3, bd.NewBorderBoxDrawing())
	//fmt.Println(tbl)
}
