package main

type PlayerState struct {
	score  []uint32
	target []Vec
	turret []EntId
	ids    IdList
}

func NewPlayerState() *PlayerState {
	p := &PlayerState{}
	p.score = make([]uint32, 0, 10)
	p.turret = make([]EntId, 0, 10)
	p.target = make([]Vec, 0, 10)
	p.ids = make(IdList, 0, 10)
	return p
}

func (dst *PlayerState) Copy(src *PlayerState) {
	copy(dst.ids, src.ids)
	copy(dst.target, src.target)
	copy(dst.score, src.score)
	copy(dst.turret, src.turret)

	if len(dst.ids) < len(src.ids) {
		dst.ids = append(dst.ids, src.ids[len(dst.ids):]...)
		dst.target = append(dst.target, src.target[len(dst.ids):]...)
		dst.score = append(dst.score, src.score[len(dst.ids):]...)
		dst.turret = append(dst.turret, src.turret[len(dst.ids):]...)

	} else {
		dst.ids = dst.ids[:len(src.ids)]
		dst.target = dst.target[:len(src.ids)]
		dst.score = dst.score[:len(src.ids)]
		dst.turret = dst.turret[:len(src.ids)]
	}
}

func (ps *PlayerState) Append(id EntId, target Vec, score uint32, turret EntId) int {
	ps.ids = append(ps.ids, id)
	ps.target = append(ps.target, target)
	ps.score = append(ps.score, score)
	ps.turret = append(ps.turret, turret)

	return len(ps.ids) - 1
}

func (ps *PlayerState) Remove(i int) {
	ps.ids[i] = ps.ids[len(ps.ids)-1]
	ps.target[i] = ps.target[len(ps.target)-1]
	ps.score[i] = ps.score[len(ps.score)-1]
	ps.turret[i] = ps.turret[len(ps.turret)-1]

	ps.ids = ps.ids[:len(ps.ids)-1]
	ps.target = ps.target[:len(ps.target)-1]
	ps.score = ps.score[:len(ps.score)-1]
	ps.turret = ps.turret[:len(ps.turret)-1]

	ents[ps.ids[i]][PlayerComp] = i
}
