package gigasecond

import "time"

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	giga := t.Add(time.Second * 1000000000)
	return giga
}
