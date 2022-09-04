package material

import (
	"image"
	"os"
	"log"
)

func GetImageFromFilePath(filePath string) image.Image {
    f, err := os.Open(filePath)
    if err != nil {
        return nil
    }
    defer f.Close()
    image, _, err := image.Decode(f)
	if err != nil {
        log.Fatal(err)
    }
    return image
}