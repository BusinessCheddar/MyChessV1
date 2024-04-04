package main

import (
	"fmt"

	"example.com/goTest/chess"
)

func main() {
	fmt.Println("Hello World")
	p := chess.Piece{Type: "King", Color: "W"}
	s := chess.Square{XCoord: 1, YCoord: 2, Piece: p}

	chess.ShowSquare(s)
}
