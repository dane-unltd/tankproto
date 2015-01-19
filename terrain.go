package main

import (
	"bytes"
	"encoding/binary"
	"github.com/dane-unltd/engine/bitmask"
	"io"
	"math"
)

type Terrain struct {
	N       int
	Dx      float64
	Heights []float64
}

func NewTerrain(N int, Dx float64) *Terrain {
	return &Terrain{
		N:       N,
		Dx:      Dx,
		Heights: make([]float64, N*(N+1)/2),
	}
}

func (t *Terrain) Set(x, y int, h float64) {
	t.Heights[x+t.N*y-y*(y-1)/2] = h
}

func (t *Terrain) At(x, y int) float64 {
	if x < 0 || y < 0 || y >= t.N || x >= t.N-y {
		panic("terrain: index out of range")
	}
	return t.Heights[x+t.N*y-y*(y-1)/2]
}

func (dest *Terrain) Copy(src *Terrain) {
	dest.N = src.N
	if len(dest.Heights) < src.N*(src.N+1)/2 {
		dest.Heights = make([]float64, src.N*(src.N+1)/2)
	}
	copy(dest.Heights, src.Heights)
}

func (t *Terrain) Serialize(buf io.Writer, serAll bool, tOld *Terrain) {
	bitMask := bitmask.New(len(t.Heights))
	bufTemp := &bytes.Buffer{}
	for i := range t.Heights {
		if serAll || t.Heights[i] != tOld.Heights[i] {
			bitMask.Set(i)
			binary.Write(bufTemp, binary.LittleEndian, t.Heights[i])
		}
	}
	buf.Write(bitMask)
	buf.Write(bufTemp.Bytes())
}

func (t *Terrain) Transform(x, y float64) (xt, yt float64) {
	xt = (x - y/math.Sqrt(3)) / t.Dx
	yt = (2 * y / math.Sqrt(3)) / t.Dx
	return
}
