package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"unicode"
)

// Counter struct
type Counter struct {
	values [][]byte
	keys   []int
}

func main() {
	//reading text file
	data, err := ioutil.ReadFile("mobydick.txt")
	if err != nil {
		log.Fatal(err)
	}
	//Trim if its not letter and divide there
	//to create matrix of bytes
	f := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	words := bytes.FieldsFunc(data, f)
	//declaring Struct
	var counter Counter

	counter.uniq(words)
	counter.sort()
	counter.print()
}

//creats uniq words and creats counts of them
func (c *Counter) uniq(words [][]byte) {
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

//checking word for existens, if not exist returnig flag,
//if yes returning index of the word
func exist(source []byte, target [][]byte) int {
	for i, word := range target {
		if bytes.EqualFold(word, source) {
			return i
		}
	}
	return -1
}

//sorting keys and by them sorting values
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

//printing by Stdout Writer
//and before print converting letter to lower case
func (c *Counter) print() {
	for i, word := range c.values[:20] {
		word = bytes.ToLower(word)
		fmt.Printf("%d ", c.keys[i])
		os.Stdout.Write(word)
		fmt.Println()
	}
}
