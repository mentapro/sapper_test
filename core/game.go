package core

type Game interface {
	Initialize()
	ClickOnCell(targetCell Cell)
}

type Cell struct {
	X int
	Y int
}
