package component

import "github.com/DanielHaba/goro/unit"

func lerp(i float64, a, b uint32) uint32 {
	return uint32(float64(b-a)*i) + a
}

type Servo struct {
	P PwmPin
	a unit.Angle
}

func (s *Servo) Setup() {
	s.P.Pwm()
	s.P.Freq(100)
}

func (s *Servo) Angle(a unit.Angle) {
	s.a = a
	d := lerp((a.Deg().Float64()+90)/180, 10, 20)
	s.P.DutyCycle(d, 100)
}

func (s *Servo) Left() {
	s.Angle(unit.Deg(-90))
}

func (s *Servo) Right() {
	s.Angle(unit.Deg(90))
}

func (s *Servo) Center() {
	s.Angle(unit.Deg(0))
}

func (s *Servo) Turn(a unit.Angle) {
	s.Angle(s.a.Add(a))
}
