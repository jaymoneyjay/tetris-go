package main

import (
	"math/rand"
	"time"
)

// Constants
const boardHeight = 30
const boardWidth = 10
const gameSpeed = 1000 * time.Millisecond
const numberSquares = 4
const numberPieces = 7

// Struct to model the tetris piece
type Piece struct {
	deltaX []int
	deltaY []int
	color  int
}

// Collection of all possible tetris pieces modeled with offset coordinates
var pieces = []Piece{
	//half cross
	newPiece([]int{0, 1, 2, 1}, []int{0, 0, 0, 1}, 0),
	//straight
	newPiece([]int{0, 1, 2, 3}, []int{0, 0, 0, 0}, 1),
	//left L
	newPiece([]int{0, 1, 1, 1}, []int{0, 0, 1, 2}, 2),
	//right L
	newPiece([]int{0, 0, 0, 1}, []int{0, 1, 2, 1}, 3),
	//square
	newPiece([]int{0, 0, 1, 1}, []int{0, 1, 0, 1}, 4),
	//left knee
	newPiece([]int{0, 0, 1, 1}, []int{0, 1, 1, 2}, 5),
	//right knee
	newPiece([]int{0, 0, 1, 1}, []int{0, 1, 0, -1}, 6),
}

// Create a new piece, not exported
func newPiece(deltaX []int, deltaY []int, color int) Piece {
	p := new(Piece)
	p.deltaX = deltaX
	p.deltaY = deltaY
	p.color = color
	return *p
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
	b := make([][]int, boardHeight)
	for y := 0; y < boardHeight; y++ {
		b[y] = make([]int, boardWidth)
		for x := 0; x < boardWidth; x++ {
			b[y][x] = initValue
		}
	}
	return b
}

// TODO: Methods for game

func NewGame() *Game {
	g := new(Game)
	g.x = 5
	g.y = 5
	g.board = initializeBoard(0)
	g.clock = time.NewTimer(gameSpeed)
	g.spawnRandomPiece()
	return g
}
func (g *Game) play() {
	for i := 0; i < 4; i++ {
		g.movePiece(down)
	}
}

func (g *Game) resetClock() {
	g.clock.Reset(gameSpeed)
}

// pieceFits()
func (g *Game) pieceFits(x int, y int) bool {
	return true
}

// movePiece(direction=[down, left, right])
func (g *Game) movePiece(dir Direction) bool {
	if dir == down {
		if g.pieceFits(g.x, g.y+1) {
			g.deletePiece()
			g.y++
			g.placePiece()
			return true
		}
	} else {
		if g.pieceFits(g.x+int(dir), g.y) {
			g.deletePiece()
			g.x += int(dir)
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
		g.board[y][x] = value
	}
}

// deletePiece()
func (g *Game) deletePiece() {
	g.updateBoard(0)
}

// placePiece()
func (g *Game) placePiece() {
	g.updateBoard(g.piece.color)
}

//spawnPiece()
func (g *Game) spawnRandomPiece() {
	color := rand.Int() % numberPieces
	g.piece = pieces[color]
	g.x = boardWidth / 2
	g.y = 0
	g.placePiece()
}
