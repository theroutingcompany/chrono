package chrono

import (
	"encoding/json"
	"errors"
	"io"
	"time"

	duration "github.com/golang/protobuf/ptypes/duration"
	"github.com/sirupsen/logrus"
)

// NewDuration instantiates a *chrono.Duration from a time.Duration
func NewDuration(d time.Duration) *Duration {
	return NewDurationFromProto(DurPB(d))
}

func MustDuration(d *Duration, err error) *Duration {
	if err != nil {
		logrus.WithError(err).Panic("Cannog return Duration")
	}

	return d
}

// NewDurationFromProto instantiates a Duration from a protocol buffer
// *duration.Duration
func NewDurationFromProto(pd *duration.Duration) *Duration {
	return &Duration{ProtoDuration: pd}
}

// NewDurationFromString instantiates a Duration from a string
func NewDurationFromString(str string) (*Duration, error) {
	gd, err := time.ParseDuration(str)

	if err != nil {
		return nil, err
	}

	return NewDuration(gd), nil
}

// D returns the Duration as a time.Duration
func (d *Duration) D() time.Duration {
	gd, err := Dur(d.ProtoDuration)

	if err != nil {
		logrus.WithError(err).
			Panic("Cannot convert from pb.duration to duration")
	}

	return gd
}

// PB returns the Duration as protocol buffer *duration.Duration
func (d *Duration) PB() *duration.Duration {
	return d.ProtoDuration
}

// UnmarshalJSON unmarshals the duration encoded as JSON into a Duration
func (d *Duration) UnmarshalJSON(data []byte) error {
	var v interface{}

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	var dur time.Duration

	switch value := v.(type) {
	case float64:
		dur = time.Duration(value)
	case string:
		var err error
		dur, err = time.ParseDuration(value)

		if err != nil {
			return err
		}
	default:
		return errors.New("invalid duration")
	}

	d.ProtoDuration = DurPB(dur)
	return nil
}

// MarshalJSON marshals the Duration into JSON
func (d *Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.D().String())
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (d *Duration) UnmarshalGQL(v interface{}) error {
	gd := &Duration{}
	var err error

	if tmpStr, ok := v.(string); ok {
		gd, err = NewDurationFromString(tmpStr)
		if err != nil {
			return err
		}
	}

	*d = *gd
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (d *Duration) MarshalGQL(w io.Writer) {
	bytes, err := json.Marshal(d.D())

	if err != nil {
		logrus.WithError(err).Panic("Cannot marshal duration for GQL")
	}

	if _, err := w.Write(bytes); err != nil {
		logrus.WithError(err).Panic("Cannot marshal duration for GQL")
	}
}
