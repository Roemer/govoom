package govoom

import "fmt"

// A type to represent a color.
type Color struct {
	R byte
	G byte
	B byte
}

// Constructor method for a color.
func NewColor(r, g, b byte) Color {
	return Color{R: r, G: g, B: b}
}

func (c Color) ToHex() string {
	return fmt.Sprintf("#%02X%02X%02X", c.R, c.G, c.B)
}

// A handfull of colors ready to be used.
var Colors colors

type colors struct {
	// Pink
	MediumVioletRed Color
	DeepPink        Color
	PaleVioletRed   Color
	HotPink         Color
	LightPink       Color
	Pink            Color
	// Red
	DarkRed     Color
	Red         Color
	FireBrick   Color
	Crimson     Color
	IndianRed   Color
	LightCoral  Color
	Salmon      Color
	DarkSalmon  Color
	LightSalmon Color
	// Orange
	OrangeRed  Color
	Tomato     Color
	DarkOrange Color
	Coral      Color
	Orange     Color
	// Yellow
	DarkKhaki            Color
	Gold                 Color
	Khaki                Color
	PeachPuff            Color
	Yellow               Color
	PaleGoldenrod        Color
	Moccasin             Color
	PapayaWhip           Color
	LightGoldenrodYellow Color
	LemonChiffon         Color
	LightYellow          Color
	// Brown
	Maroon         Color
	Brown          Color
	SaddleBrown    Color
	Sienna         Color
	Chocolate      Color
	DarkGoldenrod  Color
	Peru           Color
	RosyBrown      Color
	Goldenrod      Color
	SandyBrown     Color
	Tan            Color
	BurlyWood      Color
	Wheat          Color
	NavajoWhite    Color
	Bisque         Color
	BlanchedAlmond Color
	Cornsilk       Color
	// Purple
	Indigo          Color
	Purple          Color
	DarkMagenta     Color
	DarkViolet      Color
	DarkSlateBlue   Color
	BlueViolet      Color
	DarkOrchid      Color
	Fuchsia         Color
	Magenta         Color
	SlateBlue       Color
	MediumSlateBlue Color
	MediumOrchid    Color
	MediumPurple    Color
	Orchid          Color
	Violet          Color
	Plum            Color
	Thistle         Color
	Lavender        Color
	RebeccaPurple   Color
	// Blue
	MidnightBlue   Color
	Navy           Color
	DarkBlue       Color
	MediumBlue     Color
	Blue           Color
	RoyalBlue      Color
	SteelBlue      Color
	DodgerBlue     Color
	DeepSkyBlue    Color
	CornflowerBlue Color
	SkyBlue        Color
	LightSkyBlue   Color
	LightSteelBlue Color
	LightBlue      Color
	PowderBlue     Color
	// Cyan
	Teal            Color
	DarkCyan        Color
	LightSeaGreen   Color
	CadetBlue       Color
	DarkTurquoise   Color
	MediumTurquoise Color
	Turquoise       Color
	Aqua            Color
	Cyan            Color
	Aquamarine      Color
	PaleTurquoise   Color
	LightCyan       Color
	// Green
	DarkGreen         Color
	Green             Color
	DarkOliveGreen    Color
	ForestGreen       Color
	SeaGreen          Color
	Olive             Color
	OliveDrab         Color
	MediumSeaGreen    Color
	LimeGreen         Color
	Lime              Color
	SpringGreen       Color
	MediumSpringGreen Color
	DarkSeaGreen      Color
	MediumAquamarine  Color
	YellowGreen       Color
	LawnGreen         Color
	Chartreuse        Color
	LightGreen        Color
	GreenYellow       Color
	PaleGreen         Color
	// White
	MistyRose     Color
	AntiqueWhite  Color
	Linen         Color
	Beige         Color
	WhiteSmoke    Color
	LavenderBlush Color
	OldLace       Color
	AliceBlue     Color
	SeaShell      Color
	GhostWhite    Color
	HoneyDew      Color
	FloralWhite   Color
	Azure         Color
	MintCream     Color
	Snow          Color
	Ivory         Color
	White         Color
	// Gray and Black
	Black          Color
	DarkSlateGray  Color
	DimGray        Color
	SlateGray      Color
	Gray           Color
	LightSlateGray Color
	DarkGray       Color
	Silver         Color
	LightGray      Color
	Gainsboro      Color
}

