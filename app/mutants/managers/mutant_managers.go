package managers

import (
	"context"
	"log"
	"mutants/app/mutants"
	"mutants/app/mutants/repositories"
)

type mutantManager struct {
	readRepo          repositories.ReadRepository
	strongWriteRepo   repositories.StrongWriteRepository
	eventualWriteRepo repositories.EventualWriteRepository
}

func NewMutantManager(readRepo repositories.ReadRepository, strongRepo repositories.StrongWriteRepository, eventualRepo repositories.EventualWriteRepository) MutantManager {
	return &mutantManager{readRepo: readRepo, strongWriteRepo: strongRepo, eventualWriteRepo: eventualRepo}
}

func (m *mutantManager) IsMutant(ctx context.Context, dna [][]string) (bool, error) {
	human := mutants.NewHuman(dna)

	savedErr := m.strongWriteRepo.SaveDna(ctx, human)
	if savedErr != nil {
		log.Println(savedErr)
		return false, &ManagerErrors{Description: "SaveDna Error on the Strong Write Repo failed", Err: SaveRowError}
	}
	humanType := human.GetType()

	countsErr := m.eventualWriteRepo.IncrementCount(ctx, humanType)
	if countsErr != nil {
		log.Println(countsErr)
		return false, &ManagerErrors{Description: "IncrementCount on Eventual Write Repo failed", Err: CountsError}
	}

	return humanType == mutants.MutantType, nil
}

func (m *mutantManager) RetrieveStats(ctx context.Context) (StatsSummary, error) {
	stats, err := m.readRepo.GetStats(ctx)
	if err != nil {
		log.Println(err)
		return StatsSummary{}, &ManagerErrors{Description: "GetStats on the Read Repo failed", Err: StatsError}
	}

	return StatsSummary{
		Mutants: stats.Mutants,
		Humans:  stats.Humans,
		Ratio:   float32(float64(stats.Mutants) / float64(stats.Humans)),
	}, nil
}
