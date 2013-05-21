package main

type EntId uint32
type PlayerId uint32

type IdList []EntId
type IdMap map[EntId]struct{}

type IdGen struct {
	maxId    uint32
	freeIds  []uint32
	idOut    chan uint32
	idIn     chan uint32
	maxIdInc func()
}

func NewIdGen(idInc func()) *IdGen {
	g := IdGen{}
	g.idOut = make(chan uint32)
	g.idIn = make(chan uint32, 8)
	g.maxIdInc = idInc
	return &g
}

func (g *IdGen) NewId() uint32 {
	return <-g.idOut
}

func (g *IdGen) FreeId(id uint32) {
	g.idIn <- id
}

func (g *IdGen) Run() {
	g.maxIdInc()
	g.maxIdInc()
	currId := g.maxId + 1
	for {
		select {
		case id := <-g.idIn:
			g.freeIds = append(g.freeIds, id)
			currId = id
		case g.idOut <- currId:
			if currId > g.maxId {
				g.maxId = currId
				g.maxIdInc()
				currId++
			} else {
				g.freeIds = g.freeIds[:len(g.freeIds)-1]
				if len(g.freeIds) > 0 {
					currId = g.freeIds[len(g.freeIds)-1]
				} else {
					currId = g.maxId + 1
				}
			}
		}
	}
}

var plGen, entGen = NewIdGen(incMaxPlayers), NewIdGen(incMaxEnts)

func init() {
	go plGen.Run()
	go entGen.Run()
}

func newEntId() EntId {
	return EntId(entGen.NewId())
}

func freePlayerId(id PlayerId) {
	plGen.FreeId(uint32(id))
}

func newPlayerId() PlayerId {
	return PlayerId(plGen.NewId())
}

func freeEntId(id EntId) {
	entGen.FreeId(uint32(id))
}
