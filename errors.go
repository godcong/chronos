package chronos

import "errors"

var (
	// ErrWrongSolarTermFormat is returned when an invalid SolarTerm value is
	// provided.
	ErrWrongSolarTermFormat = errors.New("[chronos] wrong solar term format")
)
