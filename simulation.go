package main

import (
	"math"
)

func updateSimulation() {
	processInput()
	collisionCheck()
	move()
	placeTurrets()
}

func move() {
	for id, act := range active {
		if !act {
			continue
		}
		pos[id].Add(&pos[id], &vel[id])
	}
}

func processInput() {
	for id := range players {
		if !players[id].active {
			continue
		}
		pl := PlayerId(id)
		entId := players[pl].ent
		newVel := 0.0

		if Active(pl, Forward) {
			newVel += 5
		}
		if Active(pl, Backward) {
			newVel -= 5
		}

		if Active(pl, Left) {
			rot[entId] += 0.1
		}
		if Active(pl, Right) {
			rot[entId] -= 0.1
		}

		vel[entId][0] = newVel * math.Cos(rot[entId])
		vel[entId][1] = newVel * math.Sin(rot[entId])

		players[pl].target.Sub(target(pl), &pos[entId])
		players[pl].target.Normalize(&players[pl].target)
	}
}

func collisionCheck() {
	checkMap()
}

func checkMap() {
	tileSize := Vec{20, 20, 20}
	for id, act := range active {
		if !act || size[id].Equals(&Vec{}) {
			continue
		}
		px := math.Floor(pos[id][0] / 20)
		py := math.Floor(pos[id][1] / 20)

		r := size[id][0] / 2

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
				v.Sub(&pos[id], &tilePos)
				v.Clamp(&tileSize)

				d := Vec{}
				d.Sub(&pos[id], &tilePos)
				d.Sub(&d, &v)

				dist := math.Sqrt(d.Nrm2Sq())
				vProj := Dot(&vel[id], &d)
				vProj /= dist

				remove := dist - r + vProj
				if remove < 0 {
					if dist < r {
						dPos := Vec{}
						dPos.Scale(r/dist-1, &d)
						pos[id].Add(&pos[id], &dPos)
					}

					d.Scale(-remove/dist, &d)
					vel[id].Add(&vel[id], &d)
				}
			}
		}
	}
}

func placeTurrets() {
	for id := range players {
		if !players[id].active {
			continue
		}
		entId := players[id].ent
		turretId := child[entId]

		offset := Vec{}
		offset.Scale(4, &players[id].target)
		pos[turretId].Add(&pos[entId], &offset)
		rot[turretId] = math.Atan2(offset[1], offset[0])
	}
}
