package main

import (
	"fmt"
	"strings"

	"github.com/roemer/govoom"
)

var maxBoxesPerRow = 26

func main() {
	numTitles := 5
	numSampleTexts := 2

	for _, font := range []govoom.PixelFont{govoom.FontSlumbers, govoom.Fonts.Pico8, govoom.Fonts.Pico8Narrow, govoom.Fonts.Smallest} {
		titleHeight := govoom.FontPico8.LineHeight + 1
		sampleTextHeight := font.LineHeight + 1
		headerBoxHeight := govoom.FontPico8.LineHeight + 4
		headerBoxWidth := 15
		textBoxHeight := font.LineHeight + 4
		img := govoom.NewRgbImage(maxBoxesPerRow*(headerBoxWidth-1)+1, numTitles*titleHeight+numSampleTexts*sampleTextHeight+10*(headerBoxHeight-1)+10*(textBoxHeight-1)+(numTitles)*2)

		y := 0
		y = draw(img, font, "0123456789", "Numbers:", y)
		y = draw(img, font, "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "Letters (Upper):", y)
		y = draw(img, font, "abcdefghijklmnopqrstuvwxyz", "Letters (Lower):", y)
		y = draw(img, font, "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~", "Special:", y) // 33-47, 58-64, 91-96, 123-126
		str := ""
		for charCode := 128; charCode <= 255; charCode++ {
			str += string(rune(charCode))
		}
		y = draw(img, font, str, "Extended:", y)

		// Draw some sample text
		sampleText := "The quick brown fox jumps over the lazy dog"
		img.DrawText(strings.ToUpper(sampleText), 0, y, font, govoom.Colors.White, govoom.TextAlignmentLeft)
		y += sampleTextHeight
		img.DrawText(sampleText, 0, y, font, govoom.Colors.White, govoom.TextAlignmentLeft)

		// Save the image
		img.SaveToPng(strings.ToLower(font.FontName)+".png", 2)
	}
}

func draw(img *govoom.RgbImage, font govoom.PixelFont, text string, title string, startY int) int {
	newY := startY
	if title != "" {
		img.DrawText(title, 0, newY, govoom.FontPico8, govoom.Colors.White, govoom.TextAlignmentLeft)
		newY += govoom.FontPico8.LineHeight + 1
	}
	if text != "" {
		headerBoxHeight := govoom.FontPico8.LineHeight + 4
		headerBoxWidth := 15
		textBoxHeight := font.LineHeight + 4
		x := 0
		y := newY
		for _, r := range text {
			// Wrap around if x gets too large
			if x+headerBoxHeight > img.Width {
				x = 0
				y += headerBoxHeight + textBoxHeight - 2
			}

			// Draw the rectangles
			img.DrawRectangle(x, y, headerBoxWidth, headerBoxHeight, govoom.Colors.Gray)
			img.DrawRectangle(x, y+headerBoxHeight-1, headerBoxWidth, textBoxHeight, govoom.Colors.Gray)

			// Draw the texts
			img.DrawText(fmt.Sprintf("%03d", r), x+2, y+2, govoom.FontPico8, govoom.Colors.LightGray, govoom.TextAlignmentLeft)
			img.DrawText(string(r), x+headerBoxWidth/2, y+headerBoxHeight+1, font, govoom.Colors.White, govoom.TextAlignmentMiddle)

			// Increase the x position for the next box
			x += headerBoxWidth - 1
		}
		// Add the last row of boxes
		newY = y + headerBoxHeight + textBoxHeight
	}
	return newY
}
