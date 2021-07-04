package tzchanger

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestChangeTZ(t *testing.T) {
	const cleanTime = "2006-01-02 15:04:05"
	utcTime, _ := time.Parse(cleanTime, "2021-07-01 05:04:03")
	assert.Equal(t, "2021-07-01 05:04:03 +0000 UTC", utcTime.String(), "utcTime failed")
	_, err := ChangeTZ(utcTime, "Unknown/Zone")
	assert.EqualError(t, err, "invalid timezone", "invalid timezone failed")
	mskTime, _ := ChangeTZ(utcTime, "Europe/Moscow")
	assert.Equal(t, "2021-07-01 05:04:03 +0300 MSK", mskTime.String(), "mskTime failed")
	vldTime, _ := ChangeTZ(mskTime, "Asia/Vladivostok")
	assert.Equal(t, "2021-07-01 05:04:03 +1000 +10", vldTime.String(), "vldTime failed")
	pdtTime, _ := ChangeTZ(vldTime, "America/Los_Angeles")
	assert.Equal(t, "2021-07-01 05:04:03 -0700 PDT", pdtTime.String(), "pdtTime failed")
}
