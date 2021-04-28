package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"unicode"
)

// Counter .
type Counter struct {
	values [][]byte
	keys   []int
}

func main() {
	data, err := ioutil.ReadFile("mobydick.txt")
	if err != nil {
		log.Fatal(err)
	}

	f := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	var counter Counter
	words := bytes.FieldsFunc(data, f)

	counter.count(words)
	counter.sort()
	print(counter.values[:20], counter.keys[:20])
}

func (c *Counter) count(words [][]byte) {
	for _, word := range words {
		n := exist(word, c.values)
		if n == -1 {
			c.values = append(c.values, word)
			c.keys = append(c.keys, 1)
		} else {
			c.keys[n] = c.keys[n] + 1
		}
	}
}

func exist(source []byte, target [][]byte) int {
	for i, word := range target {
		if bytes.EqualFold(word, source) {
			return i
		}
	}
	return -1
}

func (c *Counter) sort() {
	for {
		isChanged := false

		for k := 0; k < len(c.keys); k++ {
			if k+1 == len(c.keys) {
				break
			}
			if c.keys[k] < c.keys[k+1] {
				c.keys[k], c.keys[k+1] = c.keys[k+1], c.keys[k]
				c.values[k], c.values[k+1] = c.values[k+1], c.values[k]
				isChanged = true
			}
		}
		if !isChanged {
			break
		}
	}
}

type Writer interface {
	Write(p []byte) /*(n int, err error)*/
}

func print(words [][]byte, keys []int) {
	for i, word := range words {
		word = bytes.ToLower(word)
		fmt.Printf("%d ", keys[i])
		f := os.Stdout
		f.Write(word)
		fmt.Println()
	}
}
