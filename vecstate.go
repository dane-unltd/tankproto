package main

type VecState []Vec

func NewVecState() VecState {
	return make([]Vec, 0, 10)
}

func (dst VecState) Copy(src Copyer) {
	s := src.(VecState)
	copy(dst, s)
}

func (vs *VecState) Append() {
	(*vs) = append((*vs), Vec{})
}
