package main

import (
	"fmt"

	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
		"image/png"
	"log"
	"os"
	//	"github.com/BurntSushi/graphics-go/graphics"
	//	"math"
)

var (
	red    = color.RGBA{255, 0, 0, 150}
	orange = color.RGBA{255, 127, 0, 150}
	yellow = color.RGBA{255, 255, 0, 150}
	green  = color.RGBA{0, 255, 0, 150}
	blue   = color.RGBA{0, 0, 255, 150}
	indigo = color.RGBA{75, 0, 130, 150}
	violet = color.RGBA{148, 0, 211, 150}
)

var colors = []color.RGBA{red, yellow, orange, green, blue, indigo, violet}

func main() {
	generateImage("test.jpg", 20, 10)
	getImgDimensions("test.jpg")
	printImageConfig("test.jpg")
	makeRainbow("rainbow.png", 200, 400)
	rainbowOverlay("gopherized.jpg", "rainbowedGopherized.jpg")
	//	rotateImage("gopherized.png", "rotatedGopherized.png") //BONUS: rotating image with an external package
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
	jpeg.Encode(outputFile, myImage, &jpeg.Options{100})
	checkErr(err)
	fmt.Println("Generated ", filename)
}

func getImgDimensions(filename string) {
	imageFile, err := os.Open(filename)
	checkErr(err)
	defer imageFile.Close()

	//decode the []bytes to image.Image
	decodedImg, err := jpeg.Decode(imageFile)
	dimensions := decodedImg.Bounds()
	fmt.Println("width: ", dimensions.Dx(), " Height: ", dimensions.Dy())
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
	fmt.Println("Made rainbow: "+filename+" width and height: ", width, height)
}

func printImageConfig(filename string) {
	imageFile, err := os.Open(filename)
	checkErr(err)
	defer imageFile.Close()

	imageConfig, err := jpeg.DecodeConfig(imageFile)
	checkErr(err)
	fmt.Println("Decoded config for "+filename+" : ", imageConfig)
}

func rainbowOverlay(givenFilename, newFilename string) {
	//get the given image's dimension
	givenFile, err := os.Open(givenFilename)
	checkErr(err)
	defer givenFile.Close()
	givenImg, _, _ := image.Decode(givenFile)
	dimensions := givenImg.Bounds()

	makeRainbow("rainbow.png", dimensions.Dx(), dimensions.Dy())
	srcFile, err := os.Open("rainbow.png")
	checkErr(err)
	defer srcFile.Close()
	srcImg, _, _ := image.Decode(srcFile)

	m := image.NewRGBA(image.Rect(0, 0, dimensions.Dx(), dimensions.Dy()))
	draw.Draw(m, m.Bounds(), givenImg, image.Point{0, 0}, draw.Src)
	draw.Draw(m, m.Bounds(), srcImg, image.Point{0, 0}, draw.Over)

	imageFile, err := os.Create(newFilename)
	jpeg.Encode(imageFile, m, &jpeg.Options{100})
	checkErr(err)
}

/*BONUS STUFF
func rotateImage(srcFilename, dstFilename string) {
	srcImgFile, err := os.Open(srcFilename)
	checkErr(err)
	defer srcImgFile.Close()

	imageConf, err := png.Decode(srcImgFile)
	dimensions := imageConf.Bounds()

	newImg := image.NewRGBA(image.Rect(0, 0, dimensions.Dy(), dimensions.Dx()))
	graphics.Rotate(newImg, imageConf, &graphics.RotateOptions{math.Pi / 2.0})

	dstImageFile, _ := os.Create(dstFilename)
	defer dstImageFile.Close()
	png.Encode(dstImageFile, newImg)

	fmt.Println(dstFilename + " is rotated version of " + srcFilename)
}
*/
