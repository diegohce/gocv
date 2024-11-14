package main

import (
	"fmt"

	"gocv.io/x/gocv"
)

func main() {

	img := gocv.IMRead("../../images/face-detect.jpeg", gocv.IMReadColor)
	if img.Empty() {
		fmt.Println("invalid read of Mat")
	}
	defer img.Close()

	dest := gocv.NewMat()
	defer dest.Close()

	err := gocv.Safe(func() {

		gocv.CvtColor(img, &dest, gocv.ColorBGRAToGray)

	})
	if err != nil {
		fmt.Println(err, "converting image to gray scale")
	}
}
