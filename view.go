package main

import (
	"math"

	"github.com/nsf/termbox-go"
)

// Constants
const boardStart = 0
const cellWidth = 3
const boardEnd = boardWidth*cellWidth + boardStart
const backgroundColor = termbox.ColorLightGray

// Colors
var pieceColors = []termbox.Attribute{
	termbox.ColorRed,
	termbox.ColorBlue,
	termbox.ColorGreen,
	termbox.ColorYellow,
	termbox.ColorLightRed,
	termbox.ColorLightCyan,
	termbox.ColorMagenta,
}

// render renders the games board
func render(g *Game) {
	termbox.Clear(backgroundColor, backgroundColor)
	for y := 0; y < boardHeight; y++ {
		for x := 0; x < boardWidth; x++ {
			pieceIndex := int(math.Abs(float64(g.board[y][x])))
			pieceColor := pieceColors[pieceIndex]
			for c := 0; c < cellWidth; c++ {
				termbox.SetCell(boardStart+cellWidth*x+c, boardStart+y, ' ', pieceColor, pieceColor)
			}
		}
	}
	termbox.Flush()
}
