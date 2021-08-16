package decorator

type IDraw interface {
	Draw() string
}

type Square struct{}

func (s Square) Draw() string {
	return "this is a square"
}

type ColorSquare struct {
	square IDraw
	color  string
}

func NewColorSquare(square IDraw, color string) ColorSquare {
	return ColorSquare{square: square, color: color}
}

func (square ColorSquare) Draw() string {
	return square.square.Draw() + ", color is " + square.color
}
