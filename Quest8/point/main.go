package main

import "github.com/01-edu/z01"

type point struct {
	x, y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func RsmNumber(n int) {
	s := '0'
	if n/10 > 0 {
		RsmNumber(n / 10)
	}
	for i := 0; i < n%10; i++ {
		s++
	}
	z01.PrintRune(s)
}

func main() {
	points := &point{}

	setPoint(points)
	for _, v := range "x = " {
		z01.PrintRune(v)
	}
	RsmNumber(points.x)
	for _, v := range ", y = " {
		z01.PrintRune(v)
	}
	RsmNumber(points.y)
	z01.PrintRune('\n')
}
