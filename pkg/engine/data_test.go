package engine

import (
	bytes2 "bytes"
	"strings"
	"testing"
)

func contains(items []string, item string) bool {
	for _, s := range items {
		if item == s {
			return true
		}
	}
	return false
}

func TestLoadDictionary(t *testing.T) {
	// Produce a test list of items
	var items = []string{"who", "what", "where"}

	var buffer bytes2.Buffer
	buffer.WriteString(strings.Join(items, "\n"))

	// Test that we get back what we expected
	output := LoadDictionary(&buffer)
	for _, item := range items {
		t.Run(item+" is in returned list",
			func(t *testing.T) {
				if !contains(output, item) {
					t.Errorf("Did not find %s in output list.", item)
				}
			})
	}
}
