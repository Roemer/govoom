package govoom

var (
	ColorWhite   = Color{R: 255, G: 255, B: 255}
	ColorBlack   = Color{R: 000, G: 000, B: 000}
	ColorRed     = Color{R: 255, G: 000, B: 000}
	ColorGreen   = Color{R: 000, G: 255, B: 000}
	ColorBlue    = Color{R: 000, G: 000, B: 255}
	ColorYellow  = Color{R: 255, G: 255, B: 000}
	ColorMagenta = Color{R: 255, G: 000, B: 255}
	ColorAqua    = Color{R: 000, G: 255, B: 255}
)

type Color struct {
	R byte
	G byte
	B byte
}

func NewColor(r, g, b byte) Color {
	return Color{R: r, G: g, B: b}
}
