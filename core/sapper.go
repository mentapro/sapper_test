package core

import (
	"fmt"
	"math/rand"
	"time"
)

type sapperGame struct {
	board map[Cell]cellState
	n, k  int
}

type cellState struct {
	isBlackHole              bool
	adjacentBlackHoleCounter int
	isOpen                   bool
}

// NewSapperGame - constructor for creating game state
func NewSapperGame(n, k int) Game {
	game := sapperGame{
		board: make(map[Cell]cellState),
		n:     n,
		k:     k,
	}
	return game
}

// Initialize - Implementing an interface method for initial initialization
func (game sapperGame) Initialize() {
	fmt.Println("Initial board...")
	game.initializeBoard()

	fmt.Println()
	fmt.Println("Generating black holes...")
	game.generateBlackHoles()

	fmt.Println()
	fmt.Println("Calculating adjacent black hole counter...")
	game.calculateAdjacentBlackHoleCounter()
	game.printBoard()
}

// ClickOnCell - Implementing an interface method for clicking on cell
func (game sapperGame) ClickOnCell(targetCell Cell) {
	fmt.Println()
	fmt.Println("Handle some click...")
	game.handleCellClick(targetCell)
	game.printBoard()
}

func (game sapperGame) printBoard() {
	fmt.Println("== B O A R D  S T A T E ==")
	for yi := 0; yi < game.n; yi++ {
		for xi := 0; xi < game.n; xi++ {
			if xi == 0 {
				fmt.Print("|")
			}
			if cell := game.board[Cell{X: xi, Y: yi}]; cell.isBlackHole {
				fmt.Print("_")
			} else if cell.isOpen {
				fmt.Print(cell.adjacentBlackHoleCounter)
			} else {
				fmt.Print("_")
			}
			fmt.Print("|")
		}
		fmt.Println()
	}
}

func (game sapperGame) initializeBoard() {
	for yi := 0; yi < game.n; yi++ {
		for xi := 0; xi < game.n; xi++ {
			game.board[Cell{X: xi, Y: yi}] = cellState{
				isBlackHole:              false,
				adjacentBlackHoleCounter: -1,
				isOpen:                   false,
			}
		}
	}
}

func (game sapperGame) generateBlackHoles() {
	rand.Seed(time.Now().UnixNano())
	// we need to keep track of remaining empty cells because we want to generate exact count of black holes
	remainingCoordinates := make([]Cell, 0, game.n*game.n)
	for yi := 0; yi < game.n; yi++ {
		for xi := 0; xi < game.n; xi++ {
			remainingCoordinates = append(remainingCoordinates, Cell{X: xi, Y: yi})
		}
	}
	for i := 0; i < game.k; i++ {
		randomCellIndex := rand.Intn(len(remainingCoordinates))
		randomCellCoordinate := remainingCoordinates[randomCellIndex]
		cellState := game.board[randomCellCoordinate]
		cellState.isBlackHole = true
		game.board[randomCellCoordinate] = cellState
		// remove used cell to avoid repetition in randomizing coordinates
		remainingCoordinates = append(remainingCoordinates[:randomCellIndex], remainingCoordinates[randomCellIndex+1:]...)
	}
}

func (game sapperGame) calculateAdjacentBlackHoleCounter() {
	for yi := 0; yi < game.n; yi++ {
		for xi := 0; xi < game.n; xi++ {
			targetCell := game.board[Cell{X: xi, Y: yi}]
			if targetCell.isBlackHole {
				continue
			}

			adjacentBlackHoleCounter := 0
			adjacentCellsCoordinates := getRoundaboutCells(Cell{X: xi, Y: yi})
			for _, adjacentCellCoord := range adjacentCellsCoordinates {
				adjacentCell, ok := game.board[adjacentCellCoord]
				if ok && adjacentCell.isBlackHole {
					adjacentBlackHoleCounter++
				}
			}

			targetCell.adjacentBlackHoleCounter = adjacentBlackHoleCounter
			game.board[Cell{X: xi, Y: yi}] = targetCell
		}
	}
}

func (game sapperGame) handleCellClick(clickedCell Cell) {
	cellState := game.board[clickedCell]
	if cellState.isOpen {
		fmt.Printf("Cell %v is opened already. Do nothing\n", clickedCell)
		return
	}

	if cellState.isBlackHole {
		fmt.Printf("Cell %v is black hole. You lose\n", clickedCell)
		return
	}

	game.openCell(clickedCell)
}

func (game sapperGame) openCell(coordinate Cell) {
	targetCell, ok := game.board[coordinate]
	if !ok || targetCell.isOpen {
		return
	}

	targetCell.isOpen = true
	game.board[coordinate] = targetCell

	// if counter is zero - then we should open all adjacent cells
	if targetCell.adjacentBlackHoleCounter == 0 {
		adjacentCellsCoordinates := getRoundaboutCells(coordinate)
		for _, adjacentCellCoord := range adjacentCellsCoordinates {
			game.openCell(adjacentCellCoord)
		}
	}
}

func getRoundaboutCells(coordinate Cell) []Cell {
	cells := make([]Cell, 0, 8)

	cells = append(cells, Cell{X: coordinate.X + 1, Y: coordinate.Y - 1})
	cells = append(cells, Cell{X: coordinate.X + 1, Y: coordinate.Y})
	cells = append(cells, Cell{X: coordinate.X + 1, Y: coordinate.Y + 1})

	cells = append(cells, Cell{X: coordinate.X - 1, Y: coordinate.Y - 1})
	cells = append(cells, Cell{X: coordinate.X - 1, Y: coordinate.Y})
	cells = append(cells, Cell{X: coordinate.X - 1, Y: coordinate.Y + 1})

	cells = append(cells, Cell{X: coordinate.X, Y: coordinate.Y - 1})
	cells = append(cells, Cell{X: coordinate.X, Y: coordinate.Y + 1})

	return cells
}
