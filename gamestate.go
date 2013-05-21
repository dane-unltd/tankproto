package main

import ()

type CompId uint32

type StateComp interface {
	Copyer
	Clone() StateComp
	Append()
	Val(id EntId) interface{}
	Equal(v interface{}, id EntId) bool
}

type Copyer interface {
	Copy(src Copyer)
}

//world state
var tileMapOld = NewTileMap(10, 10)
var tileMap = NewTileMap(10, 10)

//entity state
var comp = make([]BitMask, 0, 10)

var active = NewBoolState()
var child = NewEntIdState()

var model = NewModelState()
var pos = NewVecState()
var size = NewVecState()
var rot = NewFloat64State()
var vel = NewVecState()

//state collection
var stateComps = make([]StateComp, 0)
var networkedComps = make([]CompId, 0)
var oldStateComps = make([]StateComp, 0)

func init() {
	RegisterComp(ActiveComp, 0, &active)
	RegisterComp(ChildComp, 0, &child)
	RegisterComp(VelComp, 0, &vel)

	RegisterComp(ModelComp, 1, &model)
	RegisterComp(PosComp, 2, &pos)
	RegisterComp(SizeComp, 3, &size)
	RegisterComp(RotComp, 4, &rot)
}

func incMaxEnts() {
	for i := range stateComps {
		stateComps[i].Append()
	}
	for i := range oldStateComps {
		if oldStateComps[i] != nil {
			oldStateComps[i].Append()
		}
	}
}

func AddNetworkedComp(id CompId, pr int) {
	if len(networkedComps) < pr+1 {
		temp := make([]CompId, pr+1)
		copy(temp, networkedComps)
		networkedComps = temp
	}
	if networkedComps[pr] != 0 {
		panic("two components at same serialization step")
	}
	networkedComps[pr] = id
}

func RegisterComp(id CompId, pr int, sc StateComp) {
	if len(stateComps) < int(id)+1 {
		temp := make([]StateComp, id+1)
		copy(temp, stateComps)
		stateComps = temp

		temp = make([]StateComp, id+1)
		copy(temp, oldStateComps)
		oldStateComps = temp
	}
	if stateComps[id] != nil {
		panic("two components with same id")
	}

	stateComps[id] = sc
	oldSc := sc.Clone()
	oldStateComps[id] = oldSc

	if pr > 0 {
		AddNetworkedComp(id, pr-1)
	}
}

func copyState() {
	for i := range stateComps {
		if oldStateComps[i] != nil {
			oldStateComps[i].Copy(stateComps[i])
		}
	}
	tileMapOld.Copy(tileMap)
}
