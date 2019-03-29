package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
   +---+---+
   | 10| 20|
   +---+---+
   |400|
   +---+
*/

func main() {
	x := []int{1, 10, 999, 9999999, 7, 6, 15, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	table(x, 2)
}

func cellWidth(s []int) int {
	var cellWidth int
	for _, i := range s {
		if i > cellWidth {
			cellWidth = i
		}
	}

	return len(fmt.Sprintf("%d", cellWidth))
}

func table(s []int, c int) {
	cw := cellWidth(s)
	rowArt := strings.Repeat("-", cw)
	fmt.Print("+")
	var columnCount int
	for range s {
		if columnCount == c {
			break
		}
		columnCount++
		fmt.Printf("%s+", rowArt)
	}
	columnCount = 0
	fmt.Print("\n|")
	for _, i := range s {
		if columnCount == c {
			columnCount = 0
			fmt.Print("\n+")
			for j := 0; j < c; j++ {
				fmt.Printf("%s+", rowArt)
			}
			fmt.Print("\n|")
		}
		columnCount++
		fmt.Printf("%"+strconv.Itoa(cw)+"d|", i)
	}
	fmt.Printf("\n+")

	if len(s)%c == 0 && c != 1 {
		for range s {
			fmt.Printf("%s+", rowArt)
		}
	} else if c == 1 {
		fmt.Printf("%s+", rowArt)
	} else {
		for i := 0; i < len(s)%c; i++ {
			fmt.Printf("%s+", rowArt)
		}
	}
	fmt.Print("\n")
}
