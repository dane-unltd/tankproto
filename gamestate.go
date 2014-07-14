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
var terrainOld = NewTerrain(10, 1)
var terrain = NewTerrain(10, 1)

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
}

func serialize(buf io.Writer, serAll bool) {
	entstate.Serialize(buf, serAll, active)
	terrain.Serialize(buf, serAll, terrainOld)
}

func copyState() {
	entstate.CopyState()
	terrainOld.Copy(terrain)
}
