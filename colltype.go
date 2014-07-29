package main

import (
	"github.com/dane-unltd/engine/entstate"
)

type CollTypeState []CollType

func NewCollTypeState() CollTypeState {
	return make(CollTypeState, 0, 10)
}

func (st *CollTypeState) Clone() interface{} {
	res := make(CollTypeState, len(*st))
	copy(res, *st)
	return &res
}

func (st CollTypeState) Zero(id entstate.EntId) {
	st[id] = 0
}

func (dst *CollTypeState) Copy(src interface{}) {
	s := src.(*CollTypeState)
	if len(*dst) < len(*s) {
		*dst = make(CollTypeState, len(*s))
	}
	copy(*dst, *s)
}

func (st CollTypeState) Val(id entstate.EntId) interface{} {
	return st[id]
}

func (st CollTypeState) Equal(v interface{}, id entstate.EntId) bool {
	ct := v.(CollType)
	return ct == st[id]
}

func (st *CollTypeState) Append(n uint32) {
	for len(*st) <= int(n) {
		(*st) = append((*st), 0)
	}
}
