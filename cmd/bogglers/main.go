package main

import (
	"flag"
	"github.com/goyder/bogglers/pkg/engine"
	"log"
	"os"
)

func main() {
	// Pull in configs
	minLetters := flag.Int("minLetters", 3, "Minimum letters to allow as a word.")
	maxLetters := flag.Int("maxLetters", 8, "Maximum number of letters to search for.")
	dictionaryPath := flag.String("dict", "", "Filepath to dictionary.txt file.")

	// All flags are declared - call it
	flag.Parse()

	// Read in the dictionary
	dictionaryData, err := os.Open(*dictionaryPath)
	if err != nil {
		log.Fatal(err)
	}
	dictionary := engine.LoadDictionary(dictionaryData)
	dictionaryData.Close()

	// Spin up our inputs
	game := engine.BoggleGame{
		Network:       engine.GenerateNetwork(),
		LetterMapping: engine.GenerateRandomNetworkLetterMapping(1),
		Dictionary:    dictionary,
	}

	// And go
	words := engine.SolveBoggleNetwork(game, *minLetters, *maxLetters)

	// Outputs
	for _, word := range words {
		println(word)
	}
}
