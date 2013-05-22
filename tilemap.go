package main

import (
	"bytes"
	"encoding/binary"
	"github.com/dane-unltd/engine/bitmask"
	"io"
)

type TileMap struct {
	rows, cols int
	Tiles      []byte
}

func NewTileMap(rows, cols int) *TileMap {
	tm := &TileMap{}
	tm.rows = rows
	tm.cols = cols

	tm.Tiles = make([]byte, rows*cols)

	return tm
}

func (tm *TileMap) Set(x, y int, t byte) {
	tm.Tiles[x+tm.rows*y] = t
}
func (tm *TileMap) At(x, y int) byte {
	return tm.Tiles[x+tm.rows*y]
}

func (dest *TileMap) Copy(src *TileMap) {
	dest.rows = src.rows
	dest.cols = src.cols

	if len(dest.Tiles) < len(src.Tiles) {
		dest.Tiles = make([]byte, dest.rows*dest.cols)
	}
	dest.Tiles = dest.Tiles[:dest.rows*dest.cols]
	copy(dest.Tiles, src.Tiles)
}

func (tm *TileMap) Serialize(buf io.Writer, serAll bool, tmOld *TileMap) {
	bitMask := bitmask.New(len(tm.Tiles))
	bufTemp := &bytes.Buffer{}
	for i := range tm.Tiles {
		if serAll || tm.Tiles[i] != tmOld.Tiles[i] {
			bitMask.Set(i)
			binary.Write(bufTemp, binary.LittleEndian, tm.Tiles[i])
		}
	}
	buf.Write(bitMask)
	buf.Write(bufTemp.Bytes())
}
