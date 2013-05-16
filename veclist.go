package main

type VecList []Vec

func NewVecList(n int) []Vec {
	return make(VecList, n)
}
