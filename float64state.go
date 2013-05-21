package main

type Float64State []float64

func NewFloat64State() Float64State {
	return make(Float64State, 0, 10)
}

func (fs *Float64State) Clone() StateComp {
	res := make(Float64State, len(*fs))
	copy(res, *fs)
	return &res
}

func (dst *Float64State) Copy(src Copyer) {
	s := src.(*Float64State)
	if len(*dst) < len(*s) {
		*dst = make(Float64State, len(*s))
	}
	copy(*dst, *s)
}

func (fs Float64State) Val(id EntId) interface{} {
	return fs[id]
}

func (fs Float64State) Equal(v interface{}, id EntId) bool {
	f64 := v.(float64)
	return f64 == fs[id]
}

func (fs *Float64State) Append() {
	(*fs) = append((*fs), 0)
}
