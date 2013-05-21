package main

import "fmt"

type Player struct {
	active bool
	dirty  bool
	score  uint32
	target Vec
	ent    EntId
}

var numPlayers = 0
var players = make([]Player, 0, 10)

func incMaxPlayers() {
	players = append(players, Player{})
}

func login(id PlayerId) {
	if numPlayers > 4 {
		return
	}
	numPlayers++

	turretId := newEntId()

	fmt.Println(turretId)
	active[turretId] = true
	model[turretId] = Tank
	pos[turretId] = Vec{50, 100, 0}
	size[turretId] = Vec{40, 4, 4}
	rot[turretId] = 0

	entId := newEntId()

	active[entId] = true
	model[entId] = Tank
	pos[entId] = Vec{50, 100, 0}
	size[entId] = Vec{20, 16, 5}
	vel[entId] = Vec{}
	rot[entId] = 0
	child[entId] = turretId

	players[id].active = true
	players[id].target = Vec{1, 0, 0}
	players[id].score = 0
	players[id].ent = entId
}

func disconnect(id PlayerId) {
	if players[id].active == false {
		return
	}

	numPlayers--

	entId := players[id].ent
	turretId := child[entId]

	active[entId] = false
	active[turretId] = false

	players[id].active = false

	freeEntId(turretId)
	freeEntId(entId)

	freePlayerId(id)
}
