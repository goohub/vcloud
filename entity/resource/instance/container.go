package instance

type Container struct {
	id   int
	mips float64
	ram  float64
	bw   float64
}

func (container *Container) SetQuota(mips, ram, bw float64) {
	container.mips = mips
	container.ram = ram
	container.bw = bw
}

func (container *Container) GetBw() float64 {
	return container.bw
}

func (container *Container) GetRam() float64 {
	return container.ram
}

func (container *Container) GetMips() float64 {
	return container.mips
}

func(container *Container) SetId(id int){
	container.id = id
}

func(container *Container)GetId()int{
	return container.id
}
