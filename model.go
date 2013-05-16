package main

type ModelState struct {
	ids IdList
	m   []ModelId
}

func NewModelState() *ModelState {
	m := &ModelState{}
	m.ids = make(IdList, 0, 10)
	m.m = make([]ModelId, 0, 10)
	return m
}

func (dst *ModelState) Copy(src *ModelState) {
	copy(dst.ids, src.ids)
	copy(dst.m, src.m)

	if len(dst.ids) < len(src.ids) {
		dst.ids = append(dst.ids, src.ids[len(dst.ids):]...)
		dst.m = append(dst.m, src.m[len(dst.ids):]...)
	} else {
		dst.ids = dst.ids[:len(src.ids)]
		dst.m = dst.m[:len(src.ids)]
	}
}

func (mst *ModelState) Append(id EntId, model ModelId) int {
	mst.ids = append(mst.ids, id)
	mst.m = append(mst.m, model)
	return len(mst.ids) - 1
}

func (mst *ModelState) Remove(i int) {
	mst.ids[i] = mst.ids[len(mst.ids)-1]
	mst.m[i] = mst.m[len(mst.m)-1]

	mst.ids = mst.ids[:len(mst.ids)-1]
	mst.m = mst.m[:len(mst.m)-1]

	ents[mst.ids[i]][ModelComp] = i
}
