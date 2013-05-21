package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

type Entity map[CompId]int

type StateComp interface {
	Copyer
}

type Copyer interface {
	Copy(src Copyer)
}

//world state
var tileMapOld = NewTileMap(10, 10)
var tileMap = NewTileMap(10, 10)

//entity state
var comp = make([]BitMask, 0, 10)

var active = make([]bool, 0, 10)
var child = make([]EntId, 0, 10)

var pos = make([]Vec, 0, 10)
var posOld = make([]Vec, 0, 10)

var size = make([]Vec, 0, 10)
var sizeOld = make([]Vec, 0, 10)

var vel = make([]Vec, 0, 10)
var velOld = make([]Vec, 0, 10)

var rot = make([]float64, 0, 10)
var rotOld = make([]float64, 0, 10)

var model = make([]ModelId, 0, 10)
var modelOld = make([]ModelId, 0, 10)

//entity mangagement

//initializing map
func init() {
	fmt.Println("initializing")

	for i := 0; i < 10; i++ {
		tileMap.Set(0, i, 1)
		tileMap.Set(i, 0, 1)
		tileMap.Set(9, i, 1)
		tileMap.Set(i, 9, 1)
	}
}

func incMaxEnts() {
	model = append(model, 0)
	modelOld = append(model, 0)
	pos = append(pos, Vec{})
	posOld = append(posOld, Vec{})
	vel = append(vel, Vec{})
	velOld = append(velOld, Vec{})
	size = append(size, Vec{})
	sizeOld = append(sizeOld, Vec{})
	rot = append(rot, 0)
	rotOld = append(rotOld, 0)

	comp = append(comp, NewBitMask(5))
	active = append(active, false)
	child = append(child, 0)
}

func copyState() {
	copy(modelOld, model)
	copy(velOld, vel)
	copy(posOld, pos)
	copy(sizeOld, size)
	copy(rotOld, rot)

	tileMapOld.Copy(tileMap)
}

func serialize(buf io.Writer, serAll bool) {
	nEnts := 0
	for _, act := range active {
		if act {
			nEnts++
		}
	}
	binary.Write(buf, binary.LittleEndian, uint32(nEnts))
	for id, act := range active {
		if !act {
			continue
		}
		binary.Write(buf, binary.LittleEndian, EntId(id))

		bitMask := NewBitMask(4)
		bufTemp := &bytes.Buffer{}
		if serAll || model[id] != modelOld[id] {
			bitMask.Set(0)
			binary.Write(bufTemp, binary.LittleEndian, model[id])
		}
		if serAll || !pos[id].Equals(&posOld[id]) {
			bitMask.Set(1)
			binary.Write(bufTemp, binary.LittleEndian, &pos[id])
		}
		if serAll || !size[id].Equals(&sizeOld[id]) {
			bitMask.Set(2)
			binary.Write(bufTemp, binary.LittleEndian, &size[id])
		}
		if serAll || rot[id] != rotOld[id] {
			bitMask.Set(3)
			binary.Write(bufTemp, binary.LittleEndian, rot[id])
		}
		buf.Write(bitMask)
		buf.Write(bufTemp.Bytes())
	}

	tileMap.Serialize(buf, serAll, tileMapOld)
}
