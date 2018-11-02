package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"math"
	"os"
	//Initilize so image.Decode can understand these formats
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/BurntSushi/graphics-go/graphics"
)

var (
	violet = color.RGBA{148, 0, 211, 255}
	indigo = color.RGBA{75, 0, 130, 255}
	blue   = color.RGBA{0, 0, 255, 255}
	green  = color.RGBA{0, 255, 0, 255}
	yellow = color.RGBA{255, 255, 0, 255}
	orange = color.RGBA{255, 127, 0, 255}
	red    = color.RGBA{255, 0, 0, 255}
)

var colors = []color.RGBA{violet, indigo, blue, green, yellow, orange, red}

func main() {
	generateImage("test.png", 200, 400)
	getImgDimensions("test.png")
	rotateImage("papadu.jpg", "papadu2.jpg")
	printImageConfig("papadu.jpg")
	makeRainbow("rainbow.png", 200, 400)
}

func getImgDimensions(filename string) {
	imageFile, err := os.Open(filename)
	checkErr(err)
	defer imageFile.Close()

	decodedImg, err := png.Decode(imageFile) //decode the whole image
	dimensions := decodedImg.Bounds()
	fmt.Println(dimensions.Dx(), dimensions.Dy())
}

func makeRainbow(filename string, width, height int) {
	imageFile, err := os.Create(filename) //byteio
	checkErr(err)
	defer imageFile.Close()

	newImg := image.NewRGBA(image.Rect(0, 0, width, height)) //image.Image type

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			row := int(float64(j) / (float64((height + 7) / 7))) //height-7 eliminate the edge case when row is calculated to 7
			newImg.SetRGBA(i, j, colors[row])
		}
	}

	png.Encode(imageFile, newImg)
	checkErr(err)
	fmt.Println("Rainbow shooting from "+filename+" width and height: ", width, height)
}

func printImageConfig(filename string) {
	imageFile, err := os.Open(filename)
	checkErr(err)
	defer imageFile.Close()

	imageConfig, err := jpeg.DecodeConfig(imageFile)
	checkErr(err)
	fmt.Println("Decoded config for "+filename+" : ", imageConfig)

}
func rotateImage(srcFilename, dstFilename string) {
	srcImgFile, err := os.Open(srcFilename)
	checkErr(err)
	defer srcImgFile.Close()

	imageConf, err := jpeg.Decode(srcImgFile)
	dimensions := imageConf.Bounds()

	newImg := image.NewRGBA(image.Rect(0, 0, dimensions.Dy(), dimensions.Dx()))
	graphics.Rotate(newImg, imageConf, &graphics.RotateOptions{math.Pi / 2.0})

	dstImageFile, _ := os.Create(dstFilename)
	defer dstImageFile.Close()
	jpeg.Encode(dstImageFile, newImg, &jpeg.Options{jpeg.DefaultQuality})

	fmt.Println(dstFilename + " is rotated version of " + srcFilename)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func generateImage(filename string, width, height int) {

	myImage := image.NewRGBA(image.Rect(0, 0, width, height))
	for i := 0; i < (width * height * 4); i++ {
		myImage.Pix[i] = 122
	}
	// outputFile is a File type which satisfies Writer interface
	outputFile, err := os.Create(filename)
	checkErr(err)
	defer outputFile.Close()
	// Encode takes a writer interface and an image interface
	// We pass it the File and the RGBA
	png.Encode(outputFile, myImage)
	checkErr(err)
	fmt.Println(filename + " created")
}
