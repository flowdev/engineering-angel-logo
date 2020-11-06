package main

import (
	"fmt"
	"math"
)

// SVG logo for EngineeringAngel, EngNgl, E-NGL, }|->o{, ->o:)

func main() {
	// global
	oSize := 120
	blue := "#00569a"
	//red := "#b00b0b"
	//green := "#0aa00a"
	//orange := "#db6200"
	red := blue
	green := blue
	orange := blue
	margin := oSize / 10
	if margin < 2 {
		margin = 2
	}

	// wing1 (1)
	wing1X := oSize - margin*3/2

	// O (1)
	oR := oSize / 4
	oY := oR

	// wing1 (2)
	wingY := oY + oR + margin/2

	// ^
	arrWidth := margin * 2
	arrHeadR := arrWidth / 2
	arrHeadLen := oSize * 3 / 4
	arrHeadX := wing1X + arrHeadLen*1/6
	arrHeadY := oY + oR + margin

	// O (2)
	oX := arrHeadX

	// |
	arrShaftY := arrHeadY + arrWidth*2
	arrShaftX := arrHeadX - arrHeadR
	arrShaftLen := arrHeadLen

	// square
	squareY := arrShaftY + arrShaftLen + margin
	squareWidth := oSize
	squareHeight := arrWidth

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
		h%d a%d,%d 0 0,0 %d,%d
		v%d a%d,%d 0 0,0 %d,0 v%d
		h%d
		v%d a%d,%d 0 0,0 %d,0 v%d
		H%d
		v%d a%d,%d 0 0,0 %d,0
		Z"
		transform="rotate(45 %d %d)" fill="%s" />`,
		wing1X, wingY,
		-arrWidth*5/2-margin*2, arrHeadR, arrHeadR, -arrWidth/2, arrWidth/2,
		arrWidth, arrHeadR, arrHeadR, arrWidth, -arrWidth/2,
		margin,
		arrWidth*3/2, arrHeadR, arrHeadR, arrWidth, -arrWidth*3/2,
		wing1X-arrWidth,
		arrWidth*2, arrHeadR, arrHeadR, arrWidth,
		wing1X, wingY, orange,
	)
	fmt.Println()

	// O
	fmt.Printf(`  <circle cx="%d" cy="%d" r="%d" fill="%s" />`, oX, oY, oR, green)
	fmt.Println()

	// ^
	sin45 := 0.5 * math.Sqrt(2.0) // sin(45 degree) == (1/2) * sqrt(2) == cos(45 degree)
	arrHeadShift := int(math.Round(float64(arrHeadR) * sin45))
	arrHeadLowShift := pythagorasSmall(arrWidth)
	arrHeadBigLineWH := pythagorasSmall(arrHeadLen - arrWidth + arrHeadShift)
	arrHeadSmlLineWH := pythagorasSmall(arrHeadLen - arrWidth - arrHeadR + arrHeadShift)
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
		arrShaftX, arrShaftY, arrWidth, arrShaftLen, arrHeadR, arrHeadR, blue)
	fmt.Println()

	// square
	fmt.Printf(`  <rect x="%d" y="%d" width="%d" height="%d" fill="%s" />`,
		arrHeadX-squareWidth/2, squareY, squareWidth, squareHeight, red)
	fmt.Println()

	// wing2
	fmt.Printf(`  <path d="M%d,%d
		h%d a%d,%d 0 0,1 %d,%d
		v%d a%d,%d 0 0,1 %d,0 v%d
		h%d
		v%d a%d,%d 0 0,1 %d,0 v%d
		H%d
		v%d a%d,%d 0 0,1 %d,0
		Z"
		transform="rotate(-45 %d %d)" fill="%s" />`,
		wing2X, wingY,
		arrWidth*5/2+margin*2, arrHeadR, arrHeadR, arrWidth/2, arrWidth/2,
		arrWidth, arrHeadR, arrHeadR, -arrWidth, -arrWidth/2,
		-margin,
		arrWidth*3/2, arrHeadR, arrHeadR, -arrWidth, -arrWidth*3/2,
		wing2X+arrWidth,
		arrWidth*2, arrHeadR, arrHeadR, -arrWidth,
		wing2X, wingY, orange,
	)
	fmt.Println()

	/*
		fmt.Printf(`  <line x1="%d" y1="%d" x2="%d" y2="%d" stroke="red" stroke-width="1" />`,
			0, allHeight, allWidth, allHeight)
		fmt.Printf(`  <line x1="%d" y1="%d" x2="%d" y2="%d" stroke="red" stroke-width="1" />`,
			allWidth, 0, allWidth, allHeight)
	*/

	fmt.Println(`</svg>`)
}

func pythagorasSmall(c int) int {
	return int(math.Round(math.Sqrt(float64(c * c / 2))))
}
