package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {

	t, err := time.Parse("1/2/2006 15:04:05", date)

	if err != nil {
		panic("invalid format")
	}

	return t
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {

	t := time.Now()
	d, err := time.Parse("January 2, 2006 15:04:05", date)

	if err != nil {
		panic("invalid format")
	}

	return t.Sub(d) > 0
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {

	d, err := time.Parse("Monday, January 2, 2006 15:04:05", date)

	if err != nil {
		panic("invalid format")
	}

	return d.Hour() <= 18 && d.Hour() >= 12
}

// Description returns a formatted string of the appointment time
func Description(date string) string {

	d := Schedule(date)

	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.",
		d.Weekday(),
		d.Month().String(),
		d.Day(),
		d.Year(),
		d.Hour(),
		d.Minute(),
	)
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	t := time.Now()
	return time.Date(t.Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}
