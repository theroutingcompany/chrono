package chrono

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testTimeJSON struct {
	Time *Time `json:"time"`
}

var (
	timeStr     = "2020-05-07T00:00:00Z"
	timeGo, _   = time.Parse(time.RFC3339, timeStr)
	encodedTime = []byte(fmt.Sprintf(`{"time":"%s"}`, timeStr))
	decodedTime = testTimeJSON{
		Time: Must(NewFromString(timeStr)),
	}
)

func TestTime_Unmarshal(t *testing.T) {
	re := require.New(t)
	as := assert.New(t)

	foo := testTimeJSON{}
	err := json.Unmarshal(encodedTime, &foo)
	re.Nil(err)
	as.Equal(decodedTime.Time.T(), foo.Time.T())

	t.Log(foo)
	t.Log(foo.Time.T())
	t.Log(foo.Time.PB())
}

func TestTime_Marshal(t *testing.T) {
	re := require.New(t)
	as := assert.New(t)

	foo, err := json.Marshal(decodedTime)
	re.Nil(err)
	as.Equal(encodedTime, foo)

	t.Log(foo)
}

func TestTime_FromString(t *testing.T) {
	actual, err := NewFromString(timeStr)
	require.Nil(t, err)

	assert.Equal(t, timeGo, actual.T())
}
