package main

import (
	"fmt"
	"github.com/dane-unltd/engine/math3"
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
	checkTerrain(terrain)
	checkEnts()
}

func checkTerrain(t *Terrain) {
	for id, ct := range collType {
		if ct == CollNone {
			continue
		}
		px, py := t.Transform(pos[id][0], pos[id][1])

		x, y := math.Ceil(px), math.Ceil(py)

		fmt.Println(pos[id][0], pos[id][1], px, py)
		pos[id][2] = t.At(int(x), int(y))
	}
}

func checkMap(tileMap *TileMap) {
	tileSize := math3.Vec{20, 20, 20}
	for id, ct := range collType {
		if ct == CollNone {
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
				tilePos := math3.Vec{tx*20 + 10, ty*20 + 10, 0}

				v := math3.Vec{}
				v.Sub(&pos[id], &tilePos)
				v.Clamp(&tileSize)

				d := math3.Vec{}
				d.Sub(&pos[id], &tilePos)
				d.Sub(&d, &v)

				dist := math.Sqrt(d.Nrm2Sq())
				vProj := math3.Dot(&vel[id], &d)
				vProj /= dist

				remove := dist - r + vProj
				if remove < 0 {
					if dist < r {
						dPos := math3.Vec{}
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

func checkEnts() {
	for id1, ct1 := range collType {
		if ct1 == CollNone {
			continue
		}
		for id2 := id1 + 1; id2 < len(collType); id2++ {
			ct2 := collType[id2]
			if ct2 == CollNone {
				continue
			}
			if ct1 == CollProjectile && ct2 == CollProjectile {
				continue
			}
			r1, r2 := size[id1][0]/2, size[id2][0]/2
			d := math3.Vec{}

			d.Sub(&pos[id1], &pos[id2])
			dist := d.Nrm2()
			vDiff := math3.Vec{}
			vDiff.Sub(&vel[id1], &vel[id2])
			vProj := math3.Dot(&vDiff, &d)
			vProj /= dist

			remove := dist - r1 - r2 + vProj
			if remove < 0 {
				if dist < r1+r2 {
					dPos := math3.Vec{}
					dPos.Scale(0.5*((r1+r2)/dist-1), &d)
					pos[id1].Add(&pos[id1], &dPos)
					pos[id2].Sub(&pos[id2], &dPos)
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

		offset := math3.Vec{}
		offset.Scale(4, &players[id].target)
		pos[turretId].Add(&pos[entId], &offset)
		rot[turretId] = math.Atan2(offset[1], offset[0])
	}
}
