package govoom

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

type RgbImage struct {
	Data   []byte
	Width  int
	Height int
}

func NewRgbImage(width, height int) *RgbImage {
	return &RgbImage{
		Data:   make([]byte, width*height*3),
		Width:  width,
		Height: height,
	}
}

func (i *RgbImage) SaveToPng(path string, scale int) error {
	// Create the image
	img := image.NewRGBA(image.Rect(0, 0, i.Width*scale, i.Height*scale))

	// Set the pixels
	for x := 0; x < i.Width; x++ {
		for y := 0; y < i.Height; y++ {
			index := (x * i.Height) + y
			dataIndex := index * 3
			for sx := 0; sx < scale; sx++ {
				for sy := 0; sy < scale; sy++ {
					img.Set(x*scale+sx, y*scale+sy, color.RGBA{i.Data[dataIndex+0], i.Data[dataIndex+1], i.Data[dataIndex+2], 0xff})
				}
			}
		}
	}

	// Encode as PNG.
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	return png.Encode(f, img)
}

func (image *RgbImage) Clear() {
	image.DrawRectangleFilled(0, 0, image.Width-1, image.Height-1, ColorBlack, ColorBlack)
}

func (image *RgbImage) Fill(color Color) {
	image.DrawRectangleFilled(0, 0, image.Width-1, image.Height-1, color, color)
}

func (image *RgbImage) DrawPixel(x, y int, color Color) {
	// The pixel is outside of the bounds, skip it
	if x < 0 || y < 0 || x >= image.Width || y >= image.Height {
		return
	}

	// Calculate the index
	index := (x * image.Height) + y

	// Draw the pixel at the index
	image.DrawPixelAtIndex(index, color)
}

func (image *RgbImage) DrawPixelAtIndex(index int, color Color) {
	// The index outside of the bounds, skip it
	if index < 0 || index >= image.Width*image.Height {
		return
	}

	// Calculate the index in the data array
	dataIndex := index * 3

	// Set the pixel
	image.Data[dataIndex+0] = color.R
	image.Data[dataIndex+1] = color.G
	image.Data[dataIndex+2] = color.B
}

func (image *RgbImage) DrawLine(startX, startY int, endX, endY int, color Color) {
	// Use Bresenham’s Line Algorithm
	dx := int(math.Abs(float64(endX) - float64(startX)))
	sx := -1
	if startX < endX {
		sx = 1
	}
	dy := -int(math.Abs(float64(endY) - float64(startY)))
	sy := -1
	if startY < endY {
		sy = 1
	}
	e := dx + dy
	x := startX
	y := startY
	for {
		image.DrawPixel(x, y, color)
		if x == endX && y == endY {
			break
		}
		e2 := 2 * e
		if e2 >= dy {
			if x == endX {
				break
			}
			e = e + dy
			x = x + sx
		}
		if e2 <= dx {
			if y == endY {
				break
			}
			e = e + dx
			y = y + sy
		}
	}
}

func (image *RgbImage) DrawRectangle(left, top int, right, bottom int, borderColor Color) {
	image.DrawLine(left, top, right, top, borderColor)
	image.DrawLine(right, top, right, bottom, borderColor)
	image.DrawLine(right, bottom, left, bottom, borderColor)
	image.DrawLine(left, bottom, left, top, borderColor)
}

func (image *RgbImage) DrawRectangleFilled(left, top int, right, bottom int, borderColor Color, fillColor Color) {
	image.DrawRectangle(left, top, right, bottom, borderColor)
	for x := left + 1; x < right; x++ {
		for y := top + 1; y < bottom; y++ {
			image.DrawPixel(x, y, fillColor)
		}
	}
}

func (image *RgbImage) DrawImage(x, y int, img image.Image) {
	w, h := img.Bounds().Dx(), img.Bounds().Dy()
	for py := 0; py < h; py++ {
		for px := 0; px < w; px++ {
			r, g, b, a := img.At(px, py).RGBA()
			if a == 0 {
				// Skip pixels with full transparency
				continue
			}
			if a < 0xff {
				// Todo: Merge colors like:
				// Target.R = ((1 - Source.A) * BGColor.R) + (Source.A * Source.R)
				// Target.G = ((1 - Source.A) * BGColor.G) + (Source.A * Source.G)
				// Target.B = ((1 - Source.A) * BGColor.B) + (Source.A * Source.B)
			}
			r8, g8, b8 := r>>8, g>>8, b>>8
			image.DrawPixel(px+x, py+y, NewColor(byte(r8), byte(g8), byte(b8)))
		}
	}
}

func (image *RgbImage) DrawGlyph(glyph PixelGlyph, x, y int, color Color) int {
	for localY, dy := range glyph.Pixels {
		for localX, dx := range dy {
			if dx == 1 {
				image.DrawPixel(x+localX, y+localY+glyph.Offset, color)
			}
		}
	}
	return glyph.GetWidth()
}

func (image *RgbImage) DrawCharacter(character rune, x, y int, font PixelFont, color Color) int {
	glyph, exists := font.Glyphs[character]
	if !exists {
		return 0
	}
	return image.DrawGlyph(glyph, x, y, color)
}

func (image *RgbImage) DrawText(text string, x, y int, font PixelFont, color Color, alignment TextAlignment) int {
	// Create a list with glyphs and their local position
	glyphs := []PixelGlyph{}
	glyphsPos := []int{}
	currX := 0
	for i, rune := range []rune(text) {
		glyph, ok := font.Glyphs[rune]
		if !ok {
			continue
		}
		// Calculate the new x
		if i > 0 {
			// Add a space if needed
			if font.FixedWidth || glyph.Touches(glyphs[i-1]) {
				currX += font.FontSpacing
			}
		}
		// Add the glyph and the position
		glyphs = append(glyphs, glyph)
		glyphsPos = append(glyphsPos, currX)
		// Add the width of the glyph itself
		currX += glyph.GetWidth()
	}
	// Adjust the x position for different alignments
	totalWidth := currX
	if alignment == TextAlignmentRight {
		x = x - totalWidth + 1
	} else if alignment == TextAlignmentMiddle {
		x = x - totalWidth/2
	}
	// Draw the glyphs
	for i, glyph := range glyphs {
		image.DrawGlyph(glyph, glyphsPos[i]+x, y, color)
	}
	return totalWidth
}
