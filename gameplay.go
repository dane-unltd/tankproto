package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

type IdList []EntId
type IdMap map[EntId]struct{}
type Entity map[CompId]int

type CompState interface {
	Remove(i int)
	Serialize(i int)
}

//game state
var tileMapOld = NewTileMap(10, 10)
var tileMap = NewTileMap(10, 10)

var players = NewPlayerState()
var playersOld = NewPlayerState()

var pos = NewVecState(PosComp)
var posOld = NewVecState(PosComp)

var size = NewVecState(SizeComp)
var sizeOld = NewVecState(SizeComp)

var vel = NewVecState(VelComp)
var velOld = NewVecState(VelComp)

var rot = NewFloat64State(RotComp)
var rotOld = NewFloat64State(RotComp)

var model = NewModelState()
var modelOld = NewModelState()

//entity mangagement
var ents = make([]Entity, 0, 100)
var entsOld = make([]Entity, 0, 100)

var entIds = make(IdMap)

var comps = make(map[CompId]CompState)

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

func copyState() {
	modelOld.Copy(model)
	velOld.Copy(vel)
	posOld.Copy(pos)
	sizeOld.Copy(size)
	rotOld.Copy(rot)
	playersOld.Copy(players)

	tileMapOld.Copy(tileMap)
}

func updateSimulation() {
	processInput()
	collisionCheck()
	move()
	placeTurrets()
}

func move() {
	for i := range vel.v {
		posIx := ents[vel.ids[i]][TransComp]
		pos.v[posIx].Add(&pos.v[posIx], &vel.v[i])
	}
}

func placeTurrets() {
	for i, id := range players.ids {
		childId := players.turret[i]
		posIx := ents[id][TransComp]
		childPosIx := ents[childId][TransComp]
		offset := Vec{}
		offset.Scale(4, &players.target[i])
		pos.v[childPosIx].Add(&pos.v[posIx], &offset)
		rot.v[childPosIx] = math.Atan2(offset[1], offset[0])
	}
}

func processInput() {
	for i, pl := range players.ids {
		newVel := 0.0
		rotIx := ents[pl][RotComp]
		if active(pl, Forward) {
			newVel += 5
		}
		if active(pl, Backward) {
			newVel -= 5
		}

		if active(pl, Left) {
			rot.v[rotIx] += 0.1
		}
		if active(pl, Right) {
			rot.v[rotIx] -= 0.1
		}

		velIx := ents[pl][VelComp]
		vel.v[velIx][0] = newVel * math.Cos(rot.v[rotIx])
		vel.v[velIx][1] = newVel * math.Sin(rot.v[rotIx])

		posIx := ents[pl][PosComp]
		players.target[i].Sub(target(pl), &pos.v[posIx])
		players.target[i].Normalize(&players.target[i])
	}
}

func collisionCheck() {
	checkMap()
}

func checkMap() {
	tileSize := Vec{20, 20, 20}
	for velIx, id := range vel.ids {
		posIx := ents[id][PosComp]
		sizeIx := ents[id][SizeComp]
		px := math.Floor(pos.v[posIx][0] / 20)
		py := math.Floor(pos.v[posIx][1] / 20)

		r := size.v[sizeIx][0] / 2

		rt := math.Ceil(r / 20)
		for tx := px - rt; tx <= px+rt; tx++ {
			for ty := py - rt; ty <= py+rt; ty++ {
				if tx < 0 || ty < 0 {
					continue
				}
				if tx > 9 || ty > 9 {
					continue
				}
				if tileMap.At(int(tx), int(ty)) == 0 {
					continue
				}
				tilePos := Vec{tx*20 + 10, ty*20 + 10, 0}

				v := Vec{}
				v.Sub(&pos.v[posIx], &tilePos)
				v.Clamp(&tileSize)

				d := Vec{}
				d.Sub(&pos.v[posIx], &tilePos)
				d.Sub(&d, &v)

				dist := math.Sqrt(d.Nrm2Sq())
				vProj := Dot(&vel.v[velIx], &d)
				vProj /= dist

				remove := dist - r + vProj
				if remove < 0 {
					if dist < r {
						dPos := Vec{}
						dPos.Scale(r/dist-1, &d)
						pos.v[posIx].Add(&pos.v[posIx], &dPos)
					}

					d.Scale(-remove/dist, &d)
					vel.v[velIx].Add(&vel.v[velIx], &d)
				}
			}
		}
	}
}

