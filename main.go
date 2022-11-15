package main

import (
	"log"
	"time"

	"otosample/synth"
)

func main() {
	ctx, ready, err := synth.NewContext(48000, 1, 2)
	if err != nil {
		panic(err)
	}
	<-ready
	log.Println("oto ready")
	sig := synth.NewSig(0)
	output := &synth.Output{Source: synth.Gain(0.3, sig)}
	player := ctx.NewPlayer(output)
	player.Play()
	go func() {
		time.Sleep(time.Second)
		sig.SetFreq(440)
		time.Sleep(time.Second)
		sig.SetFreq(880)
		time.Sleep(time.Second)
		sig.SetFreq(440)
		time.Sleep(time.Second)
		sig.SetFreq(880)
		time.Sleep(time.Second)
		sig.SetFreq(0)
		output.Stop()
	}()
	for player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}
}
