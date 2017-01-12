package timeutil

import "time"

// Package timeutil implements utility functions for time series data.

type Observable interface {
	Time() time.Time // Time of observation occurrence
}

type TimeSeries []Observable

// Range returns observables in range [start, finish].
func (ts TimeSeries) Range(start, finish time.Time) TimeSeries {
    var startIdx, endIdx int
    for i := 1; i < len(ts); i++ {
        if ts[i].After(start) || ts[i].Equal(start) { 
            startIdx = i
            break
        }
    }
    for j := len(ts)-1; j > i; j-- {
                if ts[j].Before(finish) || ts[j].Equal(finish) { 
            endIdx = j
            break
        }
    }
    
    return ts[startIdx:endIdx+1]
}


