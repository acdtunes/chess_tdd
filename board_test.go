package chess

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Suite")
}

var _ = Describe("Chess", func() {
	Context("Queen", func() {
		Context("1x1", func() {
			Context("Queen on a1", func() {
				It("returns empty list", func() {
					board := NewBoardOfSize(1)
					board.PlacePiece("Q", "a1")
					Expect(board.Moves()).To(HaveLen(0))
				})
			})
		})
		Context("2x2", func() {
			Context("Queen on a1", func() {
				It("returns correct positions", func() {
					board := NewBoardOfSize(2)
					board.PlacePiece("Q", "a1")
					Expect(board.Moves()).To(ConsistOf("a2", "b1", "b2"))
				})
			})
			Context("Queen on a2", func() {
				It("returns correct positions", func() {
					board := NewBoardOfSize(2)
					board.PlacePiece("Q", "a2")
					Expect(board.Moves()).To(ConsistOf("a1", "b1", "b2"))
				})
			})
			Context("Queen on b1", func() {
				It("returns correct positions", func() {
					board := NewBoardOfSize(2)
					board.PlacePiece("Q", "b1")
					Expect(board.Moves()).To(ConsistOf("a1", "a2", "b2"))
				})
			})
			Context("Queen on b2", func() {
				It("returns correct positions", func() {
					board := NewBoardOfSize(2)
					board.PlacePiece("Q", "b2")
					Expect(board.Moves()).To(ConsistOf("a1", "a2", "b1"))
				})
			})
		})
		Context("3x3", func() {
			Context("Queen on a1", func() {
				It("returns correct positions", func() {
					board := NewBoardOfSize(3)
					board.PlacePiece("Q", "a1")
					Expect(board.Moves()).To(ConsistOf("a2", "a3", "b1", "c1", "b2", "c3"))
				})
			})
		})
		Context("8x8", func() {
			Context("Queen on b3", func() {
				It("returns correct positions", func() {
					board := NewBoardOfSize(8)
					board.PlacePiece("Q", "b3")
					Expect(board.Moves()).To(ConsistOf("b2", "b1", "b4", "b5", "b6", "b7", "b8",
						"a3", "c3", "d3", "e3", "f3", "g3", "h3",
						"a4", "a2", "c2", "d1", "c4", "d5", "e6", "f7", "g8"))
				})
			})
		})
	})
	Context("Rook", func() {
		Context("1x1", func() {
			Context("Rook on a1", func() {
				It("returns empty list", func() {
					board := NewBoardOfSize(1)
					board.PlacePiece("R", "a1")
					Expect(board.Moves()).To(HaveLen(0))
				})
			})
		})
		Context("2x2", func() {
			Context("Rook on a1", func() {
				It("returns correct positions", func() {
					board := NewBoardOfSize(2)
					board.PlacePiece("R", "a1")
					Expect(board.Moves()).To(ConsistOf("b1", "a2"))
				})
			})
			Context("Rook on a2", func() {
				It("returns correct positions", func() {
					board := NewBoardOfSize(2)
					board.PlacePiece("R", "a2")
					Expect(board.Moves()).To(ConsistOf("a1", "b2"))
				})
			})
			Context("Rook on b1", func() {
				It("returns correct positions", func() {
					board := NewBoardOfSize(2)
					board.PlacePiece("R", "b1")
					Expect(board.Moves()).To(ConsistOf("a1", "b2"))
				})
			})
			Context("Rook on b2", func() {
				It("returns correct positions", func() {
					board := NewBoardOfSize(2)
					board.PlacePiece("R", "b2")
					Expect(board.Moves()).To(ConsistOf("a2", "b1"))
				})
			})
		})
		Context("3x3", func() {
			Context("Rook on a1", func() {
				It("returns correct positions", func() {
					board := NewBoardOfSize(3)
					board.PlacePiece("R", "a1")
					Expect(board.Moves()).To(ConsistOf("a2", "a3", "b1", "c1"))
				})
			})
		})
		Context("8x8", func() {
			Context("Rook on b3", func() {
				It("returns correct positions", func() {
					board := NewBoardOfSize(8)
					board.PlacePiece("R", "b3")
					Expect(board.Moves()).To(ConsistOf("b2", "b1", "b4", "b5", "b6", "b7", "b8",
						"a3", "c3", "d3", "e3", "f3", "g3", "h3"))
				})
			})
		})
	})
	Context("Bishop", func() {
		Context("1x1", func() {
			Context("Bishop on a1", func() {
				It("returns empty list", func() {
					board := NewBoardOfSize(1)
					board.PlacePiece("B", "a1")
					Expect(board.Moves()).To(HaveLen(0))
				})
			})
		})
		Context("2x2", func() {
			Context("Bishop on a1", func() {
				It("returns correct positions", func() {
					board := NewBoardOfSize(2)
					board.PlacePiece("B", "a1")
					Expect(board.Moves()).To(ConsistOf("b2"))
				})
			})
			Context("Bishop on a2", func() {
				It("returns correct positions", func() {
					board := NewBoardOfSize(2)
					board.PlacePiece("B", "a2")
					Expect(board.Moves()).To(ConsistOf("b1"))
				})
			})
			Context("Bishop on b1", func() {
				It("returns correct positions", func() {
					board := NewBoardOfSize(2)
					board.PlacePiece("B", "b1")
					Expect(board.Moves()).To(ConsistOf("a2"))
				})
			})
			Context("Bishop on b2", func() {
				It("returns correct positions", func() {
					board := NewBoardOfSize(2)
					board.PlacePiece("B", "b2")
					Expect(board.Moves()).To(ConsistOf("a1"))
				})
			})
		})
		Context("3x3", func() {
			Context("Bishop on a1", func() {
				It("returns correct positions", func() {
					board := NewBoardOfSize(3)
					board.PlacePiece("B", "a1")
					Expect(board.Moves()).To(ConsistOf("b2", "c3"))
				})
			})
		})
		Context("8x8", func() {
			Context("Bishop on b3", func() {
				It("returns correct positions", func() {
					board := NewBoardOfSize(8)
					board.PlacePiece("B", "b3")
					Expect(board.Moves()).To(ConsistOf("a4", "a2", "c2", "d1", "c4", "d5", "e6", "f7", "g8"))
				})
			})
		})
		Context("8x8", func() {
			Context("Bishop on d4", func() {
				It("returns correct positions", func() {
					board := NewBoardOfSize(8)
					board.PlacePiece("B", "d4")
					Expect(board.Moves()).To(ConsistOf("e5", "f6", "g7", "h8", "c3", "b2", "a1",
						"e3", "f2", "g1", "c5", "b6", "a7"))
				})
			})
		})
	})
	Context("Knight", func() {
		Context("8x8", func() {
			Context("Knight on b1", func() {
				It("returns correct positions", func() {
					board := NewBoardOfSize(8)
					board.PlacePiece("K", "b1")
					Expect(board.Moves()).To(ConsistOf("a3", "c3", "d2"))
				})
			})
			Context("Knight on c4", func() {
				It("returns correct positions", func() {
					board := NewBoardOfSize(8)
					board.PlacePiece("K", "c4")
					Expect(board.Moves()).To(ConsistOf("e5", "d6", "d2", "e3", "b2", "a3", "a5", "b6"))
				})
			})
		})
	})
})
