package main

import "fmt"

type Position struct {
	x int
	y int
}

func main() {
	pos := Position{2,3};
	fmt.Printf("Position x: %d, y: %d\n", pos.x, pos.y)
}
