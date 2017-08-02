package tracker

import (
	"log"
)

type MipsTracker interface {
	GetMips() float64
	Claim(mips float64) bool
	Allocate(id int, mips float64)
	Deallocate(id int, mips float64) bool
}

func NewMipsTracker(
	mips float64,
	allocabledMips float64,
	allocatedMipsTable map[int]float64) MipsTracker {
	return &mipsTracker{
		mips,
		allocabledMips,
		allocatedMipsTable,
	}
}

type mipsTracker struct {
	mips               float64
	allocabledMips     float64
	allocatedMipsTable map[int]float64
}

func (mp *mipsTracker) Claim(mips float64) bool {
	return (mp.mips - mp.allocabledMips) >= mips
}

func (mp *mipsTracker) Allocate(id int, mips float64) {
	if mp.allocatedMipsTable == nil {
		mp.allocatedMipsTable = make(map[int]float64)
	}
	mp.allocatedMipsTable[id] = mips
	mp.allocabledMips += mips
}

func (mp *mipsTracker) Deallocate(id int, mips float64) bool {
	if _, ok := mp.allocatedMipsTable[id]; !ok {
		log.Printf("deallocate failed")
		return false
	}
	delete(mp.allocatedMipsTable, id)
	mp.allocabledMips -= mips
	return true
}

func (mp *mipsTracker) GetMips() float64 {
	return mp.mips
}