func init() {
	Colors = colors{
		// Pink
		MediumVioletRed: NewColor(199, 21, 133),
		DeepPink:        NewColor(255, 20, 147),
		PaleVioletRed:   NewColor(219, 112, 147),
		HotPink:         NewColor(255, 105, 180),
		LightPink:       NewColor(255, 182, 193),
		Pink:            NewColor(255, 192, 203),
		// Red
		DarkRed:     NewColor(139, 0, 0),
		Red:         NewColor(255, 0, 0),
		FireBrick:   NewColor(178, 34, 34),
		Crimson:     NewColor(220, 20, 60),
		IndianRed:   NewColor(205, 92, 92),
		LightCoral:  NewColor(240, 128, 128),
		Salmon:      NewColor(250, 128, 114),
		DarkSalmon:  NewColor(233, 150, 122),
		LightSalmon: NewColor(255, 160, 122),
		// Orange
		OrangeRed:  NewColor(255, 69, 0),
		Tomato:     NewColor(255, 99, 71),
		DarkOrange: NewColor(255, 140, 0),
		Coral:      NewColor(255, 127, 80),
		Orange:     NewColor(255, 165, 0),
		// Yellow
		DarkKhaki:            NewColor(189, 183, 107),
		Gold:                 NewColor(255, 215, 0),
		Khaki:                NewColor(240, 230, 140),
		PeachPuff:            NewColor(255, 218, 185),
		Yellow:               NewColor(255, 255, 0),
		PaleGoldenrod:        NewColor(238, 232, 170),
		Moccasin:             NewColor(255, 228, 181),
		PapayaWhip:           NewColor(255, 239, 213),
		LightGoldenrodYellow: NewColor(250, 250, 210),
		LemonChiffon:         NewColor(255, 250, 205),
		LightYellow:          NewColor(255, 255, 224),
		// Brown
		Maroon:         NewColor(128, 0, 0),
		Brown:          NewColor(165, 42, 42),
		SaddleBrown:    NewColor(139, 69, 19),
		Sienna:         NewColor(160, 82, 45),
		Chocolate:      NewColor(210, 105, 30),
		DarkGoldenrod:  NewColor(184, 134, 11),
		Peru:           NewColor(205, 133, 63),
		RosyBrown:      NewColor(188, 143, 143),
		Goldenrod:      NewColor(218, 165, 32),
		SandyBrown:     NewColor(244, 164, 96),
		Tan:            NewColor(210, 180, 140),
		BurlyWood:      NewColor(222, 184, 135),
		Wheat:          NewColor(245, 222, 179),
		NavajoWhite:    NewColor(255, 222, 173),
		Bisque:         NewColor(255, 228, 196),
		BlanchedAlmond: NewColor(255, 235, 205),
		Cornsilk:       NewColor(255, 248, 220),
		// Purple
		Indigo:          NewColor(75, 0, 130),
		Purple:          NewColor(128, 0, 128),
		DarkMagenta:     NewColor(139, 0, 139),
		DarkViolet:      NewColor(148, 0, 211),
		DarkSlateBlue:   NewColor(72, 61, 139),
		BlueViolet:      NewColor(138, 43, 226),
		DarkOrchid:      NewColor(153, 50, 204),
		Fuchsia:         NewColor(255, 0, 255),
		Magenta:         NewColor(255, 0, 255),
		SlateBlue:       NewColor(106, 90, 205),
		MediumSlateBlue: NewColor(123, 104, 238),
		MediumOrchid:    NewColor(186, 85, 211),
		MediumPurple:    NewColor(147, 112, 219),
		Orchid:          NewColor(218, 112, 214),
		Violet:          NewColor(238, 130, 238),
		Plum:            NewColor(221, 160, 221),
		Thistle:         NewColor(216, 191, 216),
		Lavender:        NewColor(230, 230, 250),
		RebeccaPurple:   NewColor(102, 51, 153),
		// Blue
		MidnightBlue:   NewColor(25, 25, 112),
		Navy:           NewColor(0, 0, 128),
		DarkBlue:       NewColor(0, 0, 139),
		MediumBlue:     NewColor(0, 0, 205),
		Blue:           NewColor(0, 0, 255),
		RoyalBlue:      NewColor(65, 105, 225),
		SteelBlue:      NewColor(70, 130, 180),
		DodgerBlue:     NewColor(30, 144, 255),
		DeepSkyBlue:    NewColor(0, 191, 255),
		CornflowerBlue: NewColor(100, 149, 237),
		SkyBlue:        NewColor(135, 206, 235),
		LightSkyBlue:   NewColor(135, 206, 250),
		LightSteelBlue: NewColor(176, 196, 222),
		LightBlue:      NewColor(173, 216, 230),
		PowderBlue:     NewColor(176, 224, 230),
		// Cyan
		Teal:            NewColor(0, 128, 128),
		DarkCyan:        NewColor(0, 139, 139),
		LightSeaGreen:   NewColor(32, 178, 170),
		CadetBlue:       NewColor(95, 158, 160),
		DarkTurquoise:   NewColor(0, 206, 209),
		MediumTurquoise: NewColor(72, 209, 204),
		Turquoise:       NewColor(64, 224, 208),
		Aqua:            NewColor(0, 255, 255),
		Cyan:            NewColor(0, 255, 255),
		Aquamarine:      NewColor(127, 255, 212),
		PaleTurquoise:   NewColor(175, 238, 238),
		LightCyan:       NewColor(224, 255, 255),
		// Green
		DarkGreen:         NewColor(0, 100, 0),
		Green:             NewColor(0, 128, 0),
		DarkOliveGreen:    NewColor(85, 107, 47),
		ForestGreen:       NewColor(34, 139, 34),
		SeaGreen:          NewColor(46, 139, 87),
		Olive:             NewColor(128, 128, 0),
		OliveDrab:         NewColor(107, 142, 35),
		MediumSeaGreen:    NewColor(60, 179, 113),
		LimeGreen:         NewColor(50, 205, 50),
		Lime:              NewColor(0, 255, 0),
		SpringGreen:       NewColor(0, 255, 127),
		MediumSpringGreen: NewColor(0, 250, 154),
		DarkSeaGreen:      NewColor(143, 188, 139),
		MediumAquamarine:  NewColor(102, 205, 170),
		YellowGreen:       NewColor(154, 205, 50),
		LawnGreen:         NewColor(124, 252, 0),
		Chartreuse:        NewColor(127, 255, 0),
		LightGreen:        NewColor(144, 238, 144),
		GreenYellow:       NewColor(173, 255, 47),
		PaleGreen:         NewColor(152, 251, 152),
		// White
		MistyRose:     NewColor(255, 228, 225),
		AntiqueWhite:  NewColor(250, 235, 215),
		Linen:         NewColor(250, 240, 230),
		Beige:         NewColor(245, 245, 220),
		WhiteSmoke:    NewColor(245, 245, 245),
		LavenderBlush: NewColor(255, 240, 245),
		OldLace:       NewColor(253, 245, 230),
		AliceBlue:     NewColor(240, 248, 255),
		SeaShell:      NewColor(255, 245, 238),
		GhostWhite:    NewColor(248, 248, 255),
		HoneyDew:      NewColor(240, 255, 240),
		FloralWhite:   NewColor(255, 250, 240),
		Azure:         NewColor(240, 255, 255),
		MintCream:     NewColor(245, 255, 250),
		Snow:          NewColor(255, 250, 250),
		Ivory:         NewColor(255, 255, 240),
		White:         NewColor(255, 255, 255),
		// Gray and Black
		Black:          NewColor(0, 0, 0),
		DarkSlateGray:  NewColor(47, 79, 79),
		DimGray:        NewColor(105, 105, 105),
		SlateGray:      NewColor(112, 128, 144),
		Gray:           NewColor(128, 128, 128),
		LightSlateGray: NewColor(119, 136, 153),
		DarkGray:       NewColor(169, 169, 169),
		Silver:         NewColor(192, 192, 192),
		LightGray:      NewColor(211, 211, 211),
		Gainsboro:      NewColor(220, 220, 220),
	}
}
