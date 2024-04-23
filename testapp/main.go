package main

import "github.com/roemer/govoom"

func main() {
	img := govoom.NewRgbImage(256, 64)

	str := ""
	for charCode := 33; charCode < 127; charCode++ {
		str += string(rune(charCode))
	}
	img.DrawText(str+"."+",", 0, 56, govoom.Fonts.Slumbers, govoom.ColorBlue, govoom.TextAlignmentLeft)

	img.DrawText("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.,:;_-+%*(){}[]<=>?!^|¦/\\'\"`", 0, 0, govoom.Fonts.Slumbers, govoom.ColorAqua, govoom.TextAlignmentLeft)
	img.DrawText("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.,:;_-+%*(){}[]<=>?!^|¦/\\'\"`", 0, 6, govoom.Fonts.Pico8, govoom.ColorAqua, govoom.TextAlignmentLeft)
	img.DrawText("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.,:;_-+%*(){}[]<=>?!^|¦/\\'\"`", 0, 12, govoom.Fonts.Pico8Narrow, govoom.ColorAqua, govoom.TextAlignmentLeft)

	img.SaveToPng("image.png", 1)
}
