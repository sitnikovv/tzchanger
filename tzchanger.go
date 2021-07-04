package tzchanger

import (
	"errors"
	"time"
)

// ChangeTZ change timezone without change time
func ChangeTZ(t time.Time, newTimeZone string) (time.Time, error) {
	var (
		zone *time.Location
		err  error
	)
	if zone, err = time.LoadLocation(newTimeZone); err != nil {
		return t, errors.New("invalid timezone")
	}
	_, currentOffset := t.Zone()
	_, newOffset := t.In(zone).Zone()
	return t.In(zone).Add(time.Second * time.Duration(currentOffset-newOffset)), nil
}
