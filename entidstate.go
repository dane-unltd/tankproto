package main

type EntIdState []EntId

func NewEntIdState() EntIdState {
	return make(EntIdState, 0, 10)
}

func (es *EntIdState) Clone() StateComp {
	res := make(EntIdState, len(*es))
	copy(res, *es)
	return &res
}

func (dst *EntIdState) Copy(src Copyer) {
	s := src.(*EntIdState)
	if len(*dst) < len(*s) {
		*dst = make(EntIdState, len(*s))
	}
	copy(*dst, *s)
}

func (es EntIdState) Val(id EntId) interface{} {
	return es[id]
}

func (es EntIdState) Equal(v interface{}, id EntId) bool {
	entId := v.(EntId)
	return entId == es[id]
}

func (es *EntIdState) Append() {
	(*es) = append((*es), 0)
}
