package main

/*import (
	"fmt"
)
*/

type Stück struct {
	Farbe string
	Typ   string
}

type Schachfeld struct {
	Feld [][]Stück
}

func NeueInstanz() *Schachfeld {
	feld := make([][]Stück, 8)
	for i := range feld {
		feld[i] = make([]Stück, 8)
		for j := range feld[i] {
			feld[i][j] = Stück{Farbe: "Leer", Typ: "Leer"}
		}
	}
	return &Schachfeld{Feld: feld}
}
