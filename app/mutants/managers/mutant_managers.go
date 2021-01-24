package managers

import (
	"context"
	"mutants/app/mutants"
	"mutants/app/mutants/repositories"
)

type mutantManager struct {
	readRepo        repositories.ReadRepository
	strongWriteRepo repositories.StrongWriteRepository
	eventualRepo    repositories.EventualWriteRepository
}

func NewMutantManager(readRepo repositories.ReadRepository, strongRepo repositories.StrongWriteRepository, eventualRepo repositories.EventualWriteRepository) MutantManager {
	return &mutantManager{readRepo: readRepo, strongWriteRepo: strongRepo, eventualRepo: eventualRepo}
}

func (m *mutantManager) IsMutant(ctx context.Context, dna [][]string) (bool, error) {
	human := mutants.NewHuman(dna)

	savedErr := m.strongWriteRepo.SaveDna(ctx, human)
	if savedErr != nil {
		return false, savedErr // TODO: handle error
	}
	humanType := human.GetType()

	countsErr := m.eventualRepo.IncrementCount(ctx, humanType)
	if countsErr != nil {
		return false, countsErr // TODO: handle error
	}

	return humanType == mutants.MutantType, nil
}

func (m *mutantManager) RetrieveStats(ctx context.Context) (StatsSummary, error) {
	stats, err := m.readRepo.GetStats(ctx)
	if err != nil {
		return StatsSummary{}, err // TODO: handle error
	}

	return StatsSummary{
		Mutants: stats.Mutants,
		Humans:  stats.Humans,
		Ratio:   float32(float64(stats.Humans)/float64(stats.Mutants)),
	}, nil
}
