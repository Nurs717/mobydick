package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"unicode"
)

var counter int

func main() {
	data, err := ioutil.ReadFile("mobydick.txt")
	if err != nil {
		fmt.Println(err)
	}
	var n int
	var test [][]byte
	f := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	words := bytes.FieldsFunc(data, f)
	test = append(test, words[0])
	test = append(test, loop(words, test[0+counter]))

	// LOOP:
	// for {
	// for _, word := range test {
	// 	for _, words := range a {
	// 		if !bytes.EqualFold(word, words) {
	// 			// fmt.Printf("words: %s, ", word)
	// 			if counter == 0 {
	// 				test = append(test, words)
	// 				counter++
	// 			}
	// 		}
	// 	}
	// 	counter = 0
	// if word == nil {
	// 	break LOOP
	// }
	// fmt.Println(string(a[i]), coter)
	// }
	// }
	// fmt.Println(counter)

	for _, val := range test {
		fmt.Printf("%s, ", string(val))
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

func loop(words [][]byte, test []byte) []byte {
	var word2 []byte
	for _, word := range words {
		if !bytes.EqualFold(test, word) {
			word2 = word
			counter++
			// break
			return word2
		}
	}
	return word2
}

// func count(data, word []byte) int {
// 	var count int
// 	var n int
// 	for i, bytesd := range data {
// 		if !(bytesd >= 'a' && bytesd <= 'z' || bytesd >= 'A' && bytesd <= 'Z' /* || bytes == ' ' || bytes == '\n'*/) {
// 			if n == 0 {
// 				n = i
// 			} else {
// 				e := data[n+1 : i]
// 				n = i
// 				res := bytes.EqualFold(e, word)
// 				if res {
// 					count++
// 				}
// 			}

// 		}
// 	}
// 	return count
// }
