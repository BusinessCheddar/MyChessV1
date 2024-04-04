package chess

type Square struct {
	xCoord int
	yCoord int
	piece  string
}

func showSquare(s Square) {
	if s.xCoord < 0 {
		println("-")
	} else {
		return
	}

}
