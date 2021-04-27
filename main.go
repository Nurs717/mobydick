package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"unicode"
)

func main() {
	data, err := ioutil.ReadFile("mobydick.txt")
	if err != nil {
		fmt.Println(err)
	}
	var n int
	coter := 0
	f := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	a := bytes.FieldsFunc(data, f)
	for i, word := range a {
		coter = 0
		for _, words := range a {
			if bytes.EqualFold(word, words) {
				coter++
			}
		}
		fmt.Println(string(a[i]), coter)
	}

	for i, bytes := range data {
		if !(bytes >= 'a' && bytes <= 'z' || bytes >= 'A' && bytes <= 'Z' /* || bytes == ' ' || bytes == '\n'*/) {
			if n == 0 {
				n = i
			} else {
				// e := data[n+1 : i]
				n = i
				// fmt.Printf("%s ", string(e))
				// c := counter(data, e)
				// fmt.Println(c)
			}

		}
	}
}

func counter(data, word []byte) int {
	var count int
	var n int
	for i, bytesd := range data {
		if !(bytesd >= 'a' && bytesd <= 'z' || bytesd >= 'A' && bytesd <= 'Z' /* || bytes == ' ' || bytes == '\n'*/) {
			if n == 0 {
				n = i
			} else {
				e := data[n+1 : i]
				n = i
				res := bytes.EqualFold(e, word)
				if res {
					count++
				}
			}

		}
	}
	return count
}
