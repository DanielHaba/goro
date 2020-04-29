package component

type Motor struct {
	A, B OutputPin
	E    PwmPin
}

func (m *Motor) Setup() {
	m.A.Output()
	m.B.Output()
	m.E.Pwm()
	m.E.Freq(1000)
	m.E.DutyCycle(0, 100)
}

func (m *Motor) On() {
	m.E.DutyCycle(100, 100)
}

func (m *Motor) Off() {
	m.E.DutyCycle(0, 100)
}

func (m *Motor) Speed(speed uint32) {
	m.E.DutyCycle(speed, 100)
}

func (m *Motor) Forward() {
	m.B.Low()
	m.A.High()
}

func (m *Motor) Backward() {
	m.A.Low()
	m.B.High()
}
