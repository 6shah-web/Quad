package piscine

import "github.com/01-edu/z01"

func QuadC(x, y int) {
	for h := 1; h <= y; h++ {
		for w := 1; w <= x; w++ {
			if (h == 1 || h == y) && (w == 1 || w == x) {
				if h == 1 && w == 1 {
					z01.PrintRune('A')
				} else if h == 1 && w == x {
					z01.PrintRune('A')
				} else if h == y && w == 1 {
					z01.PrintRune('C')
				} else if h == y && w == x {
					z01.PrintRune('C')
				}
			} else if (h == 1 && w > 1 && w < x) || (h == y && w > 1 && w < x) {
				z01.PrintRune('B')
			} else if w == 1 || w == x {
				z01.PrintRune('B')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
