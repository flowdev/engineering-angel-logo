package main

import (
	"fmt"
)

// SVG logo for EngineeringAngel, EngNgl, |->{o

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
	wingM := oSize / 5
	wing1X := oSize - wingM*4/2 + margin/2

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
			L%d,%d
			Q%d,%d %d,%d
			Z"
			stroke="%s" stroke-width="%d" fill="none" />`,
		wing1X-margin, wingY+margin/2,
		wing1X-wingM*9/4-margin/2, wingY-wingM*9/4,
		-wingM/2, wingY-wingM/2, wing1X-wingM*3/2-margin/2, wingY+wingM*3/2,
		orange, margin,
	)
	fmt.Println()

	// O
	fmt.Printf(`  <circle cx="%d" cy="%d" r="%d" fill="%s" />`, oX, oY, oR, green)
	fmt.Println()

	// ^
	fmt.Printf(`  <path d="M%d,%d
		h%d
		a%d,%d 0 0,1 %d,%d
		v%d"
		transform="rotate(-45 %d %d)" stroke="%s" stroke-width="%d" stroke-linecap="round" fill="none" />`,
		arrHeadX-arrHeadLen+arrHeadR+margin, arrHeadY+arrHeadR,
		arrHeadLen-arrHeadR*3/2-margin,
		arrHeadR/2, arrHeadR/2, arrHeadR/2, arrHeadR/2,
		arrHeadLen-arrHeadR*3/2-margin,
		arrHeadX, arrHeadY+arrHeadR, blue, arrWidth,
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
			L%d,%d
			Q%d,%d %d,%d
			Z"
			stroke="%s" stroke-width="%d" fill="none" />`,
		wing2X+margin, wingY+margin/2,
		wing2X+wingM*9/4+margin/2, wingY-wingM*9/4,
		allWidth+wingM/2, wingY-wingM/2, wing2X+wingM*3/2+margin/2, wingY+wingM*3/2,
		orange, margin,
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
