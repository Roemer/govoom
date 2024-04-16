package govoom

import "image"

func ImageToRGB24Bytes(img image.Image) []byte {
	w, h := img.Bounds().Dx(), img.Bounds().Dy()
	imgData := make([]byte, w*h*3)
	var i int
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			r8, g8, b8 := r>>8, g>>8, b>>8
			imgData[i] = byte(r8)
			imgData[i+1] = byte(g8)
			imgData[i+2] = byte(b8)
			i += 3
		}
	}
	return imgData
}
