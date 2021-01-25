package managers

import "errors"

var (
	SaveRowError = errors.New("dna could not saved")
	CountsError  = errors.New("counts could not updated")
	StatsError   = errors.New("stats are not available")
)

type ManagerErrors struct {
	Description string
	Err         error
}

func (e *ManagerErrors) Error() string { return e.Description + ": " + e.Err.Error() }

func (e *ManagerErrors) Unwrap() error { return e.Err }
