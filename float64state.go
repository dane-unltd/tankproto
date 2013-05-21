package main

type Float64State []float64

func NewFloat64State() Float64State {
	return make([]float64, 0, 10)
}

func (dst Float64State) Copy(src Copyer) {
	s := src.(Float64State)
	copy(dst, s)
}

func (fs *Float64State) Append() {
	(*fs) = append((*fs), 0)
}
