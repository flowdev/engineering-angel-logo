package main

import (
	"fmt"
)

// SVG logo for EngineeringAngel, EngNgl, |->Xo, E|->oA, Eo<-|A

func main() {
	// global
	oSize := 120
	blue := "#0091d3"
	red := "#b00b0b"
	green := "#0aa00a"
	//orange := "#db6200"
	//red := blue
	//green := blue
	orange := blue
	margin := oSize / 10
	if margin < 2 {
		margin = 2
	}

	// wing1 (1)
	wingM := oSize / 5
	wing1X := oSize - 3*margin

	// O (1)
	oR := oSize / 4
	oY := oR

	// wing1 (2)
	wingY := oY + oR + margin/2

	// ^
	arrWidth := margin * 2
	arrHeadR := arrWidth / 2
	arrHeadLen := oSize * 3 / 4
	arrHeadX := wing1X + arrHeadLen*1/2 - margin
	arrHeadY := oY + oR + margin/2

	// O (2)
	oX := arrHeadX

	// |
	arrShaftY := arrHeadY + arrWidth*2 + margin
	arrShaftX := arrHeadX
	arrShaftLen := arrHeadLen - arrWidth

	// square
	squareY := arrShaftY + arrShaftLen + margin*3/2
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
	fmt.Printf(`  <path d="M%d,%d h%d
	    M%d,%d v%d
	    M%d,%d v%d
	    M%d,%d v%d"
		transform="rotate(45 %d %d)" stroke="%s" stroke-width="%d" stroke-linecap="round" />`,
		wing1X, wingY, -arrWidth*2-margin*2,
		wing1X, wingY, arrWidth*4/2,
		wing1X-arrWidth-margin, wingY, arrWidth*3/2,
		wing1X-arrWidth*2-margin*2, wingY, arrWidth*2/2,
		wing1X, wingY, orange, arrWidth,
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
	fmt.Printf(`  <path d="M%d,%d v%d"
		stroke="%s" stroke-width="%d" stroke-linecap="round" fill="none" />`,
		arrShaftX, arrShaftY, arrShaftLen,
		blue, arrWidth,
	)
	fmt.Println()

	// square
	fmt.Printf(`  <rect x="%d" y="%d" width="%d" height="%d" fill="%s" />`,
		arrHeadX-squareWidth/2, squareY, squareWidth, squareHeight, red)
	fmt.Println()

	// wing2
	fmt.Printf(`  <path d="M%d,%d v%d
		m0,0 h%d
	    m0,0 L%d,%d
		M%d,%d h%d"
		transform="rotate(45 %d %d)" stroke="%s" stroke-width="%d" stroke-linecap="round" />`,
		wing2X, wingY, -arrWidth*2-margin*2,
		arrWidth,
		wing2X+arrWidth*2, wingY,
		wing2X, wingY-arrWidth, arrWidth*3/2,
		wing2X, wingY, orange, arrWidth,
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
