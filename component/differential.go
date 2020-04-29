package component

type Differential struct {
	L, R Motor
}

func (d *Differential) Setup() {
	d.L.Setup()
	d.R.Setup()
}

func (d *Differential) On() {
	d.L.On()
	d.R.On()
}

func (d *Differential) Off() {
	d.L.Off()
	d.R.Off()
}

func (d *Differential) Forward() {
	d.L.Forward()
	d.R.Forward()
}

func (d *Differential) Backward() {
	d.L.Backward()
	d.R.Backward()
}

func (d *Differential) Left() {
	d.L.Forward()
	d.R.Backward()
}

func (d *Differential) Right() {
	d.L.Backward()
	d.R.Forward()
}
