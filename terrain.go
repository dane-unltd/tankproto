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
		Heights: make([]float64, N*N),
	}
}

func (t *Terrain) Set(x, y int, h float64) {
	t.Heights[x+t.N*y] = h
}

func (t *Terrain) At(x, y int) float64 {
	return t.Heights[x+t.N*y]
}

func (dest *Terrain) Copy(src *Terrain) {
	dest.N = src.N
	if len(dest.Heights) < src.N*src.N {
		dest.Heights = make([]float64, src.N*src.N)
	}
	copy(dest.Heights, src.Heights[:src.N*src.N])
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
