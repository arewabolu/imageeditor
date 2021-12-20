package main

import (
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

//image converter

var (
	height   int64
	width    int64
	inputimg string
)

func main() {
	inputimg = "nohorny.jpg"
	file, err := os.Open(inputimg)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	nwimgfile, err := os.Create("img.png")
	if err != nil {
		log.Fatal(err)
	}
	defer nwimgfile.Close()

	enc := png.Encoder{
		CompressionLevel: png.NoCompression,
	}

	err = enc.Encode(nwimgfile, img)
	if err != nil {
		log.Fatal(err)
	}

}
