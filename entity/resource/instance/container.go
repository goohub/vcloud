package instance

type Container struct {
	id   int
	mips float64
	ram  float64
	bw   float64
}

func NewContainer(
	id int,
	mips float64,
	ram float64,
	bw float64) *Container {
	return &Container{
		id,
		mips,
		ram,
		bw,
	}
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

func(container *Container)GetId()int{
	return container.id
}
