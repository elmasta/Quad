package main

import (
	"github.com/01-edu/z01"
	"os"
	"strconv"
)

func printChar(l int, a rune, b rune, c rune) {
	/* fait par Valentin
	quad b & c l'utilisent */
	for i := 0; i < l; i++{
		if i == 0 {
			z01.PrintRune(a)
		} else if i+1 == l {
			z01.PrintRune(c)
		} else {
			z01.PrintRune(b)
		}
	}
	z01.PrintRune('\n')
}

func QuadE(x,y int) {
	/* fait par Valentin */
	a := 'A'
	b := 'B'
	c := 'C'
	d := ' '
	if y >= 1 && x >= 1 {
		for i := 0; i < y; i++ {
			if i == 0 {
				printChar(x, a, b, c)
			} else if i+1 == y {
				printChar(x, c, b, a)
			} else {
				printChar(x, b, d, b)
			}
		}
	}
}

func main() {
	if len(os.Args) == 3 {
		x, errx := strconv.Atoi(os.Args[1])
		y, erry := strconv.Atoi(os.Args[2])
		if errx == nil && erry == nil {
			QuadE(x, y)
		}
	}
}
