package main

import (
	"reflect"
	"testing"
)

func TestVectors(t *testing.T) {
	for _, tc := range []struct {
		i []complex128
		o []complex128
	}{
		{i: []complex128{1, 0, 0, 0, 0, 0, 0, 0}, o: []complex128{.125, .125, .125, .125, .125, .125, .125, .125}},
		{i: []complex128{0, 1, 0, 0, 0, 0, 0, 0}, o: []complex128{0.125, 0.088, 0, -0.088, -0.125, -0.088, 0, 0.088}},
	} {
		have := FT(tc.i...)
		if want := tc.o; !reflect.DeepEqual(have, want) {
			t.Logf("bad output: have %v want %v", have, want)
		}
	}
}
