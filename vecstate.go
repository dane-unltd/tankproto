package main

type VecState struct {
	ids IdList
	v   []Vec
	id  CompId
}

func NewVecState(id CompId) *VecState {
	vs := &VecState{}
	vs.ids = make(IdList, 0, 10)
	vs.v = make([]Vec, 0, 10)
	vs.id = id
	return vs
}

func (dst *VecState) Copy(src *VecState) {
	copy(dst.ids, src.ids)
	copy(dst.v, src.v)

	if len(dst.ids) < len(src.ids) {
		dst.ids = append(dst.ids, src.ids[len(dst.ids):]...)
		dst.v = append(dst.v, src.v[len(dst.ids):]...)
	} else {
		dst.ids = dst.ids[:len(src.ids)]
		dst.v = dst.v[:len(src.ids)]
	}
}

func (vs *VecState) Append(id EntId, v Vec) int {
	vs.ids = append(vs.ids, id)
	vs.v = append(vs.v, v)
	return len(vs.ids) - 1
}

func (vs *VecState) Remove(i int) {
	vs.ids[i] = vs.ids[len(vs.ids)-1]
	vs.v[i] = vs.v[len(vs.v)-1]

	vs.ids = vs.ids[:len(vs.ids)-1]
	vs.v = vs.v[:len(vs.v)-1]

	ents[vs.ids[i]][vs.id] = i
}
