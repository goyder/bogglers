package engine

import (
	bytes2 "bytes"
	"strings"
	"testing"
)

// Create some standard items that can be used in tests
var items = []string{"DOG", "CAT", "SWAT", "BANE", "QUOTE", "WANE", "COAT", "SWAB", "WAN", "TEA", "OAT", "TAG"}
var letters = [16]string{
"C", "A", "T", "QU",
"D", "O", "G", "O",
"S", "W", "A", "T",
"B", "A", "N", "E",
}
var letterMapping = GenerateNetworkLetterMapping(letters)

// TestSolveBoggleNetwork will test the solution of a Boggle network to get some standard words back.
func TestSolveBoggleNetwork(t *testing.T) {
	// Create a test game
	var buffer bytes2.Buffer
	buffer.WriteString(strings.Join(items, "\n"))
	var dictionary = LoadDictionary(&buffer)
	game := BoggleGame{
		network:       GenerateNetwork(),
		letterMapping: letterMapping,
		dictionary:    dictionary,
	}

	words := SolveBoggleNetwork(game, 3, 10)
	for _, item := range items {
		if !contains(words, item) {
			t.Errorf("Did not find expected word in returned words: %s", item)
		}
	}
}

