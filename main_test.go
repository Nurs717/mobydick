package main

import (
	"bytes"
	"testing"
)

type Check struct {
	words        [][]byte
	resultvalues [][]byte
	resultkeys   []int
}

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
	var testCounter Counter

	check := []Check{
		{words: [][]byte{[]byte("menin"), []byte("atym"), []byte("Qoja")},
			resultvalues: [][]byte{[]byte("menin"), []byte("atym"), []byte("Qoja")},
			resultkeys:   []int{1, 1, 1},
		},

		{words: [][]byte{[]byte("menin"), []byte("atym"), []byte("menin"), []byte("menin"), []byte("Qoja"), []byte("Qoja"), []byte("atym"), []byte("menin"), []byte("Qoja")},
			resultvalues: [][]byte{[]byte("menin"), []byte("atym"), []byte("Qoja")},
			resultkeys:   []int{4, 2, 3},
		},

		{words: [][]byte{[]byte("menin")},
			resultvalues: [][]byte{[]byte("menin")},
			resultkeys:   []int{1},
		},
	}

	for _, test := range check {

		testCounter.uniq(test.words)

		testCounter.equalMatrix(&test, t)
		testCounter.equalInt(&test, t)

		testCounter.values = [][]byte{}
		testCounter.keys = []int{}

	}
}

func TestSort(t *testing.T) {
	testCounter := []Counter{
		{
			values: [][]byte{[]byte("menin"), []byte("atym"), []byte("Qoja")},
			keys:   []int{4, 2, 3},
		},
		{
			values: [][]byte{[]byte("menin")},
			keys:   []int{4},
		},
	}

	check := []Check{
		{
			resultvalues: [][]byte{[]byte("menin"), []byte("Qoja"), []byte("atym")},
			resultkeys:   []int{4, 3, 2},
		},
		{
			resultvalues: [][]byte{[]byte("menin")},
			resultkeys:   []int{4},
		},
	}

	for i, test := range check {
		testCounter[i].sort()

		testCounter[i].equalMatrix(&test, t)
		testCounter[i].equalInt(&test, t)

	}
}

func (c *Counter) equalMatrix(b *Check, t *testing.T) {
	for i, word := range c.values {
		if !bytes.EqualFold(word, b.resultvalues[i]) {
			t.Errorf("expect for %v but got %v", string(b.resultvalues[i]), string(word))
			break
		}
	}

}

func (c *Counter) equalInt(b *Check, t *testing.T) {
	for i, num := range c.keys {
		if num != b.resultkeys[i] {
			t.Errorf("expect for %d but got %d", b.resultkeys[i], num)
			break
		}
	}

}
