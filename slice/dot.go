package slice

func Dot(x, y []float32) float32 {
	var d float32
	for i := range x {
		d += x[i] * y[i]
	}
	return d
}
