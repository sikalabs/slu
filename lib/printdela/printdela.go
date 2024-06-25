package printdela

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg"
	"log"
	"math/rand"

	"github.com/ondrejsika/go-dela"
	"github.com/qeesung/image2ascii/convert"
)

func PrintDela() {
	img, _, err := image.Decode(bytes.NewReader(dela.DELA1_JPG))
	if err != nil {
		log.Fatalln(err)
	}
	converter := convert.NewImageConverter()
	fmt.Print(converter.Image2ASCIIString(img, &convert.DefaultOptions))
}

func PrintRandomDela() {
	images := []*[]byte{
		&dela.DELA1_JPG,
		&dela.DELA2_JPG,
	}
	randomIndex := rand.Intn(len(images))
	randomImage := images[randomIndex]

	img, _, err := image.Decode(bytes.NewReader(*randomImage))
	if err != nil {
		log.Fatalln(err)
	}
	converter := convert.NewImageConverter()
	fmt.Print(converter.Image2ASCIIString(img, &convert.DefaultOptions))
}
