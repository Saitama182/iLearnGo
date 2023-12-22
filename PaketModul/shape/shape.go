package shape

type Square struct {
	sideLenght int
}

func NewSquare(length int) Square {
	return Square{
		sideLenght: length,
	}
}
