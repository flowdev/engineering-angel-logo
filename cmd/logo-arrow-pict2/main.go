package main

import (
	"fmt"
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

	// O (1)
	oR := oSize / 4
	oY := oR

	// wing1 (1)
	wingY := oY + oR + margin/2

	// ^ (1)
	arrWidth := margin * 2
	arrHeadR := arrWidth / 2
	arrHeadLen := oSize / 2
	arrHeadY := oY + oR + margin

	// | (2)
	arrShaftY := arrHeadY + arrWidth*2
	arrShaftLen := oSize * 3 / 4

	// square
	squareY := arrShaftY + arrShaftLen + margin
	squareWidth := oSize
	squareHeight := arrWidth

	// all
	allHeight := squareY + squareHeight
	allWidth := allHeight

	// ^ (2)
	arrHeadX := allWidth / 2

	// O (2)
	oX := arrHeadX

	// | (2)
	arrShaftX := arrHeadX - arrHeadR

	// wing1 (2)
	wing1X := oSize - margin*3/2

	// wing2
	//wing2X := oX + (oX - wing1X)

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
	fmt.Printf(`  <path d="M%d,%d
		h%d
		a%d,%d 0 0,1 %d,%d
		v%d"
		transform="rotate(-45 %d %d)" stroke="%s" stroke-width="%d" stroke-linecap="round" fill="none" />`,
		arrHeadX-arrHeadLen-arrHeadR/2, arrHeadY+arrHeadR,
		arrHeadLen-arrHeadR/2,
		arrHeadR, arrHeadR, arrHeadR, arrHeadR,
		arrHeadLen-arrHeadR/2,
		arrHeadX, arrHeadY+arrHeadR, blue, arrWidth,
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
		h%d
		a%d,%d 0 0,1 %d,%d
		V%d
		a%d,%d 0 0,0 %d,%d
		h%d
		h%d
		a%d,%d 0 0,0 %d,%d
		a%d,%d 0 0,1 %d,%d
		h%d"
		stroke="%s" stroke-width="%d" stroke-linecap="round" fill="none" />`,
		allWidth-arrHeadR, squareY-arrWidth/2,
		-margin,
		arrWidth, arrWidth, -arrWidth, -arrWidth,
		wingY+arrWidth*3/2+margin/2,
		arrWidth*3/2, arrWidth*3/2, -arrWidth*3/2, -arrWidth*3/2,
		-arrHeadR,
		arrHeadR,
		arrWidth*3/2, arrWidth*3/2, arrWidth*3/2, -arrWidth*3/2,
		arrWidth, arrWidth, arrWidth, -arrWidth,
		margin,
		orange, arrWidth,
	)
	fmt.Println()

	fmt.Printf(`  <line x1="%d" y1="%d" x2="%d" y2="%d" stroke="red" stroke-width="1" />`,
		0, allHeight, allWidth, allHeight)
	fmt.Printf(`  <line x1="%d" y1="%d" x2="%d" y2="%d" stroke="red" stroke-width="1" />`,
		allWidth, 0, allWidth, allHeight)

	fmt.Println(`</svg>`)
}
