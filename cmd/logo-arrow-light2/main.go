package main

import (
	"fmt"
	"math"
)

func main() {
	// global
	oSize := 200
	//red := "#b00b0b"
	blue := "#00569a"
	//green := "#0aa00a"
	//orange := "#db6200"
	red := blue
	green := blue
	orange := blue
	margin := oSize / 10

	// wing1 (1)
	wingM := oSize / 5
	wing1X := oSize - wingM/2

	// O (1)
	oR := oSize / 4
	oY := oR

	// wing1 (2)
	wingY := oY + oR + margin/2

	// ^
	arrWidth := margin * 2
	arrHeadLen := oSize * 3 / 4
	arrHeadX := wing1X + arrHeadLen*1/6
	arrHeadY := oY + oR + margin

	// O (2)
	oX := arrHeadX

	// |
	arrShaftY := arrHeadY + arrWidth*2
	arrShaftX := arrHeadX - arrWidth/2
	arrShaftLen := arrHeadLen

	// square
	squareY := arrShaftY + arrShaftLen + margin
	squareWidth := oSize
	squareHeight := wingM

	// wing2
	wing2X := oX + (oX - wing1X)

	// all
	allWidth := oX * 2
	allHeight := squareY + squareHeight

	// header
	fmt.Println(`<?xml version="1.0" standalone="no" ?>`)
	fmt.Printf(`<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="%dpx" height="%dpx">`,
		allWidth, allHeight)
	fmt.Println()

	// wing1
	fmt.Printf(`  <path d="M%d,%d
		Q%d,%d %d,%d
		T%d,%d
		T%d,%d
		Z"
		fill="%s" />`,
		wing1X-wingM*2, wingY+wingM*2,
		margin/3, wingY+wingM*2, margin/3, wingM*2,
		wingM, margin,
		wing1X, wingY,
		orange,
	)
	/*
		fmt.Printf(`  <path d="M%d,%d
			Q%d,%d %d,%d
			T%d,%d
			T%d,%d
			Z"
			fill="%s" />`,
			wing1X, wingY,
			wingM+margin, wingM*2, wingM, margin*2,
			margin/2, wingM*2,
			wing1X-wingM*2, wingY+wingM*2,
			orange,
		)
	*/
	fmt.Println()

	// O
	fmt.Printf(`  <circle cx="%d" cy="%d" r="%d" fill="%s" />`, oX, oY, oR, green)
	fmt.Println()

	// ^
	sin45 := 0.5 * math.Sqrt(2.0) // sin(45 degree) == (1/2) * sqrt(2) == cos(45 degree)
	arrHeadR := arrWidth / 2
	arrHeadShift := int(math.Round(float64(arrHeadR) * sin45))
	arrHeadLowShift := pythagorasSmall(arrWidth)
	arrHeadBigLineWH := pythagorasSmall(arrHeadLen - arrWidth + arrHeadShift)
	arrHeadSmlLineWH := pythagorasSmall(arrHeadLen - arrWidth - arrWidth/2 + arrHeadShift)
	fmt.Printf(`  <path d="M%d,%d
		a%d,%d 0 0,1 %d,%d
		l %d,%d
		a%d,%d 0 0,1 %d,%d
		l %d,%d
		l %d,%d
		a%d,%d 0 0,1 %d,%d
		Z"
		fill="%s" />`,
		arrHeadX-arrHeadShift, arrHeadY+arrHeadR-arrHeadShift,
		arrHeadR, arrHeadR, arrHeadShift*2, 0,
		arrHeadBigLineWH, arrHeadBigLineWH,
		arrHeadR, arrHeadR, -arrHeadLowShift, arrHeadLowShift,
		-arrHeadSmlLineWH, -arrHeadSmlLineWH,
		-arrHeadSmlLineWH, arrHeadSmlLineWH,
		arrHeadR, arrHeadR, -arrHeadLowShift, -arrHeadLowShift,
		blue,
	)
	fmt.Println()

	// |
	fmt.Printf(`  <rect x="%d" y="%d" width="%d" height="%d" rx="%d" ry="%d" fill="%s" />`,
		arrShaftX, arrShaftY, arrWidth, arrShaftLen, arrWidth/2, arrWidth/2, blue)
	fmt.Println()

	// square
	fmt.Printf(`  <rect x="%d" y="%d" width="%d" height="%d" fill="%s" />`,
		arrHeadX-squareWidth/2, squareY, squareWidth, squareHeight, red)
	fmt.Println()

	// wing2
	fmt.Printf(`  <path d="M%d,%d
		Q%d,%d %d,%d
		T%d,%d
		Z"
		fill="%s" />`,
		wing2X, wingY,
		allWidth-margin/5, wingY, allWidth-margin/5, wingY+wingM*7/2,
		allWidth-wing1X+wingM*5/2, wingY+wingM*5/2,
		orange,
	)
	fmt.Println()

	fmt.Printf(`  <line x1="%d" y1="%d" x2="%d" y2="%d" stroke="red" stroke-width="1" />`,
		0, allHeight, allWidth, allHeight)
	fmt.Printf(`  <line x1="%d" y1="%d" x2="%d" y2="%d" stroke="red" stroke-width="1" />`,
		allWidth, 0, allWidth, allHeight)

	fmt.Println(`</svg>`)
}

func pythagorasSmall(c int) int {
	return int(math.Round(math.Sqrt(float64(c * c / 2))))
}