func login(id EntId) {
	if len(players.ids) > 4 {
		return
	}
	entIds[id] = struct{}{}

	childId := newId()
	child := make(map[CompId]int)
	ents[childId] = child
	entIds[childId] = struct{}{}

	child[ModelComp] = model.Append(childId, Tank)
	child[PosComp] = pos.Append(childId, Vec{50, 100, 0})
	child[SizeComp] = size.Append(childId, Vec{40, 4, 4})
	child[RotComp] = rot.Append(childId, 0)

	ent := make(map[CompId]int)
	ents[id] = ent

	ent[ModelComp] = model.Append(id, Tank)
	ent[PosComp] = pos.Append(id, Vec{50, 100, 0})
	ent[SizeComp] = size.Append(id, Vec{20, 16, 5})
	ent[RotComp] = rot.Append(id, 0)
	ent[VelComp] = vel.Append(id, Vec{})
	ent[PlayerComp] = players.Append(id, Vec{1, 0, 0}, 0, childId)
}

func startGame() {
}

func disconnect(id EntId) {
	if ents[id] == nil {
		return
	}
	ent := ents[id]
	childId := players.turret[ent[PlayerComp]]
	child := ents[childId]
	model.Remove(child[ModelComp])
	pos.Remove(child[PosComp])
	size.Remove(child[SizeComp])
	rot.Remove(child[RotComp])
	delete(entIds, childId)
	ents[childId] = nil

	model.Remove(ent[ModelComp])
	pos.Remove(ent[PosComp])
	size.Remove(ent[SizeComp])
	rot.Remove(ent[RotComp])
	players.Remove(ent[PlayerComp])
	vel.Remove(ent[VelComp])
	delete(entIds, id)
	ents[id] = nil
}

func stopGame() {
}

func serialize(buf io.Writer, serAll bool) {

	bitMask := NewBitMask(4)
	bufTemp := &bytes.Buffer{}

	binary.Write(buf, binary.LittleEndian, uint32(len(entIds)))
	for id := range entIds {
		binary.Write(buf, binary.LittleEndian, id)
		ent := ents[id]
		entOld := entsOld[id]
		doDelta := serAll
		if entOld == nil {
			doDelta = false
		}

		modelIx := ent[ModelComp]
		modelIxOld := entOld[ModelComp]

		if doDelta || model.m[modelIx] != modelOld.m[modelIxOld] {
			bitMask.Set(0)
			binary.Write(bufTemp, binary.LittleEndian,
				model.m[modelIx])
		}

		posIx := ent[PosComp]
		posIxOld := entOld[PosComp]

		if doDelta || !pos.v[posIx].Equals(&posOld.v[posIxOld]) {
			bitMask.Set(1)
			binary.Write(bufTemp, binary.LittleEndian,
				&pos.v[posIx])
		}

		sizeIx := ent[SizeComp]
		sizeIxOld := entOld[SizeComp]

		if doDelta || !size.v[sizeIx].Equals(&sizeOld.v[sizeIxOld]) {
			bitMask.Set(2)
			binary.Write(bufTemp, binary.LittleEndian,
				&size.v[sizeIx])
		}

		rotIx := ent[RotComp]
		rotIxOld := entOld[RotComp]

		if doDelta || rot.v[rotIx] != rotOld.v[rotIxOld] {
			bitMask.Set(3)
			binary.Write(bufTemp, binary.LittleEndian,
				rot.v[rotIx])
		}
	}
	buf.Write(bitMask)
	buf.Write(bufTemp.Bytes())

	tileMap.Serialize(buf, serAll, tileMapOld)
}
