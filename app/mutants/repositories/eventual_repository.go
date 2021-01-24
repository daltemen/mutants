package repositories

import (
	"context"
	"mutants/app/mutants"
)

// EventualWriteRepository
// This works similar to a Command Service in a CQRS pattern.
// Responsible for all of the insert/updates in the system.
// more related to CAP
type EventualWriteRepository interface {
	// Increment Count using a count strategy
	IncrementCount(ctx context.Context, humanType mutants.HumanTypes) error
}
