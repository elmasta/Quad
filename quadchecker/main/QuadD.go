package main

import (
	"github.com/01-edu/z01"
	"os"
	"strconv"
)

func QuadD(x, y int) {
	if x <= 0 || y <= 0 {
		os.Exit(0)
	}

	for hauteur := 1; hauteur <= y; hauteur++ {
		for longueur := 1; longueur <= x; longueur++ {
			if hauteur == 1 || hauteur == y {
				if longueur == 1 || longueur == x {
					if longueur == 1 {
						z01.PrintRune('A')
					} else {
						z01.PrintRune('C')
					}
				} else {
					z01.PrintRune('B')
				}
			} else {
				if longueur == 1 || longueur == x {
					z01.PrintRune('B')
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}
}

func main() {
	if len(os.Args) == 3 {
		x, errx := strconv.Atoi(os.Args[1])
		y, erry := strconv.Atoi(os.Args[2])
		if errx == nil && erry == nil {
			QuadD(x, y)
		}
	}
}
