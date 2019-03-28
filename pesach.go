package calendar

import "time"

// Pesach date of Pesach using the Gauss formula
func Pesach(year int) time.Time {
	return date(year, pesach(year))
}

func pesach(year int) int {
	var a, b, y, cumdays int
	var d float64

	const (
		t = 33.0 + 14.0/24.0
		l = (1.0 + 485.0/1080.0) / 24.0 / 19.0
		k = (29.0 + (12.0+793.0/1080.0)/24.0) / 19.0
	)

	y = year + 3760

	a = (12*y + 17) % 19
	b = y % 4
	d = (t - 10.0*k + l + 14.0) + k*float64(a) + float64(b)/4.0 - l*float64(y)
	cumdays = int(d)

	/* the postponement */
	switch (cumdays + 3*y + 5*b + 5) % 7 {
	case 1:
		if a > 6 && d-float64(cumdays) >= (15.0+204.0/1080.0)/24.0 {
			cumdays += 2
		}
	case 0, 2, 4, 6:
		if a > 11 || d-float64(cumdays) >= (21.0+589.0/1080.0)/24.0 {
			cumdays++
		}
	}

	if year > 1582 {
		cumdays += year/100 - year/400 - 2
	}

	cumdays += 31 + 28
	if isleap(year) {
		cumdays += 1
	}
	return cumdays
}
