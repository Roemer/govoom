package govoom

type fonts struct {
	Pico8       PixelFont
	Pico8Narrow PixelFont
	Slumbers    PixelFont
}

var Fonts fonts

func init() {
	Fonts = fonts{
		Pico8:       FontPico8,
		Pico8Narrow: FontPico8Narrow,
		Slumbers:    FontSlumbers,
	}
}

// This type represents a font.
type PixelFont struct {
	// The name of ththe font.
	FontName string
	// A short description of the font.
	Description string
	// This flag is true if the font has a fixed width, false otherwise.
	FixedWidth bool
	// The number of pixels added between each character.
	FontSpacing int
	// The height of the font.
	LineHeight int
	// A map which contains the pixel information for each rune.
	Glyphs map[rune]PixelGlyph
}

// This type represents a glyph inside a font.
type PixelGlyph struct {
	// An offset of pixels that are empty from the top.
	Offset int
	// The array with rows/columns of pixels.
	Pixels [][]byte
}

func (p PixelGlyph) GetWidth() int {
	if len(p.Pixels) == 0 {
		return 0
	}
	return len(p.Pixels[0])
}

// Checks if the glyph touches the other glyph on the left side.
// Touch means either a direct or a diagonal connection.
func (rightGlyph PixelGlyph) Touches(leftGlyph PixelGlyph) bool {
	// Loop thru all rows of the left glyph
	for rowIndex, row := range leftGlyph.Pixels {
		// Fix the index by adding the offset
		realRowIndex := rowIndex + leftGlyph.Offset
		// Get the rightmost bit of the left glyph
		leftBit := row[len(row)-1]
		// Continue checking only if it is set (1)
		if leftBit == 1 {
			// Check against -1 to +1 pixel from the right glyph
			for j := -1; j <= 1; j++ {
				secondIndex := realRowIndex + j + rightGlyph.Offset
				if secondIndex >= 0 && secondIndex < len(rightGlyph.Pixels) && rightGlyph.Pixels[secondIndex][0] == 1 {
					return true
				}
			}
		}
	}
	return false
}
