package main

import (
	"log"
	"time"

	"github.com/nobonobo/oto-synth-sample/synth"
)

func main() {
	ctx, ready, err := synth.NewContext(48000, 1, 2)
	if err != nil {
		panic(err)
	}
	<-ready
	log.Println("oto ready")
	sig1 := synth.NewSig(440)
	sig2 := synth.NewSig(493.8833012561241)
	sig3 := synth.NewSig(523.2511306011972)
	multi := synth.Multiplex(
		synth.Gain(0.3, sig1),
		synth.Gain(0.3, sig2),
		synth.Gain(0.3, sig3),
	)
	output := &synth.Output{Source: multi}
	player := ctx.NewPlayer(output)
	player.Play()
	go func() {
		time.Sleep(time.Second)
		output.Stop()
	}()
	for player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}
}
