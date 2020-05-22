package main

import (
	"fmt"
)

func main() {
	// global
	oSize := 100
	margin := oSize / 10
	color := "black"
	stroke := oSize / 10
	featherN := 5

	// wing1
	wing1ClipX := margin - stroke/2
	wing1ClipY := margin - stroke/2
	wing1ClipWidth := oSize + stroke/2
	wing1ClipHeight := oSize + stroke/2
	wing1CircX := wing1ClipX + wing1ClipWidth
	wing1CircY := wing1ClipY + wing1ClipHeight
	wing1CircR := oSize

	// O
	oR := oSize / 2
	oX := wing1ClipX + wing1ClipWidth + oR
	oY := margin + oR

	// ( )
	parenR := oSize * 3 / 4
	parenClipX := oX + oR + margin
	parenClipY := oY - oR - stroke/2
	parenClipWidth := parenR*2 + stroke
	parenClipHeight := oSize + stroke
	parenX := parenClipX + parenR + stroke/2
	parenY := oY

	// n
	nClipX := parenClipX + oSize*4/10 - stroke
	nClipY := parenClipY
	nClipWidth := oSize + stroke
	nClipHeight := oSize/2 + stroke
	nCircR := oSize / 3
	nCircX := parenX
	nCircY := nClipY + nClipHeight
	nLine1X := nCircX - nCircR
	nLine1Y1 := nClipY + nClipHeight - nCircR - stroke/2
	nLineY1 := nClipY + nClipHeight
	nLineY2 := nLineY1 + oSize/2
	nLine2X := nCircX + nCircR

	// wing2
	wing2X := parenClipX + parenClipWidth
	wing2Y := margin
	wing2ClipWidth := oSize + stroke
	wing2ClipHeight := oSize*2 + stroke/2
	wing2CircR := oSize
	wing2ElliRX := wing2ClipWidth - stroke/2
	wing2ElliRY := wing2ClipHeight - stroke/2

	// all
	widthAll := wing2X + wing2ClipWidth + margin
	heightAll := oSize*2 + margin*2

	// header
	fmt.Println(`<?xml version="1.0" ?>`)
	fmt.Printf(`<svg version="1.1" xmlns="http://www.w3.org/2000/svg" width="%dpx" height="%dpx" stroke="%s" stroke-width="%d" fill="transparent">`,
		widthAll, heightAll, color, stroke)
	fmt.Println()

	// wing 1
	fmt.Println(`  <clipPath id="wing1Clip">`)
	fmt.Printf(`    <rect x="%d" y="%d" width="%d" height="%d" stroke-width="1"/>`,
		wing1ClipX, wing1ClipY, wing1ClipWidth, wing1ClipHeight)
	fmt.Println()
	fmt.Println(`  </clipPath>`)
	fmt.Printf(`  <circle clip-path="url(#wing1Clip)" cx="%d" cy="%d" r="%d" />`,
		wing1CircX, wing1CircY, wing1CircR)
	fmt.Println()
	featherR := (wing1CircR - stroke) / featherN / 2
	feather0X := wing1ClipX + stroke/2
	feather0Y := wing1CircY + oSize - featherR
	fmt.Printf(`  <line x1="%d" y1="%d" x2="%d" y2="%d" />`,
		feather0X, wing1CircY, feather0X, feather0Y)
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
			feather0Y-(i+1)*featherR*2-featherR-(featherN-i)*stroke,
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

	// base line
	fmt.Printf(`  <line x1="%d" y1="%d" x2="%d" y2="%d" stroke="red" stroke-width="1" />`,
		margin+oSize,
		parenClipY+parenClipHeight+1,
		parenClipX+parenClipWidth,
		parenClipY+parenClipHeight+1,
	)
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
		nLine2X, nLineY1, nLine2X, nLineY2)
	fmt.Println()

	// wing 2
	fmt.Println(`  <clipPath id="wing2Clip">`)
	fmt.Printf(`    <rect x="%d" y="%d" width="%d" height="%d" />`,
		wing2X, wing2Y, wing2ClipWidth, wing2ClipHeight)
	fmt.Println()
	fmt.Println(`  </clipPath>`)
	fmt.Printf(`  <ellipse clip-path="url(#wing2Clip)" cx="%d" cy="%d" rx="%d" ry="%d" />`,
		wing2X, wing2Y, wing2ElliRX, wing2ElliRY)
	fmt.Printf(`  <circle clip-path="url(#wing2Clip)" cx="%d" cy="%d" r="%d" />`,
		wing2X, wing2Y, wing2CircR)
	fmt.Println()

	// lowest line
	fmt.Printf(`  <line x1="%d" y1="%d" x2="%d" y2="%d" stroke="blue" stroke-width="1" />`,
		0, heightAll, widthAll, heightAll)
	fmt.Println()

	fmt.Println(`</svg>`)
}
