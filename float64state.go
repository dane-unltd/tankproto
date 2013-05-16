package main

type Float64State struct {
	ids IdList
	v   []float64
	id  CompId
}

func NewFloat64State(id CompId) *Float64State {
	fs := &Float64State{}
	fs.ids = make(IdList, 0, 10)
	fs.v = make([]float64, 0, 10)
	fs.id = id
	return fs
}

func (dst *Float64State) Copy(src *Float64State) {
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

func (fs *Float64State) Append(id EntId, v float64) int {
	fs.ids = append(fs.ids, id)
	fs.v = append(fs.v, v)

	return len(fs.ids) - 1
}

func (fs *Float64State) Remove(i int) {
	fs.ids[i] = fs.ids[len(fs.ids)-1]
	fs.v[i] = fs.v[len(fs.v)-1]

	fs.ids = fs.ids[:len(fs.ids)-1]
	fs.v = fs.v[:len(fs.v)-1]

	ents[fs.ids[i]][fs.id] = i
}
