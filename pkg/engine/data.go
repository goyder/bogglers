package engine

import (
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"strconv"
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

var columnNames = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func LoadDictionary(input io.Reader) []string {
	text, err := ioutil.ReadAll(input)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(text), "\n")
}

func GenerateNetwork() map[string][]string {
	// I heard you like for loops
	network := make(map[string][]string)
	for row := 0; row < 4; row++ {
		for column := 0; column < 4; column++ {
			var connections []string
			for _, columnI := range []int{-1, 0, 1} {
				for _, rowI := range []int{-1, 0, 1} {
					// Work through our reasons *not* to add a connection
					if row+rowI < 0 || row+rowI > 3 {
						continue
					}
					if column+columnI < 0 || column+columnI > 3 {
						continue
					}
					if columnI == 0 && rowI == 0 {
						continue
					}

					// But if we're all good, add the connection
					var target_cell = string(columnNames[column+columnI]) + strconv.FormatInt(int64(row+rowI), 10)
					connections = append(connections, target_cell)
				}
			}
			var cell = string(columnNames[column]) + strconv.FormatInt(int64(row), 10)
			network[cell] = connections
		}
	}
	return network
}

func GenerateNetworkLetterMapping(seed int64) map[string]string {
	rand.Seed(seed)
	letterMapping := make(map[string]string)

	// We'll work through and map our 16 potential node values onto the dice values
	i := 0
	for row := 0; row<4; row++ {
		for col := 0; col<4; col++ {
			node := string(columnNames[col]) + strconv.FormatInt(int64(row), 10)
			letterMapping[node] = dice[i][rand.Intn(5)]
			i++
		}
	}

	return letterMapping
}
