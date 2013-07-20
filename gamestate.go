package main

import (
	"fmt"
	"github.com/dane-unltd/engine/entstate"
	"io"
)

const (
	PosComp entstate.CompId = iota
	SizeComp
	RotComp
	VelComp
	ModelComp
	ActiveComp
	ChildComp
	CollTypeComp
)

//world state
var tileMapOld = NewTileMap(10, 10)
var tileMap = NewTileMap(10, 10)

//entity state
var active = entstate.NewBoolState()
var child = entstate.NewEntIdState()
var collType = NewCollTypeState()

var model = NewModelState()
var pos = entstate.NewVecState()
var size = entstate.NewVecState()
var rot = entstate.NewFloat64State()
var vel = entstate.NewVecState()

func init() {
	fmt.Println("registering components")
	entstate.RegisterComp(ActiveComp, false, &active)
	entstate.RegisterComp(ChildComp, false, &child)
	entstate.RegisterComp(VelComp, false, &vel)
	entstate.RegisterComp(CollTypeComp, false, &collType)

	entstate.RegisterComp(ModelComp, true, &model)
	entstate.RegisterComp(PosComp, true, &pos)
	entstate.RegisterComp(SizeComp, true, &size)
	entstate.RegisterComp(RotComp, true, &rot)

	//initializing map
	fmt.Println("initializing")

	for i := 0; i < 10; i++ {
		tileMap.Set(0, i, 1)
		tileMap.Set(i, 0, 1)
		tileMap.Set(9, i, 1)
		tileMap.Set(i, 9, 1)
	}
}

func serialize(buf io.Writer, serAll bool) {
	entstate.Serialize(buf, serAll, active)
	tileMap.Serialize(buf, serAll, tileMapOld)
}

func copyState() {
	entstate.CopyState()
	tileMapOld.Copy(tileMap)
}
