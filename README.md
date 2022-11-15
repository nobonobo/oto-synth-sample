# oto-synth-sample

Synthesizer sample

# sample code

```go
import "github.com/nobonobo/oto-synth-sample/synth"
```

```go
ctx, ready, err := synth.NewContext(48000, 1, 2)
if err != nil {
	panic(err)
}
<-ready
log.Println("oto ready")
```

```go
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
```

# harmony tone

```go
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
```

# run

```shell
$ go mod tidy
$ go run .
2022/11/15 20:19:18 oto ready
```
