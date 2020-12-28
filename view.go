package main

import (
	"fmt"
	"math"

	"github.com/nsf/termbox-go"
)

// Layout
const boardStart = 0
const cellWidth = 3
const cellHeight = 1
const boardEndX = boardWidth*cellWidth + boardStart
const boardEndY = boardHeight*cellHeight + boardStart
const backgroundColor = termbox.ColorLightGray

// Text
const title = "TETRIS"

// Instructions
var instructions = []string{
	"q: quit",
	"arrows: left, right, down",
	"up: rotate",
}

// Colors
var pieceColors = []termbox.Attribute{
	termbox.ColorBlack,
	termbox.ColorLightCyan,
	termbox.ColorBlue,
	termbox.ColorLightRed,
	termbox.ColorYellow,
	termbox.ColorGreen,
	termbox.ColorRed,
	termbox.ColorMagenta,
}

var yPos = 0

func render(g *Game) {
	termbox.Clear(backgroundColor, backgroundColor)

	// render renders the board of the game
	for y := 0; y < boardHeight; y++ {
		for x := 0; x < boardWidth; x++ {
			pieceIndex := int(math.Abs(float64(g.board[y][x])))
			pieceColor := pieceColors[pieceIndex]
			for cw := 0; cw < cellWidth; cw++ {
				for ch := 0; ch < cellHeight; ch++ {
					termbox.SetCell(boardStart+cellWidth*x+cw, boardStart+cellHeight*y+ch, ' ', pieceColor, pieceColor)
				}
			}
		}
	}

	// render title
	tbPrint(boardEndX, 0, termbox.ColorYellow, backgroundColor, title)

	// render instructions
	for i, instr := range instructions {
		tbPrint(boardEndX, i+1, termbox.ColorBlack, backgroundColor, instr)
	}

	scoreText := fmt.Sprintf("Score: %d\n", Score)
	tbPrint(0, boardEndY, termbox.ColorGreen, backgroundColor, scoreText)

	termbox.Flush()
}

// Function tbprint draws a string.
func tbPrint(x, y int, fg, bg termbox.Attribute, text string) {
	for _, c := range text {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func tbPrintln(x int, fg, bg termbox.Attribute, text string) {
	tbPrint(x, yPos, fg, bg, text)
	yPos++
}
