package synth

import "github.com/hajimehoshi/oto/v2"

type Configuration struct {
	SampleRate      int
	ChannelCount    int
	BitDepthInBytes int
}

var (
	g = Configuration{
		SampleRate:      48000,
		ChannelCount:    2,
		BitDepthInBytes: 2,
	}
)

func NewContext(sampleRate int, channelCount int, bitDepthInBytes int) (*oto.Context, chan struct{}, error) {
	g = Configuration{
		SampleRate:      sampleRate,
		ChannelCount:    channelCount,
		BitDepthInBytes: bitDepthInBytes,
	}
	return oto.NewContext(sampleRate, channelCount, bitDepthInBytes)
}
