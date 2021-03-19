package easter

import "time"

func date(year, days int) time.Time {
	return time.Date(year, time.January, days, 0, 0, 0, 0, time.UTC)
}
