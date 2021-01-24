package repositories

import (
	"context"
	"mutants/app/mutants"
)

// StrongWriteRepository
// This works similar to a Command Service in a CQRS pattern.
// Responsible for all of the insert/updates in the system
// ensuring ACID
type StrongWriteRepository interface {
	// Save Dna from a human in a datasource
	SaveDna(ctx context.Context, human mutants.Human) error
}
