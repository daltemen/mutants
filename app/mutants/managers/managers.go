package managers

import "context"

// MutantManager
// Is the interface to handle all of the mutants business rules
// Is the way to access to the domain layer
type MutantManager interface {
	// Is the way how to know if a human is either a mutant or not
	// dna: is a matrix with the dna of the human
	// just (A,T,C,G) characters are supported
	// i.e. of a dna matrix
	// {"A", "T", "G", "C", "G", "A"},
	// {"C", "A", "G", "T", "G", "C"},
	// {"T", "T", "A", "T", "G", "T"},
	// {"A", "G", "A", "A", "G", "G"},
	// {"C", "C", "C", "C", "T", "A"},
	// {"T", "C", "A", "C", "T", "G"},
	IsMutant(ctx context.Context, dna [][]string) (bool, error)
	// Obtains all the stats like number of mutants, humans and other useful stats
	RetrieveStats(ctx context.Context) (StatsSummary, error)
}

// StatsSummary holds all the important stats
type StatsSummary struct {
	Mutants int64
	Humans  int64
	Ratio   float32
}
