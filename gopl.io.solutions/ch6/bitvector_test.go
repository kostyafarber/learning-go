package main

import (
	"reflect"
	"testing"

	BitVector "gopl.io.solutons/ch6/bitvector"
)

func TestIntAdd(t *testing.T) {
	tests := map[string]struct {
		got  []int
		add  int
		want []int
	}{
		"add one number": {[]int{1, 2, 3}, 1, []int{1, 2, 3, 4}},
	}

	for name, tt := range tests {
		var vector_got BitVector.IntSet
		vector_got.AddAll(tt.got...)
		vector_got.Add(tt.add)

		var vector_want BitVector.IntSet
		vector_want.AddAll(tt.want...)

		if !reflect.DeepEqual(vector_got, vector_want) {
			t.Fatalf("%s: expected: %s, got: %s", name, vector_want, vector_got)
		}

	}
}
