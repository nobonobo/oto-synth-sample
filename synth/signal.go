package synth

import (
	"math"

	"golang.org/x/mobile/exp/f32"
)

type Signal struct {
	tick int64
	k    float32
}

func (sg *Signal) SetFreq(f float32) {
	sg.k = 2 * math.Pi * f / float32(g.SampleRate)
}

func (sg *Signal) Next() float32 {
	sg.tick++
	return f32.Sin(sg.k * float32(sg.tick))
}

func NewSig(freq float32) *Signal {
	sg := &Signal{}
	sg.SetFreq(freq)
	return sg
}
