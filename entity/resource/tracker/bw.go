package tracker

import (
	"log"
)

type BwTracker interface {
	GetBw() float64
	Claim(bw float64) bool
	Allocate(id int, bw float64)
	Deallocate(id int, bw float64)
}

func NewBwTracker(
	bw float64,
	allocabledBw float64,
	allocatedBwTable map[int]float64) BwTracker {
	return &bwTracker{
		bw,
		allocabledBw,
		allocatedBwTable,
	}
}

type bwTracker struct {
	bw               float64
	allocabledBw     float64
	allocatedBwTable map[int]float64
}

func (bp *bwTracker) Claim(bw float64) bool {
	return bp.allocabledBw+bw <= bp.bw
}

func (bp *bwTracker) Allocate(id int, bw float64) {
	if bp.allocatedBwTable == nil {
		bp.allocatedBwTable = make(map[int]float64)
	}
	bp.allocatedBwTable[id] = bw
	bp.allocabledBw += bw
}

func (bp *bwTracker) Deallocate(id int, bw float64) {
	if _, ok := bp.allocatedBwTable[id]; !ok {
		log.Panicf("deallocate failed")
	}
	delete(bp.allocatedBwTable, id)
	bp.allocabledBw -= bw
}

func (bp *bwTracker) GetBw() float64 {
	return bp.bw
}
