package main

import (
	"time"
)

// Constants
const boardHeight = 10
const boardWidth = 20
const gameSpeed = 1000 * time.Millisecond

// Collection of all possible tetris pieces modeled with offset coordinates
pieces = []Pieces{
	//half cross
	newPiece({0, 1, 2, 1}, {0, 0, 0, 1}),
	//straight
	newPiece({0, 1, 2, 3}, {0, 0, 0, 0}),
	//left L
	newPiece({0, 1, 1, 1}, {0, 0, 1, 2}),
	//right L
	newPiece({0, 0, 0, 1}, {0, 1, 2, 1}),
	//square
	newPiece({0, 0, 1, 1}, {0, 1, 0, 1}),
	//left knee
	newPiece({0, 0, 1, 1}, {0, 1, 1, 2}),
	//right knee
	newPiece({0, 0, 1, 1}, {0, 1, 0, -1})
}

// Struct to model the tetris piece
type Piece struct {
	deltaX []int
	deltaY []int
}

// Create a new piece, not exported
func newPiece(deltaX []int, deltaY []int) *Piece {
	p := new(Piece)
	p.deltaX = deltaX
	p.deltaY = deltaY
	return p
}

//Constast type for direction
type Direction int

const (
	down Direction = iota
	left
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

// deletePiece()

// placePiece()

// pieceFits()

// movePiece(direction=[down, left, right])

//rotate()

//checkRow()

//fillBoard()

//spawnPiece()



