package timeutil

import "time"

// Package timeutil implements utility functions for time series data.

type Observable interface {
	Time() time.Time // Time of observation occurrence
}

type TimeSeries []Observable




