package main

import "main/core"

func main() {
	game := core.NewSapperGame(8, 10)
	game.Initialize()
	game.ClickOnCell(core.Cell{X: 0, Y: 0})
}
