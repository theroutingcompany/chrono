package chrono

import (
	"time"

	"github.com/golang/protobuf/ptypes"
)

// RFC3339 converted chrono.Time into an RFC3339 string
func RFC3339(t *Time) string {
	return t.T().Format(time.RFC3339)
}

// Max returns the latest of all times given
func Max(times ...time.Time) time.Time {
	maxTime := times[0]

	for _, t := range times {
		if t.After(maxTime) {
			maxTime = t
		}
	}

	return maxTime
}

// Secs converts a float representation of seconds to a time.Duration
func Secs(f float64) time.Duration {
	dur := time.Duration(f * float64(time.Second))
	return dur
}

// TimeDifference returns the Duration between the given start and end times
func TimeDifference(start *Time, end *Time) time.Duration {
	startTime := start.T()
	endTime := end.T()

	return endTime.Sub(startTime)
}

// Add adds a time.Duration to a chrono.Time
func Add(t *Time, dur time.Duration) *Time {
	return New(t.T().Add(dur))
}

// Average returns the average of two *Times
func Average(a *Time, b *Time) *Time {
	diff := a.T().Sub(b.T())
	avg := b.T().Add(diff / 2)
	return New(avg)
}

var (
	// DurPB is a shorthand for this long function
	DurPB = ptypes.DurationProto
	// TimePB is a shorthand for this long function
	TimePB = ptypes.TimestampProto
	// ToTime is a shorthand for this long function
	ToTime = ptypes.Timestamp
	// NowPB is shorthand for this long function
	NowPB = ptypes.TimestampNow
	// Dur is shorthand for this long function
	Dur = ptypes.Duration
)

// DurString converted chrono.Duration into an Influx GROUP BY string
func DurString(d *Duration) string {
	return d.D().String()
}
