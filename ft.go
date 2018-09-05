package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func main() {
	fmt.Println(FT(1, 0, 0, 0, 0, 0, 0, 0))
}

func FT(x ...complex128) (X []complex128) {
	N := len(x)
	X = make([]complex128, N)
	for k := 0; k < N; k++ {
		for n := 0; n < N; n++ {
			X[k] += x[n] * cmplx.Exp(-cmplx.Sqrt(-1)*2*math.Pi*complex(float64(k*n)/float64(N), 0))
		}
		X[k] /= complex(float64(N), 0)
	}
	return X
}
