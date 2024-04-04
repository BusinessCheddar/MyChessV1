package main

import (
	"fmt"

	"example.com/goTest/chess"
)

func main() {
	fmt.Println("Hello World")
	s := chess.Square{XCoord: 1, YCoord: 2, Piece: "king"}

	chess.ShowSquare(s)
}
