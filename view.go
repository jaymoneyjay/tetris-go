package main

import (
	"github.com/nsf/termbox-go"
)

const boardStart = 0
const cellWidth = 3
const boardEnd = boardWidth*cellWidth + boardStart

const backgroundColor = termbox.ColorLightGray

var pieceColors = []termbox.Attribute{
	termbox.ColorRed,
	termbox.ColorBlue,
	termbox.ColorGreen,
	termbox.ColorYellow,
	termbox.ColorLightRed,
	termbox.ColorLightCyan,
	termbox.ColorMagenta,
}

func render(g *Game) {
	termbox.Clear(backgroundColor, backgroundColor)
	for y := 0; y < boardHeight; y++ {
		for x := 0; x < boardWidth; x++ {
			pieceIndex := g.board[y][x]
			pieceColor := pieceColors[pieceIndex]
			for c := 0; c < cellWidth; c++ {
				termbox.SetCell(boardStart+cellWidth*x+c, boardStart+y, ' ', pieceColor, pieceColor)
			}
		}
	}
	termbox.Flush()
}
