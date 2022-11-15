package synth

import "io"

type Output struct {
	Source  Source
	stopped bool
}

func (out *Output) Stop() {
	out.stopped = true
}

func (out *Output) Read(buf []byte) (int, error) {
	num := (g.BitDepthInBytes) * (g.ChannelCount)
	switch g.BitDepthInBytes {
	case 1:
		for i := 0; i < len(buf)/num; i++ {
			b := int(out.Source.Next() * 127)
			for ch := 0; ch < g.ChannelCount; ch++ {
				buf[num*i+ch] = byte(b + 128)
			}
		}
	case 2:
		for i := 0; i < len(buf)/num; i++ {
			b := int16(out.Source.Next() * 32767)
			for ch := 0; ch < g.ChannelCount; ch++ {
				buf[num*i+2*ch] = byte(b)
				buf[num*i+1+2*ch] = byte(b >> 8)
			}
		}
	}
	n := len(buf)
	if out.stopped {
		return n, io.EOF
	}
	return len(buf), nil
}
