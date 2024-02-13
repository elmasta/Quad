package piscine

import "github.com/01-edu/z01"

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

func printW(n int, a rune, b rune, c rune) {
	/* fait par Basile
	quad d l'utilise */
	for i := 0; i < n; i++ {
		if i == 0 {
			z01.PrintRune(a)
		} else if i+1 == n {
			z01.PrintRune(c)
		} else {
			z01.PrintRune(b)
		}
	}
	z01.PrintRune('\n')
}

func printH(x int, y int, c rune) {
	/* fait par Basile
	quad e l'utilise */
	for i := 0; i < y-2; i++ {
		z01.PrintRune(c)
		for j := 0; j < x-2; j++ {
			z01.PrintRune(' ')
			if j == x-3 {
				z01.PrintRune(c)
			}
		}
		z01.PrintRune('\n')
	}
}

func QuadA(x,y int) {
	/* fait par Nikolozi */
	/* gestion si x ou y est supérieur à 0
	si supérieur, ne rien faire*/
	if y >= 1 && x >= 1 {
		/* Boucle comptage de ligne*/
		for line := 0; line < y; line++ {
			/* Gestion première et dernière ligne car elles sont différentes des lignes du milieux*/
			if line == 0 || line+1 == y {
				/* Comptage de caractères sur une ligne */
				for column := 0; column < x; column++ {
					/* si début ou fin de ligne */
					if column == 0 || column+1 == x {
						z01.PrintRune('o')
					} else {
						z01.PrintRune('-')
					}
				}
			} else {
				for column := 0; column < x; column++ {
					if column == 0 || column+1 == x {
						z01.PrintRune('|')
					} else {
						z01.PrintRune(' ')
					}
				}
			}
			z01.PrintRune('\n')
		}			
	}
}

func QuadB(x,y int) {
	/* fait par Valentin */
	a := '/'
	b := '*'
	c := '\\'
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

func QuadC(x,y int) {
	/* fait par Valentin */
	a := 'A'
	b := 'B'
	c := 'C'
	d := ' '
	if y >= 1 && x >= 1 {
		for i := 0; i < y; i++ {
			if i == 0 {
				printChar(x, a, b, a)
			} else if i+1 == y {
				printChar(x, c, b, c)
			} else {
				printChar(x, b, d, b)
			}
		}
	}
}

func QuadD(x, y int) {
	/* fait par Basile */
	a := 'A'
	b := 'B'
	c := 'C'
	if x == 1 && y == 1 {
		z01.PrintRune(a)
		z01.PrintRune('\n')
	} else if x > 1 && y == 1 {
		printW(x, a, b, c)
	} else if x == 1 && y > 1 {
		z01.PrintRune(a)
		z01.PrintRune('\n')
		printH(x, y, b)
		z01.PrintRune(a)
		z01.PrintRune('\n')
	} else if y > 1 && x > 1 {
		printW(x, a, b, c)
		printH(x, y, b)
		printW(x, a, b, c)
	}
}

func QuadE(x, y int) {
	/* fait par Basile */
	a := 'A'
	b := 'B'
	c := 'C'
	if x == 1 && y == 1 {
		z01.PrintRune(a)
		z01.PrintRune('\n')
	} else if x > 1 && y == 1 {
		printW(x, a, b, c)
	} else if x == 1 && y > 1 {
		z01.PrintRune(a)
		z01.PrintRune('\n')
		printH(x, y, b)
		z01.PrintRune(c)
		z01.PrintRune('\n')
	} else if y > 1 && x > 1 {
		printW(x, a, b, c)
		printH(x, y, b)
		printW(x, c, b, a)
	}
}
