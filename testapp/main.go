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

	img.DrawRectangleFilled(5, 20, 1, 1, govoom.Colors.Orange, govoom.Colors.Salmon)
	img.DrawRectangleFilled(8, 20, 2, 2, govoom.Colors.Orange, govoom.Colors.Salmon)
	img.DrawRectangleFilled(12, 20, 3, 3, govoom.Colors.Orange, govoom.Colors.Salmon)
	img.DrawRectangleFilled(18, 20, 20, 5, govoom.Colors.Orange, govoom.Colors.Salmon)

	img.DrawRectangle(5, 30, 1, 1, govoom.Colors.Red)
	img.DrawRectangle(8, 30, 2, 2, govoom.Colors.Red)
	img.DrawRectangle(12, 30, 3, 3, govoom.Colors.Red)
	img.DrawRectangle(18, 30, 20, 5, govoom.Colors.Red)

	for radius := 0; radius <= 10; radius++ {
		img.DrawCircle(40+(radius*radius)+2*radius, 30, radius, govoom.Colors.Green)
	}
	for radius := 0; radius <= 10; radius++ {
		img.DrawCircleFilled(40+(radius*radius)+2*radius, 40, radius, govoom.Colors.Blue, govoom.Colors.Green)
	}

	img.SaveToPng("image.png", 10)
}
