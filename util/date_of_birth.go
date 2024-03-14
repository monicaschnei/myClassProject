package util

import (
	"fmt"
	"time"
)

func TransformDateOfBirth(layout, date string) (time.Time, error) {
	dateOfBirth, err := time.Parse(layout, date)
	if err != nil {
		return time.Time{}, fmt.Errorf("Could not parse the date of birth into time")
	}
	return dateOfBirth, nil
}
