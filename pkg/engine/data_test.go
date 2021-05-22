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

func TestNetworkGeneration(t *testing.T) {
	// Produce a basic network
	network := GenerateNetwork()

	// Basic assessment
	if len(network) != 16 {
		t.Errorf("Did not find 16 items in the network.")
	}
}

func TestGenerateNetworkAndAssessNodes(t *testing.T) {
	// Produce a network
	network := GenerateNetwork()

	// Assess whether key nodes have the right connections
	tests := []struct {
		node        string
		connections []string
	}{
		{
			node: "A0",
			connections: []string{
				"A1", "B0", "B1",
			},
		},
		{
			node: "B1",
			connections: []string{
				"A0", "A1", "A2", "B0", "B2", "C0", "C1", "C2",
			},
		},
		{
			node: "D3",
			connections: []string{
				"D2", "C3", "C2",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.node+" connections",
			func(t *testing.T) {
				// First run - did we have the right number of connections?
				var expectedConnections = len(test.connections)
				var actualConnections = len(network[test.node])
				if expectedConnections != actualConnections {
					t.Errorf("Did not get expected number of connections on node %s. Expected %d, got %d",
						test.node, expectedConnections, actualConnections)
				}

				// Next - were they the same?
				for _, connection := range network[test.node] {
					if !contains(test.connections, connection) {
						t.Errorf("Did not get expected connection %s in network.",
							connection)
					}
				}
			})
	}
}

func TestGenerateNetworkLetterMapping(t *testing.T) {
	// Just do the darn thing
	letterMapping := GenerateNetworkLetterMapping(0)

	// Assess basic dimensions
	if len(letterMapping) != 16 {
		t.Errorf("Did not get expected number of items in letter mapping. Expected %d, got %d",
			16, len(letterMapping))
	}
}

func TestGenerateNetworkLetterMappingDoesNotRepeat(t *testing.T) {
	// A simple error would be for the same "die" to be called for all values.
	// Do a basic sense check to ensure this hasn't occurred.

	letterMapping := GenerateNetworkLetterMapping(0)

	// Extract all the values - not really necessary but a good exploration in unrolling maps
	var values []string
	for _, v := range letterMapping {
		values = append(values, v)
	}

	// Ensure that we have different values...
	identicalValues := true
	for i := 1; i < 16; i++ {
		if values[i] != values[0] {
			identicalValues = false
		}
	}
	if identicalValues {
		t.Errorf("All values in letter mapping were the same letter.")
	}
}