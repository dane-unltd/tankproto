package main

import (
	"github.com/dane-unltd/engine/entstate"
)

type ModelState []ModelId

func NewModelState() ModelState {
	return make(ModelState, 0, 10)
}

func (st *ModelState) Clone() interface{} {
	res := make(ModelState, len(*st))
	copy(res, *st)
	return &res
}

func (st ModelState) Zero(id entstate.EntId) {
	st[id] = 0
}

func (dst *ModelState) Copy(src interface{}) {
	s := src.(*ModelState)
	if len(*dst) < len(*s) {
		*dst = make(ModelState, len(*s))
	}
	copy(*dst, *s)
}

func (st ModelState) Val(id entstate.EntId) interface{} {
	return st[id]
}

func (st ModelState) Equal(v interface{}, id entstate.EntId) bool {
	model := v.(ModelId)
	return model == st[id]
}

func (st *ModelState) Append(n uint32) {
	for len(*st) <= int(n) {
		(*st) = append((*st), 0)
	}
}
