package synth

type SourceFunc func() float32

func (sf SourceFunc) Next() float32 {
	return sf()
}

type Source interface {
	Next() float32
}

func Gain(gain float32, src Source) Source {
	return SourceFunc(func() float32 {
		return gain * src.Next()
	})
}

func Multiplex(srcs ...Source) Source {
	return SourceFunc(func() float32 {
		sum := float32(0.0)
		for _, v := range srcs {
			sum += v.Next()
		}
		return sum
	})
}
