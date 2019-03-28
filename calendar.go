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
