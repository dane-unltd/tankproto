package main

type VecState []Vec

func NewVecState() VecState {
	return make(VecState, 0, 10)
}

func (dst *VecState) Copy(src Copyer) {
	s := src.(*VecState)
	if len(*dst) < len(*s) {
		*dst = make(VecState, len(*s))
	}
	copy(*dst, *s)
}

func (vs *VecState) Clone() StateComp {
	res := make(VecState, len(*vs))
	copy(res, *vs)
	return &res
}

func (vs VecState) Val(id EntId) interface{} {
	return &vs[id]
}

func (vs VecState) Equal(v interface{}, id EntId) bool {
	vec := v.(*Vec)
	return vec.Equals(&vs[id])
}

func (vs *VecState) Append() {
	(*vs) = append((*vs), Vec{})
}
