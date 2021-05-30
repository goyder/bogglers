package engine

import "strings"

type SearchConfig struct {
	minLetters int
	maxLetters int
}

func contains(items []string, item string) bool {
	for _, s := range items {
		if item == s {
			return true
		}
	}
	return false
}

// We implement our engine as an interface.
type engine interface {
	solve(game *BoggleGame, config *SearchConfig) []string
}

type exhaustiveSearchEngine struct {
}

func (e exhaustiveSearchEngine) solve(game *BoggleGame, config *SearchConfig) []string {
	var visitedNodes []string
	var visitedLetters []string
	var words []string

	for node := range game.Network {
		searchNode(node, game, &(words), &(visitedNodes), &(visitedLetters), config)
	}

	return words
}

// SolveBoggleNetwork is the interface to solve a Boggle game.
func SolveBoggleNetwork(game BoggleGame, engine engine, minLetters int, maxLetters int) []string {
	// Create the storage locations for what we'll catch
	config := SearchConfig{
		minLetters: minLetters,
		maxLetters: maxLetters,
	}

	words := engine.solve(&game, &config)
	return words
}

// searchNode is a recursive call to exhaustively dig through their connections and see if we can find words.
func searchNode(node string, game *BoggleGame, words *[]string, visitedNodes *[]string, visitedLetters *[]string, config *SearchConfig) {
	// Look at the new "word" that has been formed and check if it is a word
	*visitedLetters = append(*visitedLetters, game.LetterMapping[node])
	*visitedNodes = append(*visitedNodes, node)
	newWord := strings.Join(*visitedLetters, "")
	if contains(game.Dictionary, newWord) && len(newWord) >= config.minLetters && !contains(*words, newWord) {
		*words = append(*words, newWord)
	}

	// If this is long enough, stop and go back a level
	if len(*visitedNodes) == config.maxLetters {
		exitSearch(visitedNodes, visitedLetters)
		return
	}

	// Look at all the connected values
	for _, connectedNode := range game.Network[node] {
		// Make sure we're not going somewhere we've already been
		if contains(*visitedNodes, connectedNode) {
			continue
		}

		// If not, check out that node
		searchNode(connectedNode, game, words, visitedNodes, visitedLetters, config)
	}

	// And if we're here, we've done all the searching we can do
	exitSearch(visitedNodes, visitedLetters)
	return
}

func exitSearch(visitedNodes *[]string, visitedLetters *[]string) {
	// We de-reference for readability and simplicity...
	vN, vL := *visitedNodes, *visitedLetters
	*visitedNodes = vN[:(len(vN)-1)]
	*visitedLetters = vL[:(len(vL)-1)]
}