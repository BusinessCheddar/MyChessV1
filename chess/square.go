package chess

type Square struct {
	XCoord int
	YCoord int
	Piece  string
}

func ShowSquare(s Square) {
	if s.XCoord > 0 {
		println("-")
	} else {
		return
	}

}
