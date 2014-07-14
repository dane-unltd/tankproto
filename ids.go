package main

import (
	"fmt"
	"github.com/dane-unltd/engine/entstate"
	"github.com/dane-unltd/engine/idgen"
)

type PlayerId uint32

type IdList []entstate.EntId
type IdMap map[entstate.EntId]struct{}

var plGen = idgen.New(incMaxPlayers)

func freePlayerId(id PlayerId) {
	fmt.Println("free", id)
	plGen.Free(uint32(id))
}

func newPlayerId() PlayerId {
	id := PlayerId(plGen.Next())
	fmt.Println("new", id)
	return id
}
