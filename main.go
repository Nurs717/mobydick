package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"unicode"
)

var values [][]byte

func main() {
	data, err := ioutil.ReadFile("mobydick.txt")
	if err != nil {
		fmt.Println(err)
	}

	f := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	words := bytes.FieldsFunc(data, f)

	var counters []int

	for _, word := range words {
		n := exist(word, values)
		if n == -1 {
			values = append(values, word)
			counters = append(counters, 1)
		} else {
			counters[n] = counters[n] + 1
		}
	}

	sort(counters, values)

	print(values[:20], counters[:20])

}

func exist(source []byte, target [][]byte) int {
	for i, word := range target {
		if bytes.EqualFold(word, source) {
			return i
		}
	}
	return -1
}

func sort(counters []int, values [][]byte) {
	for {
		isChanged := false

		for k := 0; k < len(counters); k++ {
			if k+1 == len(counters) {
				break
			}
			if counters[k] < counters[k+1] {
				counters[k], counters[k+1] = counters[k+1], counters[k]
				values[k], values[k+1] = values[k+1], values[k]
				isChanged = true
			}
		}
		if !isChanged {
			break
		}
	}
}

func print(words [][]byte, counters []int) {
	for i, word := range words {
		word = bytes.ToLower(word)
		fmt.Printf("%v ", counters[i])
		for _, letter := range word {
			fmt.Printf("%v", rune(letter))
		}
		fmt.Println()
	}
}
