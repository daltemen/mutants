package repositories

import (
	"context"
	"mutants/app/mutants"
)

// ReadRepository
// This works similar to a Query Service in a CQRS pattern.
// Responsible for all of the reads in the system.
type ReadRepository interface {
	// Get Stats from the datasource
	GetStats(ctx context.Context) (mutants.HumanStats, error)
}
