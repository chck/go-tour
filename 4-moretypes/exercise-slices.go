package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	img := make([][]uint8, 0, dy)
	for y := range make([]uint8, dy) {
		a := make([]uint8, 0, dx)
		for x := range make([]uint8, dx) {
			a = append(a, uint8((x+y)/2))
		}
		img = append(img, a)
	}
	return img
}

func main() {
	pic.Show(Pic)
}
