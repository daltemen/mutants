package rest

type DnaRequest struct {
	Dna []string `json:"dna"`
}

type StatsResponse struct {
	CountMutantDna int64   `json:"count_mutant_dna"`
	CountHumanDna  int64   `json:"count_human_dna"`
	Ratio          float32 `json:"ratio"`
}
