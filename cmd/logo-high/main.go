package main

import (
	"fmt"
)

func main() {
	// global
	oSize := 300
	margin := oSize / 10
	color := "black"
	stroke := oSize / 10
	featherN := 5

	// wing1
	wing1ClipX := margin - stroke/2
	wingClipY := margin + oSize
	wingClipWidth := oSize + stroke/2 + margin
	wingClipHeight := oSize + stroke/2
	wing1CircX := wing1ClipX + wingClipWidth - margin
	wingCircY := wingClipY + wingClipHeight
	wingCircR := oSize

	// O (1)
	oR := oSize / 2
	oY := margin + oR

	// ( )
	parenR := oSize * 3 / 4
	parenClipX := wing1CircX
	parenClipY := oY + oR + stroke/2
	parenClipWidth := parenR*2 + stroke
	parenClipHeight := oSize + stroke
	parenX := parenClipX + parenR + stroke/2
	parenY := parenClipY + oR + stroke/2

	// O (2)
	oX := parenX

	// n
	nCircR := oSize / 3
	nClipX := parenClipX + oSize*4/10 - stroke
	nClipY := parenClipY
	nClipWidth := oSize + stroke
	nClipHeight := nCircR + stroke/2 + margin
	nCircX := parenX
	nCircY := nClipY + nClipHeight
	nLine1X := nCircX - nCircR
	nLine1Y1 := nCircY - nCircR - stroke/2
	nLine2X := nCircX + nCircR
	nLine2Y1 := nCircY
	nLineY2 := parenClipY + parenClipHeight - margin

	// wing2
	wing2ClipX := parenClipX + parenClipWidth - margin
	wing2CircX := wing2ClipX + margin

	// all
	widthAll := wing2ClipX + wingClipWidth + margin - stroke/2
	heightAll := wingCircY + oSize/2 + margin //oSize*3 + margin*2

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
	featherR := (wingCircR - stroke) / featherN / 2
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

	// ( )
	fmt.Println(`  <clipPath id="parenClip">`)
	fmt.Printf(`    <rect x="%d" y="%d" width="%d" height="%d" />`,
		parenClipX, parenClipY, parenClipWidth, parenClipHeight)
	fmt.Println()
	fmt.Println(`  </clipPath>`)
	fmt.Printf(`  <circle clip-path="url(#parenClip)" cx="%d" cy="%d" r="%d" />`,
		parenX, parenY, parenR)
	fmt.Println()

	// n
	fmt.Println(`  <clipPath id="nClip">`)
	fmt.Printf(`    <rect x="%d" y="%d" width="%d" height="%d" />`,
		nClipX, nClipY, nClipWidth, nClipHeight)
	fmt.Println()
	fmt.Println(`  </clipPath>`)
	fmt.Printf(`  <line x1="%d" y1="%d" x2="%d" y2="%d" />`,
		nLine1X, nLine1Y1, nLine1X, nLineY2)
	fmt.Println()
	fmt.Printf(`  <circle clip-path="url(#nClip)" cx="%d" cy="%d" r="%d" />`,
		nCircX, nCircY, nCircR)
	fmt.Println()
	fmt.Printf(`  <line x1="%d" y1="%d" x2="%d" y2="%d" />`,
		nLine2X, nLine2Y1, nLine2X, nLineY2)
	fmt.Println()

	// wing 2
	fmt.Println(`  <clipPath id="wing2Clip">`)
	fmt.Printf(`    <rect x="%d" y="%d" width="%d" height="%d" />`,
		wing2ClipX, wingClipY, wingClipWidth, wingClipHeight)
	fmt.Println()
	fmt.Println(`  </clipPath>`)
	fmt.Printf(`  <circle clip-path="url(#wing2Clip)" cx="%d" cy="%d" r="%d" />`,
		wing2CircX, wingCircY, wingCircR)
	fmt.Println()
	featherR = (wingCircR - stroke) / featherN / 2
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
			//feather0Y-(i+1)*featherR*2-featherR-(featherN-i)*stroke,
			featherX-featherR,
			featherY,
		)
		fmt.Println()
	}

	fmt.Printf(`  <line x1="%d" y1="%d" x2="%d" y2="%d" stroke="blue" stroke-width="1" />`,
		0, heightAll, widthAll, heightAll)
	fmt.Printf(`  <line x1="%d" y1="%d" x2="%d" y2="%d" stroke="blue" stroke-width="1" />`,
		widthAll, 0, widthAll, heightAll)
	fmt.Println(`</svg>`)
}
