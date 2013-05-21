package main

type ModelState []ModelId

func NewModelState() ModelState {
	return make([]ModelId, 0, 10)
}

func (dst ModelState) Copy(src Copyer) {
	s := src.(ModelState)
	copy(dst, s)
}

func (ms *ModelState) Append() {
	(*ms) = append((*ms), 0)
}
