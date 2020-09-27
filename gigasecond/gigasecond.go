// Package gigasecond solves the gigasecond problem
package gigasecond

import "time"

// AddGigasecond adds a gigasecond to the input time.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1e9)
}
