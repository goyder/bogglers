package main

import (
	"flag"
	"github.com/goyder/bogglers/pkg/engine"
	"log"
	"os"
	"strings"
)

func main() {
	// Pull in configs
	minLetters := flag.Int("minLetters", 3, "Minimum letters to allow as a word.")
	maxLetters := flag.Int("maxLetters", 6, "Maximum number of letters to search for.")
	dictionaryPath := flag.String("dict", "", "Filepath to dictionary.txt file.")
	boggleBoard := flag.String("board", "", "Implement Boggle board as text, letters comma-separated.")

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
	// First, check if a Boggle board has been provided
	randomBoard := true
	boggleLetters := strings.Split(*boggleBoard, ",")
	var boggleLettersArray [16]string
	if len(boggleLetters) == 16 {
		randomBoard = false
		copy(boggleLettersArray[:], boggleLetters[:16])
	}

	var letterMapping map[string]string
	if randomBoard {
		letterMapping = engine.GenerateRandomNetworkLetterMapping(1)
	} else {
		letterMapping = engine.GenerateNetworkLetterMapping(boggleLettersArray)
	}

	game := engine.BoggleGame{
		Network:       engine.GenerateNetwork(),
		LetterMapping: letterMapping,
		Dictionary:    dictionary,
	}
	var trieEngine = engine.TrieSearchEngine{}

	// And go
	println("Solving board:")
	game.DisplayBoard()
	println("================")
	println("Beginning execution.")
	words := engine.SolveBoggleNetwork(game, trieEngine, *minLetters, *maxLetters)

	// Outputs
	println("Execution complete. Words found:")
	println("================")
	for _, word := range words {
		println(word)
	}
}
