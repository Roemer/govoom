# govoom
Go Library to interact with devices from divoom (eg. Pixoo-64) via the HTTP API.

The official documentation for the API is here: https://docin.divoom-gz.com/web/#/5/23
Alternative Link: http://doc.divoom-gz.com/web/#/12?page_id=143

## Installation

```go
go get github.com/roemer/govoom
```

## Usage

```go
devices, err := govoom.FindDevices()
if err != nil {
    // Handle error
}
client := devices[0].GetClient()
// Interact with the client now, eg:
client.SetBrightness(100)
```

### Animation and Images

The usual approach to create a custom screen is to create a background image (or animation) and then display text above.

For background images, there are two variants:
1. Load an image from the disk and use that one
2. Create an image in-memory with the `RgbImage` helper type and its methods.

Loading an image from a file is pretty straight forward like:
```go
// Load and decode the image
f, err := os.Open(pathToImage)
if err != nil {
    // Handle error
}
defer f.Close()
loadedImage, _, err := image.Decode(f)
if err != nil {
    // Handle error
}
// Convert the image to a byte array
imageData := govoom.ImageToRGB24Bytes(loadedImage)
// Reset the ID as otherwise the image is not updated
client.ResetSendingAnimationPicId()
// Send the image as a single frame to the display
client.SendAnimation(1, 1, 0, 64, 1000, imageData)
```

Manually creating an image and using that as background looks like this:
```go
// Create the image in the appropriate size
img := govoom.NewRgbImage(64, 64)
// Use the methods on the image to paint the image or add text directly:
// Drawing pixels
img.DrawPixel(5, 5, govoom.ColorWhite)
// Drawing lines
img.DrawLine(10, 0, 20, 5, govoom.ColorRed)
// Drawing rectangles
img.DrawRectangleFilled(25, 0, 30, 10, govoom.ColorGreen, govoom.ColorAqua)
// Writing text in a PICO8-like font
img.DrawText("Left", 0, 15, govoom.Fonts.Pico8, govoom.ColorWhite, govoom.TextAlignmentLeft)
img.DrawText("Centered", 31, 21, govoom.Fonts.Pico8, govoom.ColorWhite, govoom.TextAlignmentMiddle)
img.DrawText("Right", 63, 27, govoom.Fonts.Pico8, govoom.ColorWhite, govoom.TextAlignmentRight)
// Add an image from the disk
mySmallImage := loadImage(pathToImage) // Same as for the background image
img.DrawImage(30, 30, mySmallImage)
// Reset the id and send the image as background
client.ResetSendingAnimationPicId()
client.SendAnimation(1, 1, 0, 64, 1000, img.Data)
```

### Texts

If you want to write text with the api above a background image, you can use the `SendDisplayList` method on the client.
That way, you can keep a single background image (or animation) and just regularly refresh the texts.

```go
// Send background image first as described before
// Clear the texts
client.ClearAllTextArea()
// Send new texts
client.SendDisplayList(
    govoom.DisplayListElement{
        Id:            1,
        TextType:      govoom.TextTypeText,
        X:             59,
        Y:             0,
        Font:          34,
        Width:         64,
        Height:        5,
        TextAlignment: govoom.TextAlignmentRight,
        Color:         "#FFFFFF",
        Text:          "Hello",
    },
    govoom.DisplayListElement{
        Id:            2,
        TextType:      govoom.TextTypeText,
        X:             63,
        Y:             8,
        Font:          34,
        Width:         64,
        Height:        5,
        TextAlignment: govoom.TextAlignmentRight,
        Color:         "#FFFFFF",
        Text:          "World",
    },
)
```
