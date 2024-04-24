package main

import "github.com/roemer/govoom"

func main() {
	img := govoom.NewRgbImage(256, 64)

	str := ""
	for charCode := 33; charCode < 127; charCode++ {
		str += string(rune(charCode))
	}
	img.DrawText(str+"."+",", 0, 56, govoom.Fonts.Slumbers, govoom.Colors.Blue, govoom.TextAlignmentLeft)

	img.DrawText("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.,:;_-+%*(){}[]<=>?!^|¦/\\'\"`", 0, 0, govoom.Fonts.Slumbers, govoom.Colors.Aqua, govoom.TextAlignmentLeft)
	img.DrawText("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.,:;_-+%*(){}[]<=>?!^|¦/\\'\"`", 0, 6, govoom.Fonts.Pico8, govoom.Colors.Aqua, govoom.TextAlignmentLeft)
	img.DrawText("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.,:;_-+%*(){}[]<=>?!^|¦/\\'\"`", 0, 12, govoom.Fonts.Pico8Narrow, govoom.Colors.Aqua, govoom.TextAlignmentLeft)

	img.DrawRectangleFilled(5, 20, 20, 5, govoom.Colors.Orange, govoom.Colors.Salmon)

	for radius := 0; radius <= 10; radius++ {
		img.DrawCircle(40+(radius*radius)+2*radius, 30, radius, govoom.Colors.Green)
	}
	for radius := 0; radius <= 10; radius++ {
		img.DrawCircleFilled(40+(radius*radius)+2*radius, 40, radius, govoom.Colors.Blue, govoom.Colors.Green)
	}

	img.SaveToPng("image.png", 10)
}
