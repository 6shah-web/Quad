package main

import "github.com/01-edu/z01"

func main() {
	width := 1
	height := 5

	// Header rectangle
	if width > 1 {
		z01.PrintRune('o')
	}
	for i := 1; i < width-1; i++ {
		z01.PrintRune('-')
	}
	z01.PrintRune('o')
	z01.PrintRune('\n')

	// body rectangle and side
	for k := 1; k < height-1; k++ {
		if height > 1 {
			z01.PrintRune('|')
			for j := 1; j < width-1; j++ {
				z01.PrintRune(' ')
			}
			if width > 1 {
				z01.PrintRune('|')
			}
		}
		z01.PrintRune('\n')
	}

	// footer rectangle
	if height > 1 {
		if width > 1 {
			z01.PrintRune('o')
		}
		for i := 1; i < width-1; i++ {
			z01.PrintRune('-')
		}

		z01.PrintRune('o')
		z01.PrintRune('\n')
	} else {
		return
	}
}
