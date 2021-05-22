package engine

import (
	"io"
	"io/ioutil"
	"log"
	"strings"
)

var dice = [][]string{
	{"P", "S", "A", "F", "K", "F"},
	{"H", "I", "U", "N", "QU", "M"},
	{"I", "M", "T", "O", "U", "C"},
	{"I", "T", "S", "T", "D", "Y"},
	{"L", "R", "E", "I", "X", "D"},
	{"E", "E", "U", "S", "I", "N"},
	{"E", "R", "W", "T", "H", "V"},
	{"T", "Y", "E", "L", "R", "T"},
	{"N", "G", "E", "E", "A", "A"},
	{"T", "O", "E", "S", "S", "I"},
	{"B", "B", "A", "O", "O", "S"},
	{"H", "E", "E", "N", "H", "E"},
	{"W", "T", "O", "O", "T", "A"},
	{"S", "O", "A", "C", "H", "P"},
	{"R", "N", "Z", "N", "H", "L"},
	{"D", "E", "Y", "L", "V", "R"},
}

func LoadDictionary(input io.Reader) []string {
	text, err := ioutil.ReadAll(input)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(text), "\n")
}
