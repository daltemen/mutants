// Mutant is the main domain Entity
package mutants


// Human Business Entity
// Dna: Contains the dna represented as a matrix [nxn]
// 	each row contains a sequence of letters
// 	the letters just can be (A,T,C,G) other letters are not allowed
//	i.e: {"ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"};
// expressed as
// {ATGCGA}
// {CAGTGC}
// {TTATGT}
// {AGAAGG}
// {CCCCTA}
// {TCACTG}
type Human struct {
	Dna [][]string
}

// Build a new Human
func NewHuman(dna [][]string) Human {
	human := Human{
		Dna: dna,
	}
	return human
}

// It checks if the human is a mutant
// A human is a mutant if he has in his dna one of the below rules:
//
// 1. 4 equal letters horizontally
//
// 2. 4 equal letters vertically
//
// 3. 4 equal letters in the diagonal
func (h *Human) IsMutant() bool {
	return h.checkIsMutant()
}

// to check the pattern to recognize a mutant
// in the worst case we have an O(n^2) complexity
// if the performance is a concern
// this should be use coroutines to improve the performance
func (h *Human) checkIsMutant() bool {
	size := len(h.Dna)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			maxIndex := size - 1
			maxLengthRow := j+3 <= maxIndex
			if maxLengthRow {
				isMutant := h.checkHorizontally(i, j)
				if isMutant {
					return true
				}
			}
			maxLengthColumn := i+3 <= maxIndex
			if maxLengthColumn {
				isMutant := h.checkVertically(i, j)
				if isMutant {
					return true
				}
				isMutant = h.checkDiagonal(i)
				if isMutant {
					return true
				}
			}
			if !maxLengthColumn && !maxLengthRow {
				break
			}
		}
	}
	return false
}

func (h *Human) checkHorizontally(i int, j int) bool {
	if h.Dna[i][j] == h.Dna[i][j+1] && h.Dna[i][j] == h.Dna[i][j+2] && h.Dna[i][j] == h.Dna[i][j+3] {
		println("found horizontal ")
		return true
	}
	return false
}

func (h *Human) checkVertically(i int, j int) bool {
	if h.Dna[i][j] == h.Dna[i+1][j] && h.Dna[i][j] == h.Dna[i+2][j] && h.Dna[i][j] == h.Dna[i+3][j] {
		println("found vertical ")
		return true
	}
	return false
}

func (h *Human) checkDiagonal(i int) bool {
	if h.Dna[i][i] == h.Dna[i+1][i+1] && h.Dna[i][i] == h.Dna[i+2][i+2] && h.Dna[i][i] == h.Dna[i+3][i+3] {
		println("found diagonal ")
		return true
	}
	return false
}
