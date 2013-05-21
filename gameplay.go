package main

import (
	"fmt"
)

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
