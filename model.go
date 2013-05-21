package main

type ModelState []ModelId

func NewModelState() ModelState {
	return make(ModelState, 0, 10)
}

func (ms *ModelState) Clone() StateComp {
	res := make(ModelState, len(*ms))
	copy(res, *ms)
	return &res
}

func (dst *ModelState) Copy(src Copyer) {
	s := src.(*ModelState)
	if len(*dst) < len(*s) {
		*dst = make(ModelState, len(*s))
	}
	copy(*dst, *s)
}

func (ms ModelState) Val(id EntId) interface{} {
	return ms[id]
}

func (ms ModelState) Equal(v interface{}, id EntId) bool {
	model := v.(ModelId)
	return model == ms[id]
}

func (ms *ModelState) Append() {
	(*ms) = append((*ms), 0)
}
