package core

import (
	"testing"
)

func TestBoardInitialization(t *testing.T) {
	t.Run("test that board is initialized", func(t *testing.T) {
		// arrange
		game := newTestGame(8, 10)

		// act
		game.initializeBoard()

		// assert
		for yi := 0; yi < game.n; yi++ {
			for xi := 0; xi < game.n; xi++ {
				coord := Cell{X: xi, Y: yi}
				cell, ok := game.board[coord]
				if !ok {
					t.Errorf("cell %v is not initialized", coord)
					continue
				}

				if cell.adjacentBlackHoleCounter != -1 || cell.isOpen || cell.isBlackHole {
					t.Errorf("cell %v is initialized incorrectly", coord)
					continue
				}
			}
		}
	})
}

func TestBlackHoleGeneration(t *testing.T) {
	t.Run("test that black holes are generated with correct count", func(t *testing.T) {
		// arrange
		game := newTestGame(5, 10)
		game.initializeBoard()

		// act
		game.generateBlackHoles()

		// assert
		blackHoleCounter := 0
		for yi := 0; yi < game.n; yi++ {
			for xi := 0; xi < game.n; xi++ {
				coord := Cell{X: xi, Y: yi}
				cell := game.board[coord]
				if cell.isBlackHole {
					blackHoleCounter++
				}
			}
		}

		if blackHoleCounter != game.k {
			t.Errorf("got %d black holes, want %d black holes", blackHoleCounter, game.k)
		}
	})
}

func TestAdjacentBlackHoleCounter(t *testing.T) {
	t.Run("test that counter for adjacent black holes is correct", func(t *testing.T) {
		// arrange
		game := newTestGame(8, 8)
		game.initializeBoard()

		bh1 := Cell{X: 0, Y: 0}
		bh2 := Cell{X: 1, Y: 0}
		cellState1 := game.board[bh1]
		cellState1.isBlackHole = true
		cellState2 := game.board[bh2]
		cellState2.isBlackHole = true
		game.board[bh1] = cellState1
		game.board[bh2] = cellState2

		// act
		game.calculateAdjacentBlackHoleCounter()

		// assert
		checkCell1 := Cell{X: 2, Y: 0}
		checkCell2 := Cell{X: 0, Y: 1}
		cellCounter1 := game.board[checkCell1]
		cellCounter2 := game.board[checkCell2]

		if cellCounter1.adjacentBlackHoleCounter != 1 || cellCounter2.adjacentBlackHoleCounter != 2 {
			t.Errorf("expected that cell %v will have counter 1, and cell %v - counter 2", cellCounter1, cellCounter2)
		}
	})
}

func TestOpeningAdjacentZeroCells(t *testing.T) {
	t.Run("test that all adjacent cells near zero counter will be opened", func(t *testing.T) {
		// arrange
		game := newTestGame(8, 8)
		game.initializeBoard()

		bh1 := Cell{X: 2, Y: 0}
		cellState1 := game.board[bh1]
		cellState1.isBlackHole = true
		game.board[bh1] = cellState1
		game.calculateAdjacentBlackHoleCounter()

		// act
		game.openCell(Cell{X: 0, Y: 0})

		// assert
		checkCell1 := Cell{X: 0, Y: 1}
		checkCell2 := Cell{X: 0, Y: 2}
		cellCounter1 := game.board[checkCell1]
		cellCounter2 := game.board[checkCell2]

		if !cellCounter1.isOpen || !cellCounter2.isOpen {
			t.Errorf("expected that these cells should be opened: %v, %v", cellCounter1, cellCounter2)
		}
	})
}

func newTestGame(n, k int) sapperGame {
	game := sapperGame{
		board: make(map[Cell]cellState),
		n:     n,
		k:     k,
	}
	return game
}
