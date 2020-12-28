package main

import (
	"math"

	"github.com/nsf/termbox-go"
)

// Layout
const boardStart = 0
const cellWidth = 3
const boardEnd = boardWidth*cellWidth + boardStart
const backgroundColor = termbox.ColorLightGray

// Text
const title = "TETRIS"

// Colors
var pieceColors = []termbox.Attribute{
	termbox.ColorBlack,
	termbox.ColorBlue,
	termbox.ColorGreen,
	termbox.ColorYellow,
	termbox.ColorLightRed,
	termbox.ColorLightCyan,
	termbox.ColorMagenta,
	termbox.ColorRed,
}

// render renders the board of the game
func render(g *Game) {
	termbox.Clear(backgroundColor, backgroundColor)
	tbprint(0, 0, termbox.ColorYellow, backgroundColor, title)
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

// Function tbprint draws a string.
func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}
