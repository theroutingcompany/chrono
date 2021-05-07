package chrono

import (
	"encoding/json"
	"io"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/sirupsen/logrus"
)

// New instantiates a chrono Time from a Go Time. It panics if the time cannot
// be converted
func New(t time.Time) *Time {
	pt, err := TimePB(t)

	if err != nil {
		logrus.WithError(err).Panic("Cannot instantiate time from Go time")
		return nil
	}

	return NewFromPB(pt)
}

func Must(t *Time, err error) *Time {
	if err != nil {
		logrus.WithError(err).Panic("Cannot return time")
	}

	return t
}

// NewFromPB instantiates a Time from a protocol buffer *timestamp.Timestamp
func NewFromPB(pt *timestamp.Timestamp) *Time {
	return &Time{ProtoTime: pt}
}

// Now instantiates a Time from the current time
func Now() *Time {
	return New(time.Now())
}

// FromString instantiates a Time from a RFC3339 string
func NewFromString(str string) (*Time, error) {
	gt, err := time.Parse(time.RFC3339, str)

	if err != nil {
		return nil, err
	}

	return New(gt), nil
}

// T returns the Time as a time.Time
func (t Time) T() time.Time {
	gt, err := ToTime(t.ProtoTime)

	if err != nil {
		logrus.WithError(err).Panic("Cannot convert from pb.timestamp to time")
	}

	return gt
}

// PB returns the Time as a protocol buffer *timestamp.Timestamp
func (t Time) PB() *timestamp.Timestamp {
	return t.ProtoTime
}

// UnmarshalJSON unmarshals the time encoded as a RFC3339 JSON string into a
// Time
func (t *Time) UnmarshalJSON(data []byte) error {
	gt := &time.Time{}

	if err := gt.UnmarshalJSON(data); err != nil {
		return err
	}

	pt, err := TimePB(*gt)

	if err != nil {
		return err
	}

	t.ProtoTime = pt
	return nil
}

// MarshalJSON marshals the Time into JSON
func (t *Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.T())
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (t *Time) UnmarshalGQL(v interface{}) error {
	gt := &Time{}
	var err error

	if tmpStr, ok := v.(string); ok {
		gt, err = NewFromString(tmpStr)
		if err != nil {
			return err
		}
	}

	*t = *gt
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (t Time) MarshalGQL(w io.Writer) {
	bytes, err := json.Marshal(t.T())

	if err != nil {
		logrus.WithError(err).Panic("Cannot marshal time for GQL")
	}

	if _, err := w.Write(bytes); err != nil {
		logrus.WithError(err).Panic("Cannot marshal time for GQL")
	}
}
