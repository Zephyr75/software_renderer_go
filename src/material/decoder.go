package material

import (
	"image"
	"log"
	"os"
)

func GetImageFromFilePath(filePath string) Material {
	f, err := os.Open(filePath)
	if err != nil {
		return WhiteMaterial()
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	return TextureMaterial(image)
}
