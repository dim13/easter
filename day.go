package calendar

import "time"

/*
  6/15		June 15 (if ambiguous, will default to month/day).
  Jun. 15	June 15.
  15 June	June 15.
  Thursday	Every Thursday.
  June		Every June 1st.
  15 *		15th of every month.
  May Sun+2	second Sunday in May (Muttertag)
  04/SunLast	last Sunday in April, summer time in Europe
  Easter	Easter
  Easter-2	Good Friday (2 days before Easter)
  Paskha	Orthodox Easter
*/

func Parse(s string) time.Time {
	return time.Time{}
}

var (
	// 1-based month, 0-based days, cumulative
	daytab = [][]int{
		{0, -1, 30, 58, 89, 119, 150, 180, 211, 242, 272, 303, 333, 364},
		{0, -1, 30, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334, 365},
	}
	days = []string{
		"sun", "mon", "tue", "wed", "thu", "fri", "sat",
	}
	months = []string{
		"jan", "feb", "mar", "apr", "may", "jun",
		"jul", "aug", "sep", "oct", "nov", "dec",
	}
)
