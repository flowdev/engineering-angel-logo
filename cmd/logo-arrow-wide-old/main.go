package main

import (
	"fmt"
	"math"
)

func main() {
	// global
	oSize := 200
	color := "black"
	stroke := oSize / 10
	margin := stroke/2 + oSize/20
	featherN := 5

	// wing1
	wing1ClipX := margin - stroke/2
	wingClipY := margin + oSize*3/4
	wingShoulder := oSize / 6
	wingClipWidth := oSize + stroke/2 + wingShoulder
	wingClipHeight := oSize + stroke/2
	wing1CircX := wing1ClipX + wingClipWidth - wingShoulder
	wingCircY := wingClipY + wingClipHeight
	wingCircR := oSize

	// O (1)
	oR := oSize / 4
	oY := margin + oR

	// ^
	arrWidth := stroke * 2
	arrHeadLen := oSize * 3 / 4
	arrHeadX := wing1CircX + arrHeadLen*5/6
	arrHeadY := oY + oR + stroke + stroke/2

	// O (2)
	oX := arrHeadX

	// |
	arrShaftY := arrHeadY + arrWidth*2 + stroke
	arrShaftX := arrHeadX - arrWidth/2
	arrShaftLen := arrHeadLen // * 4 / 3

	// square
	squareWidth := oSize
	squareHeight := oSize / 4

	// wing2
	wing2ClipX := arrHeadX + (arrHeadX - wing1CircX) - wingShoulder
	wing2CircX := wing2ClipX + wingShoulder

	// all
	widthAll := wing2ClipX + wingClipWidth + margin - stroke/2
	heightAll := wingCircY + oSize/2 + margin

	// header
	fmt.Println(`<?xml version="1.0" ?>`)
	fmt.Printf(`<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="%dpx" height="%dpx" stroke="%s" stroke-width="%d" fill="transparent">`,
		widthAll, heightAll, color, stroke)
	fmt.Println()

	// wing 1
	fmt.Println(`  <clipPath id="wing1Clip">`)
	fmt.Printf(`    <rect x="%d" y="%d" width="%d" height="%d" />`,
		wing1ClipX, wingClipY, wingClipWidth, wingClipHeight)
	fmt.Println()
	fmt.Println(`  </clipPath>`)
	fmt.Printf(`  <circle clip-path="url(#wing1Clip)" cx="%d" cy="%d" r="%d" />`,
		wing1CircX, wingCircY, wingCircR)
	fmt.Println()
	featherR := (wingCircR - stroke/2) / featherN / 2
	feather0X := wing1ClipX + stroke/2
	feather0Y := wingCircY + oSize/2 - featherR
	fmt.Printf(`  <line x1="%d" y1="%d" x2="%d" y2="%d" />`,
		feather0X, wingCircY, feather0X, feather0Y)
	fmt.Println()
	for i := 0; i < featherN; i++ {
		featherX := feather0X + featherR + i*featherR*2
		featherY := feather0Y - i*featherR*2
		fmt.Printf(`  <clipPath id="wing1feather%dClip">`, i+1)
		fmt.Println()
		fmt.Printf(`    <rect x="%d" y="%d" width="%d" height="%d" />`,
			featherX-featherR-stroke/2,
			featherY,
			featherR*2+stroke,
			featherR+stroke/2,
		)
		fmt.Println()
		fmt.Println(`  </clipPath>`)
		fmt.Printf(`  <circle clip-path="url(#wing1feather%dClip)" cx="%d" cy="%d" r="%d" />`,
			i+1, featherX, featherY, featherR)
		fmt.Println()
		fmt.Printf(`  <line x1="%d" y1="%d" x2="%d" y2="%d" />`,
			featherX+featherR,
			featherY-featherR-(featherN-i)*(featherR),
			featherX+featherR,
			featherY,
		)
		fmt.Println()
	}

	// O
	fmt.Printf(`  <circle cx="%d" cy="%d" r="%d" />`, oX, oY, oR)
	fmt.Println()

	// ^
	sin45 := 0.5 * math.Sqrt(2.0) // sin(45 degree) == (1/2) * sqrt(2) == cos(45 degree)
	arrHeadR := arrWidth / 2
	arrHeadShift := int(math.Round(float64(arrHeadR) * sin45))
	arrHeadLowShift := pythagorasSmall(arrWidth)
	arrHeadBigLineWH := pythagorasSmall(arrHeadLen - arrWidth + arrHeadShift)
	arrHeadSmlLineWH := pythagorasSmall(arrHeadLen - arrWidth - arrWidth/2 + arrHeadShift)
	fmt.Printf(`<path d="M%d,%d
		a%d,%d 0 0,1 %d,%d
		l %d,%d
		a%d,%d 0 0,1 %d,%d
		l %d,%d
		l %d,%d
		a%d,%d 0 0,1 %d,%d
		Z" />`,
		arrHeadX-arrHeadShift, arrHeadY+arrHeadR-arrHeadShift,
		arrHeadR, arrHeadR, arrHeadShift*2, 0,
		arrHeadBigLineWH, arrHeadBigLineWH,
		arrHeadR, arrHeadR, -arrHeadLowShift, arrHeadLowShift,
		-arrHeadSmlLineWH, -arrHeadSmlLineWH,
		-arrHeadSmlLineWH, arrHeadSmlLineWH,
		arrHeadR, arrHeadR, -arrHeadLowShift, -arrHeadLowShift,
	)
	fmt.Println()

	// |
	fmt.Printf(`  <rect x="%d" y="%d" width="%d" height="%d" rx="%d" ry="%d" />`,
		arrShaftX, arrShaftY, arrWidth, arrShaftLen, arrWidth/2, arrWidth/2)
	fmt.Println()

	// square
	fmt.Printf(`  <rect x="%d" y="%d" width="%d" height="%d" />`,
		arrHeadX-squareWidth/2, feather0Y+featherR-squareHeight, squareWidth, squareHeight)
	fmt.Println()

	// wing2
	fmt.Println(`  <clipPath id="wing2Clip">`)
	fmt.Printf(`    <rect x="%d" y="%d" width="%d" height="%d" />`,
		wing2ClipX, wingClipY, wingClipWidth, wingClipHeight)
	fmt.Println()
	fmt.Println(`  </clipPath>`)
	fmt.Printf(`  <circle clip-path="url(#wing2Clip)" cx="%d" cy="%d" r="%d" />`,
		wing2CircX, wingCircY, wingCircR)
	fmt.Println()
	feather0X = wing2ClipX + wingClipWidth - stroke/2
	fmt.Printf(`  <line x1="%d" y1="%d" x2="%d" y2="%d" />`,
		feather0X, wingCircY, feather0X, feather0Y)
	fmt.Println()
	for i := 0; i < featherN; i++ {
		featherX := feather0X - featherR - i*featherR*2
		featherY := feather0Y - i*featherR*2
		fmt.Printf(`  <clipPath id="wing2feather%dClip">`, i+1)
		fmt.Println()
		fmt.Printf(`    <rect x="%d" y="%d" width="%d" height="%d" />`,
			featherX-featherR-stroke/2,
			featherY,
			featherR*2+stroke,
			featherR+stroke/2,
		)
		fmt.Println()
		fmt.Println(`  </clipPath>`)
		fmt.Printf(`  <circle clip-path="url(#wing2feather%dClip)" cx="%d" cy="%d" r="%d" />`,
			i+1, featherX, featherY, featherR)
		fmt.Println()
		fmt.Printf(`  <line x1="%d" y1="%d" x2="%d" y2="%d" />`,
			featherX-featherR,
			featherY-featherR-(featherN-i)*(featherR),
			featherX-featherR,
			featherY,
		)
		fmt.Println()
	}

	/*
		fmt.Printf(`  <line x1="%d" y1="%d" x2="%d" y2="%d" stroke="blue" stroke-width="1" />`,
			0, heightAll, widthAll, heightAll)
		fmt.Printf(`  <line x1="%d" y1="%d" x2="%d" y2="%d" stroke="blue" stroke-width="1" />`,
			widthAll, 0, widthAll, heightAll)
	*/
	fmt.Println(`</svg>`)
}

func pythagorasSmall(c int) int {
	return int(math.Round(math.Sqrt(float64(c * c / 2))))
}
