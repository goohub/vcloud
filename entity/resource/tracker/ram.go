package tracker

import (
	"log"
)

type RamTracker interface {
	Claim(ram float64) bool
	Allocate(id int, ram float64)
	Deallocate(id int, ram float64)
}

type RamProvisioner struct {
	ram               float64
	allocabledRam     float64
	allocatedRamTable map[int]float64
}

func (rp *RamProvisioner) Claim(ram float64) bool {
	return (rp.ram - rp.allocabledRam) >= ram
}

func (rp *RamProvisioner) Allocate(id int, ram float64) {
	if rp.allocatedRamTable == nil{
		rp.allocatedRamTable = make(map[int]float64)
	}
	rp.allocatedRamTable[id] = ram
	rp.allocabledRam += ram
}

func (rp *RamProvisioner) Deallocate(id int, ram float64) {
	if _, ok := rp.allocatedRamTable[id]; !ok {
		log.Panicf("deallocate failed")
	}
	delete(rp.allocatedRamTable, id)
	rp.allocabledRam -= ram
}

func (rp *RamProvisioner)GetRam()float64{
	return rp.ram
}

func (rp *RamProvisioner)SetRam(ram float64){
	rp.ram = ram
}
