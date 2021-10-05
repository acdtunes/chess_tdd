package chess

import (
	"fmt"
	"strconv"
)

type Board struct {
	size           int
	pos            string
	piece          string
	moves		   []string
}

func (b *Board) PlacePiece(piece string, pos string) {
	b.pos = pos
	b.piece = piece
}

func (b *Board) Moves() []string {
	if b.piece == "R" {
		b.addStraightMoves()
	}
	if b.piece == "B" {
		b.addDiagonalMoves()
	}
	if b.piece == "Q" {
		b.addStraightAndDiagonalMoves()
	}
	if b.piece == "K" {
		b.addKnightMoves()
	}

	return b.moves
}

func (b *Board) addStraightAndDiagonalMoves() {
	b.addStraightMoves()
	b.addDiagonalMoves()
}

func (b *Board) addStraightMoves() {
	for col := 1; col <= b.size; col++ {
		b.addMove(col, b.pieceLine())
	}
	for line := 1; line <= b.size; line++ {
		b.addMove(b.pieceCol(), line)
	}
}

func (b *Board) addDiagonalMoves() {
	for col := b.pieceCol() + 1; col <= b.size; col++ {
		b.addMove(col, b.diagonalAboveLine(col))
		b.addMove(col, b.diagonalBelowLine(col))
	}

	for col := b.pieceCol() - 1; col >= 1; col-- {
		b.addMove(col, b.diagonalAboveLine(col))
		b.addMove(col, b.diagonalBelowLine(col))
	}
}

func (b *Board) diagonalBelowLine(col int) int {
	return b.pieceLine() - (b.pieceCol() - col)
}

func (b *Board) diagonalAboveLine(col int) int {
	return b.pieceLine() + (b.pieceCol() - col)
}

func (b *Board) isInBounds(p int) bool {
	return p >= 1 && p <= b.size
}

func (b *Board) addKnightMoves() {
	b.addMove(b.pieceCol()-2, b.pieceLine()+1)
	b.addMove(b.pieceCol()-2, b.pieceLine()-1)
	b.addMove(b.pieceCol()+2, b.pieceLine()+1)
	b.addMove(b.pieceCol()+2, b.pieceLine()-1)
	b.addMove(b.pieceCol()-1, b.pieceLine()+2)
	b.addMove(b.pieceCol()-1, b.pieceLine()-2)
	b.addMove(b.pieceCol()+1, b.pieceLine()+2)
	b.addMove(b.pieceCol()+1, b.pieceLine()-2)
}

func (b *Board) pieceLine() int {
	line, _ := strconv.Atoi(string(b.pos[1]))
	return line
}

func (b *Board) pieceCol() int {
	return int(b.pos[0]) - 96
}

func (b *Board) addMove(col int, line int) {
	if !b.pieceIsOnSamePosition(col, line) && b.isInBounds(col) && b.isInBounds(line) {
		b.moves = append(b.moves, getPosition(col, line))
	}
}

func (b *Board) pieceIsOnSamePosition(col int, line int) bool {
	return getPosition(col, line) == b.pos
}

func getPosition(col int, line int) string {
	return fmt.Sprintf("%c%d", rune(col+96), line)
}

func NewBoardOfSize(size int) *Board {
	return &Board{size: size, moves: []string{}}
}

