package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func main() {
	fmt.Println(DFT(1, 0, 0, 0, 0, 0, 0, 0))
}

const (
	i2pi = (0+2i)*math.Pi
	ftPrec = 0.001
	iftPrec = 0.01
)

func DFT(x ...complex128) (X []complex128) {
	N := float64(len(x))
	Ninv := complex(1/N,0)
	X = make([]complex128, len(x))
	for k := 0; k < len(x); k++ {
		Xk := 0+0i
		for n := 0; n < len(x); n++ {
			Xk += x[n] * cmplx.Exp(-i2pi*(complex(float64(k*n), 0)*Ninv))
		}
		X[k] = Ninv * Xk
	}
	return Round(ftPrec, X...)
}

func IDFT(x ...complex128) (X []complex128) {
	N := float64(len(x))
	Ninv := complex(1/N,0)
	X = make([]complex128, len(x))
	for k := 0; k < len(x); k++ {
		Xk := 0+0i
		for n := 0; n < len(x); n++ {
			Xk += x[n] * cmplx.Exp(i2pi*(complex(float64(k*n), 0)*Ninv))
		}
		X[k] = Xk
	}
	return Round(iftPrec, X...)
}

func Round(prec float64, x ...complex128)[]complex128{
	for i := range x{
		x[i] = complex(round(prec, real(x[i])), round(prec, imag(x[i])))
	}
	return x
}

func round(prec float64, v float64) float64{
	iprec := 1/prec
	return math.Round(v*iprec)*prec
}

func ft0(x ...complex128) (X []complex128) {
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
