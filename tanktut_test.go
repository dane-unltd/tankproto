package main

import (
	"math/rand"
	"testing"
)

func Benchmark_GameStateIndexes(b *testing.B) {
	b.StopTimer()

	ents := make([]map[int]int, 1000)
	for i := range ents {
		ents[i] = make(map[int]int)
	}

	vel := make([]float64, 1000)
	velIds := make([]int, 1000)

	pos := make([]float64, 1000)
	posIds := make([]int, 1000)

	for i := range pos {
		pos[i] = rand.Float64()
		posIds[i] = i
		vel[i] = rand.Float64()
		velIds[i] = i

		ents[i][0] = i
		ents[i][1] = i
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for i := range ents {
			posIx := ents[i][0]
			velIx := ents[i][1]
			pos[posIx] = pos[posIx] + vel[velIx]
		}
	}
}

func Benchmark_GameStateSlice(b *testing.B) {
	b.StopTimer()
	vel := make([]float64, 1000)
	pos := make([]float64, 1000)

	for i := range pos {
		pos[i] = rand.Float64()
		vel[i] = rand.Float64()
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for i := range pos {
			pos[i] = pos[i] + vel[i]
		}
	}
}

func Benchmark_GameStateMap(b *testing.B) {
	b.StopTimer()
	vel := make(map[uint32]float64, 1000)
	pos := make(map[uint32]float64, 1000)

	for i := uint32(0); i < 1000; i++ {
		pos[i] = rand.Float64()
		vel[i] = rand.Float64()
	}
	b.StartTimer()
	for i := uint32(0); int(i) < b.N; i++ {
		for i := range pos {
			pos[i] = pos[i] + vel[i]
		}
	}
}
