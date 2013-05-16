package main

type ModelId uint32
type CompId uint32

const (
	TransComp  CompId = 1
	PosComp    CompId = 2
	SizeComp   CompId = 3
	RotComp    CompId = 4
	VelComp    CompId = 5
	ModelComp  CompId = 6
	PlayerComp CompId = 7
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
