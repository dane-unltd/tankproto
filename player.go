package main

import (
	"fmt"
	"github.com/dane-unltd/engine/entstate"
	"github.com/dane-unltd/engine/math3"
)

type Player struct {
	active bool
	score  uint32
	target math3.Vec
	ent    entstate.EntId
}

var numPlayers = 0
var players = make([]Player, 0, 10)

func incMaxPlayers() {
	players = append(players, Player{})
}

func login(id PlayerId) {
	fmt.Println("login", id, numPlayers)
	if numPlayers > 4 {
		return
	}
	numPlayers++

	turretId := entstate.New()

	fmt.Println(turretId)
	active[turretId] = true
	model[turretId] = Tank
	pos[turretId] = math3.Vec{50, 100, 0}
	size[turretId] = math3.Vec{40, 4, 4}
	rot[turretId] = 0

	entId := entstate.New()

	active[entId] = true
	model[entId] = Tank
	pos[entId] = math3.Vec{50, 100, 0}
	size[entId] = math3.Vec{20, 16, 5}
	vel[entId] = math3.Vec{}
	rot[entId] = 0
	child[entId] = turretId
	collType[entId] = CollUnit

	players[id].active = true
	players[id].target = math3.Vec{1, 0, 0}
	players[id].score = 0
	players[id].ent = entId
}

func disconnect(id PlayerId) {
	fmt.Println("disconnect", id, players[id])
	if players[id].active == false {
		return
	}

	numPlayers--

	entId := players[id].ent
	turretId := child[entId]

	entstate.Delete(turretId)
	entstate.Delete(entId)

	players[id].active = false
}
