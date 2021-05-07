package chrono

import "time"

// RelativeRange is a relative time range defined by the upper and lower
// durations. These values are used to create an absolute range centered
// around a time (like the current time)
type RelativeRange struct {
	Lower, Upper time.Duration
}

// Abs returns an absolute time range from the relative range centered
// around the provided time
func (r RelativeRange) Abs(t time.Time) Range {
	return Range{t.Add(r.Lower), t.Add(r.Upper)}
}

// Range defines a time range using a lower and upper time bound
type Range struct {
	Lower, Upper time.Time
}

// WithinRange returns whether or not a given time is within a time range
func WithinRange(t time.Time, rng Range) bool {
	return t.After(rng.Lower) && t.Before(rng.Upper)
}
