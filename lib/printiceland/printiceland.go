package printiceland

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg"
	"log"
	"math/rand"

	"github.com/ondrejsika/go-iceland"
	"github.com/qeesung/image2ascii/convert"
)

func PrintRadomIcelandPhoto() {
	images := []*[]byte{
		&iceland.ICELAND_SUNSET_2022,
		&iceland.ICELAND_LAVA_FIELD_2022,
		&iceland.ICELAND_RIVER_AT_POOL_2022,
		&iceland.ICELAND_ICE_2018,
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
