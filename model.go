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

// Piece represents a tetris piece by storing the color and coordinate offset
type Piece struct {
	deltaX []int
	deltaY []int
	color  int
}

// Collection of all possible tetris pieces modeled with offset coordinates
var pieces = []Piece{
	// empty piece
	{},

	//straight
	newPiece([]int{-1, 0, 1, 2}, []int{0, 0, 0, 0}, 1),

	//left L
	newPiece([]int{-1, 0, 0, 0}, []int{1, 1, 0, -1}, 2),

	//half cross
	newPiece([]int{0, 0, 0, 1}, []int{-1, 0, 1, 0}, 3),

	//square
	newPiece([]int{0, 0, 1, 1}, []int{0, 1, 0, 1}, 4),

	//left knee
	newPiece([]int{0, 0, 1, 1}, []int{-1, 0, 0, 1}, 5),

	//right knee
	newPiece([]int{0, 0, 1, 1}, []int{0, 1, 0, -1}, 6),

	//right L
	newPiece([]int{0, 0, 0, 1}, []int{-1, 0, 1, 1}, 7),
}

// newPiece returns a new instantiation of a Piece
func newPiece(deltaX []int, deltaY []int, color int) Piece {
	p := new(Piece)
	p.deltaX = deltaX
	p.deltaY = deltaY
	p.color = color
	return *p
}

// deepCopy returns a deep copy of the specified piece
func deepCopy(piece Piece) Piece {
	deltaXPrime := make([]int, numberSquares)
	deltaYPrime := make([]int, numberSquares)

	for i := 0; i < numberSquares; i++ {
		deltaXPrime[i] = piece.deltaX[i]
		deltaYPrime[i] = piece.deltaY[i]
	}

	colorPrime := piece.color

	return newPiece(deltaXPrime, deltaYPrime, colorPrime)
}

// rotate rotates a piece by swapping the delta values and flip one sign
func (p Piece) rotate() {
	for i := 0; i < numberSquares; i++ {
		newY := -p.deltaX[i]
		newX := p.deltaY[i]
		p.deltaX[i] = newX
		p.deltaY[i] = newY
	}
}

// Direction represents the directions left, down and right
type Direction int

const (
	left Direction = iota - 1
	down
	right
)

// Game represents a
type Game struct {
	x int
	y int
	//board[y][x]
	board        [][]int
	piece        Piece
	pieceRotated Piece
	clock        *time.Timer
}

// helper function to initialize a new board
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

// NewGame returns a pointer to anew instantiation of Game
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
	if g.movePiece(down) {
		g.resetClock()
	} else {
		g.lockPiece()
		g.resetClock()
	}

}

func (g *Game) resetClock() {
	g.clock.Reset(gameSpeed)
}

// pieceFits checks returns a boolean indicating wether or not a piece fits in a specified location
func (g *Game) pieceFits(x int, y int) bool {
	for i := 0; i < numberSquares; i++ {
		squareX := x + g.pieceRotated.deltaX[i]
		squareY := y + g.pieceRotated.deltaY[i]

		// The piece collides with border
		if squareX < 0 || boardWidth <= squareX || boardHeight <= squareY {
			return false
		}

		// The piece collides with other piece
		if squareY >= 0 && g.board[squareY][squareX] < 0 {
			return false
		}
	}

	return true
}

// movePiece tries to move a piece in the specified direction (dir=[left, down, right])
// returns a boolean indicating if the operation was successfull
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

// TODO: check for bug if piece collides with other after roatation
// rotate rotates a piece
func (g *Game) rotatePiece() bool {
	g.pieceRotated.rotate()
	if !g.pieceFits(g.x, g.y) {

		// reverse rotation
		for i := 0; i < 3; i++ {
			g.pieceRotated.rotate()
		}
		return false
	}

	g.deletePiece()
	g.piece.rotate()
	g.placePiece()

	return true
}

// TODO: checkRow()
// checkRow checks if there is a complete row and deletes it

// updateBoard updates the coordinates occupied by the currently active piece with the specified value
func (g *Game) updateBoard(value int) {
	for i := 0; i < numberSquares; i++ {
		x := g.x + g.piece.deltaX[i]
		y := g.y + g.piece.deltaY[i]

		if 0 <= y {
			g.board[y][x] = value
		}
	}
}

// deletePiece removes the currently active piece from the board
func (g *Game) deletePiece() {
	g.updateBoard(0)
}

// placePiece writes the pieces color value to the board at the pieces coordinates
func (g *Game) placePiece() {
	g.updateBoard(g.piece.color)
}

// spawnRandomPiece spawns a new piece and sets it to be the currently active piece
func (g *Game) spawnRandomPiece() {
	color := rand.Int()%numberPieces + 1
	g.piece = pieces[color]
	g.pieceRotated = deepCopy(pieces[color])
	g.x = boardWidth / 2
	g.y = 0
	g.placePiece()
}

// lockPiece locks a piece in its current position only if piece touches the bottom margin
// returns a boolean indication whether or not the operation was successfull
func (g *Game) lockPiece() bool {
	if !g.pieceFits(g.x, g.y+1) {
		g.updateBoard(-g.piece.color)
		g.spawnRandomPiece()
		return true
	}
	return false
}
