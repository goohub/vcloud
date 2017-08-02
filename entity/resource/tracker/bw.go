package tracker

import (
	"log"
)

type BwTracker interface {
	Claim(bw float64) bool
	Allocate(id int, bw float64)
	Deallocate(id int, bw float64)
}

type BwProvisioner struct {
	bw               float64
	allocabledBw     float64
	allocatedBwTable map[int]float64
}

func (bp *BwProvisioner) Claim(bw float64) bool {
	return bp.allocabledBw+bw <= bp.bw
}

func (bp *BwProvisioner) Allocate(id int, bw float64) {
	if bp.allocatedBwTable == nil{
		bp.allocatedBwTable = make(map[int]float64)
	}
	bp.allocatedBwTable[id] = bw
	bp.allocabledBw += bw
}

func (bp *BwProvisioner) Deallocate(id int, bw float64) {
	if _, ok := bp.allocatedBwTable[id]; !ok {
		log.Panicf("deallocate failed")
	}
	delete(bp.allocatedBwTable, id)
	bp.allocabledBw -= bw
}

func (bp *BwProvisioner)GetBw()float64{
	return bp.bw
}

func (bp *BwProvisioner)SetBw(bw float64){
	bp.bw = bw
}
