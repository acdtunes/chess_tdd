package chess

import (
	"fmt"
	"strconv"
)

type Board struct {
	size           int
	pos            string
	piece          string
}

func (b *Board) PlacePiece(piece string, pos string) {
	b.pos = pos
	b.piece = piece
}

func (b *Board) Moves() []string {
	if b.piece == "R" {
		return b.straightMoves()
	}
	if b.piece == "B" {
		return b.diagonalMoves()
	}
	if b.piece == "Q" {
		return b.straightAndDiagonalMoves()
	}
	if b.piece == "K" {
		return b.knightMoves()
	}

	return []string{}
}

func (b *Board) knightMoves() []string {
	var moves []string

	if b.pieceCol()-2 >= 1 {
		if b.pieceLine()+1 <= b.size {
			moves = append(moves, getPosition(b.pieceCol()-2, b.pieceLine()+1))
		}
		if b.pieceLine()-1 >= 1 {
			moves = append(moves, getPosition(b.pieceCol()-2, b.pieceLine()-1))
		}
	}
	if b.pieceCol()-1 >= 1 {
		if b.pieceLine()+2 <= b.size {
			moves = append(moves, getPosition(b.pieceCol()-1, b.pieceLine()+2))
		}
		if b.pieceLine()-2 >= 1 {
			moves = append(moves, getPosition(b.pieceCol()-1, b.pieceLine()-2))
		}
	}
	if b.pieceCol()+2 <= b.size {
		if b.pieceLine()+1 <= b.size {
			moves = append(moves, getPosition(b.pieceCol()+2, b.pieceLine()+1))
		}
		if b.pieceLine()-1 >= 1 {
			moves = append(moves, getPosition(b.pieceCol()+2, b.pieceLine()-1))
		}
	}
	if b.pieceCol()+1 <= b.size {
		if b.pieceLine()+2 <= b.size {
			moves = append(moves, getPosition(b.pieceCol()+1, b.pieceLine()+2))
		}
		if b.pieceLine()-2 >= 1 {
			moves = append(moves, getPosition(b.pieceCol()+1, b.pieceLine()-2))
		}
	}

	return moves
}

func (b *Board) straightAndDiagonalMoves() []string {
	moves := b.straightMoves()
	moves = append(moves, b.diagonalMoves()...)
	return moves
}

func (b *Board) straightMoves() []string {
	var moves []string

	for c := 1; c <= b.size; c++ {
		if !b.pieceIsOnSamePosition(c, b.pieceLine()) {
			moves = append(moves, getPosition(c, b.pieceLine()))
		}
	}
	for l := 1; l <= b.size; l++ {
		if !b.pieceIsOnSamePosition(b.pieceCol(), l) {
			moves = append(moves, getPosition(b.pieceCol(), l))
		}
	}

	return moves
}

func (b *Board) diagonalMoves() []string {
	var moves []string

	for c := b.pieceCol() + 1; c <= b.size; c++ {
		upper := b.upperLine(c)
		if b.isInBounds(upper) {
			moves = append(moves, getPosition(c, upper))
		}

		lower := b.lowerLine(c)
		if b.isInBounds(lower) {
			moves = append(moves, getPosition(c, lower))
		}
	}

	for c := b.pieceCol() - 1; c >= 1; c-- {
		upper := b.upperLine(c)
		if b.isInBounds(upper) {
			moves = append(moves, getPosition(c, upper))
		}

		lower := b.lowerLine(c)
		if b.isInBounds(lower) {
			moves = append(moves, getPosition(c, lower))
		}
	}

	return moves
}

func (b *Board) lowerLine(c int) int {
	return b.pieceLine() - (b.pieceCol() - c)
}

func (b *Board) upperLine(c int) int {
	return b.pieceLine() + (b.pieceCol() - c)
}

func (b *Board) isInBounds(x int) bool {
	return x >= 1 && x <= b.size
}

func (b *Board) pieceLine() int {
	line, _ := strconv.Atoi(string(b.pos[1]))
	return line
}

func (b *Board) pieceCol() int {
	return int(b.pos[0]) - 96
}

func (b *Board) pieceIsOnSamePosition(c int, l int) bool {
	return getPosition(c, l) == b.pos
}

func getPosition(c int, l int) string {
	return fmt.Sprintf("%s%d", string(c+96), l)
}

func NewBoardOfSize(size int) *Board {
	b := &Board{size: size}

	return b
}

