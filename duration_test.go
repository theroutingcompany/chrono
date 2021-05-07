package chrono

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testDurJSON struct {
	Duration *Duration `json:"duration"`
}

var (
	durStr     = "15s"
	dur        = 15 * time.Second
	encodedDur = []byte(fmt.Sprintf(`{"duration":"%s"}`, durStr))
	decodedDur = testDurJSON{
		Duration: MustDuration(NewDurationFromString(durStr)),
	}
)

func TestDuration_Unmarshal(t *testing.T) {
	re := require.New(t)
	as := assert.New(t)

	foo := testDurJSON{}
	err := json.Unmarshal(encodedDur, &foo)
	re.Nil(err)
	as.Equal(decodedDur.Duration.D(), foo.Duration.D())

	t.Log(foo)
	t.Log(foo.Duration.D())
	t.Log(foo.Duration.PB())
}

func TestDuration_Marshal(t *testing.T) {
	re := require.New(t)
	as := assert.New(t)

	foo, err := json.Marshal(decodedDur)
	re.Nil(err)
	as.Equal(encodedDur, foo)

	t.Log(foo)
}

func TestDuration_FromString(t *testing.T) {
	actual, err := NewDurationFromString(durStr)
	require.Nil(t, err)

	assert.Equal(t, dur, actual.D())
}
