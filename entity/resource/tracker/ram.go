package tracker

import (
	"log"
)

type RamTracker interface {
	GetRam() float64
	Claim(ram float64) bool
	Allocate(id int, ram float64)
	Deallocate(id int, ram float64)
}

func NewRamTracker(
	ram float64,
	allocabledRam float64,
	allocatedRamTable map[int]float64) RamTracker {
	return &ramTracker{
		ram,
		allocabledRam,
		allocatedRamTable,
	}
}

type ramTracker struct {
	ram               float64
	allocabledRam     float64
	allocatedRamTable map[int]float64
}

func (rp *ramTracker) Claim(ram float64) bool {
	return (rp.ram - rp.allocabledRam) >= ram
}

func (rp *ramTracker) Allocate(id int, ram float64) {
	if rp.allocatedRamTable == nil {
		rp.allocatedRamTable = make(map[int]float64)
	}
	rp.allocatedRamTable[id] = ram
	rp.allocabledRam += ram
}

func (rp *ramTracker) Deallocate(id int, ram float64) {
	if _, ok := rp.allocatedRamTable[id]; !ok {
		log.Panicf("deallocate failed")
	}
	delete(rp.allocatedRamTable, id)
	rp.allocabledRam -= ram
}

func (rp *ramTracker) GetRam() float64 {
	return rp.ram
}
