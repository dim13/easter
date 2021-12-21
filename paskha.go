package easter

import "time"

// Paskha date for Orthodox Easter using Gauss formula
func Paskha(year int) time.Time {
	return date(year, paskha(year))
}

func paskha(year int) int {
	a := year % 19
	b := year % 4
	c := year % 7
	d := (19*a + 15) % 30
	e := (2*b + 4*c + 6*d + 6) % 7
	cumdays := 31 + 28
	if isLeap(year) {
		cumdays++
	}
	return (cumdays + 22) + (d + e) + 13
}
