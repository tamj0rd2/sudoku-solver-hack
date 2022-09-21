package main

import (
	"fmt"
	"gocv.io/x/gocv"
	_ "gocv.io/x/gocv"
	"image"
)

// guide I'm following - https://aishack.in/tutorials/sudoku-grabber-opencv-plot/
// run this directly from the IDE or using `make run2` after installing the mac prerequisites:
// brew install opencv
// brew install pkgconfig
func main() {
	fmt.Println("hello world")

	sudoku := gocv.IMRead("cmd/cli2/testdata/sudoku.png", gocv.IMReadGrayScale)
	outerBox := gocv.NewMatWithSize(sudoku.Size()[0], sudoku.Size()[1], gocv.MatTypeCV8UC1)

	// blur the image to reduce noise
	gocv.GaussianBlur(sudoku, &sudoku, image.Point{11, 11}, 0, 0, gocv.BorderDefault)
	// threshold the image
	gocv.AdaptiveThreshold(sudoku, &outerBox, 255, gocv.AdaptiveThresholdMean, gocv.ThresholdBinary, 5, 2)
	// flip the bits so that the stuff we want to select (the lines) are white
	gocv.BitwiseNot(outerBox, &outerBox)
	// fill in any gaps that have formed
	gocv.Dilate(outerBox, &outerBox, gocv.GetStructuringElement(gocv.MorphCross, image.Point{3, 3}))

	sudokuWindow := gocv.NewWindow("Sudoku")
	sudokuWindow.IMShow(sudoku)
	sudokuWindow.MoveWindow(-1500, 0)

	outerBoxWindow := gocv.NewWindow("Outer Box")
	outerBoxWindow.IMShow(outerBox)
	outerBoxWindow.MoveWindow(-1000, 0)

	for {
		sudokuWindow.WaitKey(1)
		outerBoxWindow.WaitKey(1)
	}
}
