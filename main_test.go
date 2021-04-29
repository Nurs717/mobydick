package main

import (
	"bytes"
	"testing"
)

func TestExist(t *testing.T) {
	cases := []struct {
		target [][]byte
		source []byte
		result int
	}{
		{target: [][]byte{[]byte("hello"), []byte("world")}, source: []byte("hello"), result: 0},
		{target: [][]byte{[]byte("hello"), []byte("world")}, source: []byte("world"), result: 1},
		{target: [][]byte{[]byte("hello"), []byte("world")}, source: []byte("byebye"), result: -1},
	}

	for _, test := range cases {
		if res := exist(test.source, test.target); res != test.result {
			t.Errorf("expect for %v but got %v", test.result, res)
		}
	}

}

func TestUniq(t *testing.T) {
	cases := []struct {
		counter      Counter
		words        [][]byte
		resultValues [][]byte
		resultKeys   []int
	}{
		{words: [][]byte{[]byte("menin"), []byte("atym"), []byte("Qoja")},
			resultValues: [][]byte{[]byte("menin"), []byte("atym"), []byte("Qoja")},
			resultKeys:   []int{1, 1, 1},
		},

		{words: [][]byte{[]byte("menin"), []byte("atym"), []byte("menin"), []byte("menin"), []byte("Qoja"), []byte("Qoja"), []byte("atym"), []byte("menin"), []byte("Qoja")},
			resultValues: [][]byte{[]byte("menin"), []byte("atym"), []byte("Qoja")},
			resultKeys:   []int{4, 2, 3},
		},

		{words: [][]byte{[]byte("menin")},
			resultValues: [][]byte{[]byte("menin")},
			resultKeys:   []int{1},
		},
	}

	for _, test := range cases {

		test.counter.uniq(test.words)

		equalMatrix(test.counter.values, test.resultValues)
		equalInt(test.counter.keys, test.resultKeys)

	}
}

func TestSort(t *testing.T) {
	cases := []struct {
		counter      Counter
		resultValues [][]byte
		resultKeys   []int
	}{
		{
			counter: Counter{
				values: [][]byte{[]byte("menin"), []byte("atym"), []byte("Qoja")},
				keys:   []int{4, 2, 3},
			},
			resultValues: [][]byte{[]byte("menin"), []byte("Qoja"), []byte("atym")},
			resultKeys:   []int{4, 3, 2},
		},
	}

	for _, test := range cases {
		test.counter.sort()
		if !equalInt(test.resultKeys, test.counter.keys) {
			t.Errorf("expect for %v but got %v", test.resultKeys, test.counter.keys)
		}

		if !equalMatrix(test.resultValues, test.counter.values) {
			t.Errorf("expect for %v but got %v", test.resultValues, test.counter.values)
		}
	}

}

func equalInt(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, num := range a {
		if num != b[i] {
			return false
		}
	}
	return true
}

func equalMatrix(a, b [][]byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, num := range a {
		if !bytes.EqualFold(num, b[i]) {
			return false
		}
	}
	return true
}
