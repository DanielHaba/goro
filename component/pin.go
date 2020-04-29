package component

import (
	"math"
	"sync"
	"time"

	"github.com/stianeikeland/go-rpio"
)

type InputPin interface {
	Input()
	Read() rpio.State
}

type OutputPin interface {
	Output()
	Write(state rpio.State)
	High()
	Low()
	Toggle()
}

type PwmPin interface {
	Pwm()
	Freq(freq int)
	DutyCycle(dutyLen, cycleLen uint32)
}

type SoftwarePwm struct {
	P      OutputPin
	m      sync.Mutex
	f      int
	r      bool
	dl, cl uint32
}

type pwmable interface {
	Pwm()
}

func (p *SoftwarePwm) Input() {
	p.closePwm()
}

func (p *SoftwarePwm) Output() {
	p.closePwm()
}

func (p *SoftwarePwm) Pwm() {
	p.m.Lock()
	defer p.m.Unlock()
	if !p.r {
		p.P.Output()
		p.r = true
		go p.pwmThread()
	}
}

func (p *SoftwarePwm) Write(state rpio.State) {
	p.P.Write(state)
}

func (p *SoftwarePwm) High() {
	p.P.High()
}

func (p *SoftwarePwm) Low() {
	p.P.Low()
}

func (p *SoftwarePwm) Toggle() {
	p.P.Toggle()
}

func (p *SoftwarePwm) Freq(freq int) {
	p.m.Lock()
	defer p.m.Unlock()
	p.f = freq
}

func (p *SoftwarePwm) DutyCycle(dutyLen, cycleLen uint32) {
	p.m.Lock()
	defer p.m.Unlock()
	p.dl = dutyLen
	p.cl = cycleLen
}

// func (p *SoftwarePwm) pwmThread() {
// 	p.m.Lock()
// 	r := p.r
// 	p.m.Unlock()
// 	di := uint32(0)
// 	s := rpio.Low

// 	for r {
// 		p.m.Lock()
// 		r = p.r
// 		// todo: support negative frequency
// 		f := int(math.Abs(float64(p.f)))
// 		dl := p.dl
// 		cl := p.cl
// 		p.m.Unlock()

// 		cs := rpio.High
// 		if di > dl {
// 			cs = rpio.Low
// 		}
// 		if s != cs {
// 			p.Write(cs)
// 		}
// 		s = cs
// 		if f == 0 {
// 			time.Sleep(time.Millisecond)
// 		} else {
// 			time.Sleep(time.Duration(int(time.Second) / f))
// 		}
// 		di = (di + 1) % cl
// 	}
// }

func (p *SoftwarePwm) pwmThread() {
	p.m.Lock()
	r := p.r
	p.m.Unlock()

	for r {
		p.m.Lock()
		r = p.r
		// todo: support negative frequency
		f := int(math.Abs(float64(p.f)))
		dl := p.dl
		cl := p.cl
		p.m.Unlock()
		if cl != 0 && f != 0 {
			cd := float64(int(time.Second)/f) / float64(cl)
			hd := time.Duration(float64(dl) * cd)
			ld := time.Duration(float64(cl-dl) * cd)
	
			p.P.High()
			time.Sleep(hd)
			p.P.Low()
			time.Sleep(ld)
		}
	}
}

func (p *SoftwarePwm) closePwm() {
	p.m.Lock()
	defer p.m.Unlock()
	p.r = false
}
