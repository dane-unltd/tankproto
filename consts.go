package main

type ModelId uint32

const (
	PosComp CompId = iota
	SizeComp
	RotComp
	VelComp
	ModelComp
	ActiveComp
	ChildComp

	NumComps
)

const (
	Tank   ModelId = 1
	Bullet ModelId = 2
)

const (
	Forward  Action = 0
	Backward Action = 1
	Left     Action = 2
	Right    Action = 3
)

const TileSize = 20
