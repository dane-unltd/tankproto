package main

type BoolState []bool

func NewBoolState() BoolState {
	return make(BoolState, 0, 10)
}

func (dst *BoolState) Copy(src Copyer) {
	s := src.(*BoolState)
	if len(*dst) < len(*s) {
		*dst = make(BoolState, len(*s))
	}
	copy(*dst, *s)
}

func (bs *BoolState) Clone() StateComp {
	res := make(BoolState, len(*bs))
	copy(res, *bs)
	return &res
}

func (bs BoolState) Val(id EntId) interface{} {
	return bs[id]
}

func (bs BoolState) Equal(v interface{}, id EntId) bool {
	b := v.(bool)
	return b == bs[id]
}

func (bs *BoolState) Append() {
	(*bs) = append((*bs), false)
}
