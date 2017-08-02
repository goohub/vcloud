package tracker

import (
	"log"
)

type MipsTracker interface {
	Claim(mips float64) bool
	Allocate(id int64, mips float64)
	Deallocate(id int64, mips float64) bool
}

type MipsProvisioner struct {
	mips               float64
	allocabledMips     float64
	allocatedMipsTable map[int]float64
}

func (mp *MipsProvisioner) Claim(mips float64) bool {
	return (mp.mips - mp.allocabledMips) >= mips
}

func (mp *MipsProvisioner) Allocate(id int, mips float64){
	if mp.allocatedMipsTable == nil{
		mp.allocatedMipsTable = make(map[int]float64)
	}
	mp.allocatedMipsTable[id] = mips
	mp.allocabledMips += mips
}

func (mp *MipsProvisioner) Deallocate(id int, mips float64) bool {
	if _, ok := mp.allocatedMipsTable[id]; !ok {
		log.Printf("deallocate failed")
		return false
	}
	delete(mp.allocatedMipsTable, id)
	mp.allocabledMips -= mips
	return true
}

func (mp *MipsProvisioner) GetMips() float64{
	return mp.mips
}

func (mp *MipsProvisioner) SetMips(mips float64) {
	mp.mips = mips
}
