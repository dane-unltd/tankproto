package main

type ModelId uint32
type CollType uint32

const (
	CollNone CollType = iota
	CollProjectile
	CollUnit
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
