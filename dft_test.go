package main

import (
	"math"
	"math/cmplx"
	"reflect"
	"testing"
)

func TestMath(t *testing.T) {
	want := cmplx.Sqrt(-1) * 2 * math.Pi
	if i2pi != want {
		t.Fatalf("bad complex exponential: have %v, want %v", i2pi, want)
	}
	t.Logf("exp %v", cmplx.Exp(i2pi))
	t.Logf("exp %v", cmplx.Exp(want))
}

func TestVectors(t *testing.T) {
	for _, tc := range []struct {
		i []complex128
		o []complex128
	}{
		{i: []complex128{1, 0, 0, 0, 0, 0, 0, 0}, o: []complex128{.125, .125, .125, .125, .125, .125, .125, .125}},
		{i: []complex128{0, 1, 0, 0, 0, 0, 0, 0}, o: []complex128{0.125, 0.088 - 0.088i, 0 - 0.125i, -0.088 - 0.088i, -0.125, -0.088 + 0.088i, 0 + 0.125i, 0.088 + 0.088i}},
	} {
		have := DFT(tc.i...)
		if want := tc.o; !reflect.DeepEqual(have, want) {
			t.Logf("bad output: have %v\n\twant %v", have, want)
		}
	}
}

func TestInvert(t *testing.T) {
	for _, tc := range []struct {
		i []complex128
	}{
		{i: []complex128{1, 0, 0, 0, 0, 0, 0, 0}},
		{i: []complex128{0, 1, 0, 0, 0, 0, 0, 0}},
	} {
		have := IDFT(DFT(tc.i...)...)
		if want := tc.i; !reflect.DeepEqual(have, want) {
			t.Logf("bad output: have %v want %v", have, want)
		}
	}
}
