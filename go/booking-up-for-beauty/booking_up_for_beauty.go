package booking

import (
	"fmt"
	"time"
)

var layouts []string = []string{
		"1/_2/2006 15:04:05",
		"01/_2/2006 15:04:05",
		"January _2, 2006 15:04:05",
		"Monday, January _2, 2006 15:04:05",
	}

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	for _, layout := range layouts {
		t, err := time.Parse(layout, date)
		if err == nil {
			return t
		}
	}
	panic(fmt.Sprintf("could not parse date %s with any of the provided layouts", date))
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	now:= time.Now()
	return now.Unix() > Schedule(date).Unix()
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	hour := Schedule(date).Hour()
	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	appointmentStr := Schedule(date).Format("Monday, January 2, 2006, at 15:04.")
	return fmt.Sprint("You have an appointment on " + appointmentStr)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	thisYear := time.Now().Year()
	return time.Date(thisYear, 9, 15, 0, 0, 0, 0, time.UTC)
}
