/*
Working with Time
*/

package main

import (
	"fmt"
	"time"
)

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func current_time() {
	now := time.Now()      // current local time
	sec := now.Unix()      // number of seconds since January 1, 1970 UTC
	nsec := now.UnixNano() // number of nanoseconds since January 1, 1970 UTC

	fmt.Println(now)
	fmt.Println(sec)
	fmt.Println(nsec)
}

// difference between 2 given dates
func date_diff() {
	// The leap year 2000 had 366 days.
	t1 := Date(2000, 1, 1)
	t2 := Date(2001, 1, 1)
	days := t2.Sub(t1).Hours() / 24
	fmt.Println("Start Date:", t1, "\nEnd Date:", t2, "\nDays Diff:", days)
}

// number of days in a month
func days_in_month() {
	t := Date(2000, 3, 0) // the day before 2000-03-01
	fmt.Println(t)        // 2000-02-29 00:00:00 +0000 UTC
}

// main function
func main() {
	// current_time()
	date_diff()
}
