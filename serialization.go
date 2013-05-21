package main

import (
	"bytes"
	"encoding/binary"
	"io"
)

func serialize(buf io.Writer, serAll bool) {
	nEnts := 0
	for _, act := range active {
		if act {
			nEnts++
		}
	}
	binary.Write(buf, binary.LittleEndian, uint32(nEnts))
	for id, act := range active {
		if !act {
			continue
		}
		binary.Write(buf, binary.LittleEndian, EntId(id))

		bitMask := NewBitMask(4)
		bufTemp := &bytes.Buffer{}
		for i, compId := range networkedComps {
			v := stateComps[compId].Val(EntId(id))
			if serAll || !stateComps[compId].Equal(v, EntId(id)) {
				bitMask.Set(i)
				binary.Write(bufTemp, binary.LittleEndian, v)
			}

		}
		buf.Write(bitMask)
		buf.Write(bufTemp.Bytes())
	}

	tileMap.Serialize(buf, serAll, tileMapOld)
}
