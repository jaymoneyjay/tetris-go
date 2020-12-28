package main

import (
	"time"
	"math/rand"
)

// Constants
const boardHeight = 10
const boardWidth = 20
const gameSpeed = 1000 * time.Millisecond
const numberSquares = 4
const numberPieces = 7

// Collection of all possible tetris pieces modeled with offset coordinates
pieces = []Pieces{
	color = 0
	//half cross
	newPiece({0, 1, 2, 1}, {0, 0, 0, 1}, color++),
	//straight
	newPiece({0, 1, 2, 3}, {0, 0, 0, 0}, color++),
	//left L
	newPiece({0, 1, 1, 1}, {0, 0, 1, 2}, color++),
	//right L
	newPiece({0, 0, 0, 1}, {0, 1, 2, 1}, color++),
	//square
	newPiece({0, 0, 1, 1}, {0, 1, 0, 1}, color++),
	//left knee
	newPiece({0, 0, 1, 1}, {0, 1, 1, 2}, color++),
	//right knee
	newPiece({0, 0, 1, 1}, {0, 1, 0, -1}, color++)
}

// Struct to model the tetris piece
type Piece struct {
	deltaX []int
	deltaY []int
	color int
}

// Create a new piece, not exported
func newPiece(deltaX []int, deltaY []int, color) *Piece {
	p := new(Piece)
	p.deltaX = deltaX
	p.deltaY = deltaY
	p.color = color
	return p
}

//Constast type for direction
type Direction int

const (
	left Direction = iota - 1
	down
	right
)


// Struct to model the game
type Game struct {
	x int
	y int
	//board[y][x]
	board [][]int
	piece Piece
	clock *time.Timer
}

// helper functions
func initializeBoard(initValue int) [][]int {
	b = make([][]int, boardHeight)
	for y := 0; y < boardHeight; y++ {
		b[y] make([]int, boardWidth)
		for x := 0; x < boardWidth; x++ {
			b[y][x] = initValue
		}
	}
	return b
}

// TODO: Methods for game

func NewGame() (*Game) {
	g := new(Game)
	g.x = 1
	g.y = 1
	g.board = initializeBoard(0)
	g.clock = time.NewTimer(gameSpeed)

}
func (g *Game) play() {
	if movePiece(down) {
		g.resetClock()
	}
}

func (g *Game) resetClock() {
	g.clock.Reset(gameSpeed)
}



// pieceFits()
func (g *Game) pieceFits() bool {
	return true
}

// movePiece(direction=[down, left, right])
func (g *Game) movePiece(dir Direction) bool {
	if dir == down {
		if g.pieceFits(g.x, g.y + 1) {
			g.deletePiece()
			g.y++
			g.placePiece()
			return true
		}
	} else {
		if g.pieceFits(g.x + dir, g.y) {
			g.deletePiece()
			g.x += dir
			g.placePiece()
			return true
		}
	}
	return false
}

//rotate()

//checkRow()

//updateBoard()
func (g *Game) updateBoard(value int) {
	for i := 0; i < numberSquares; i++ {
		x := g.x + g.piece.deltaX[i]
		y := g.y + g.piece.deltaY[i]
		g.board[x][y] = value
	}
}

// deletePiece()
func (g *Game) deletePiece() {
	updateBoard(0)
}

// placePiece()
func (g *Game) placePiece() {
	updateBoard(piece.color)
}

//spawnPiece()
func (g *Game) spawnRandomPiece() {
	color = rand.Int() % numberPieces
	g.piece = pieces[color]
	g.x = boardWidth / 2
	g.y = 0
	g.placePiece()
}



