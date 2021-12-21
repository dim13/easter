package easter

import "time"

// Easter date
func Easter(year int) time.Time {
	return date(year, easter(year))
}

func easter(year int) int {
	a := year % 19
	b := year / 100
	c := year % 100
	d := b / 4
	e := b % 4
	f := (b + 8) / 25
	g := (b + 1 - f) / 3
	h := ((19 * a) + 15 + b - (d + g)) % 30
	i := c / 4
	k := c % 4
	l := (32 + 2*e + 2*i - (h + k)) % 7
	m := (a + 11*h + 22*l) / 451
	n := (h + l + 114 - (7 * m)) / 31
	p := (h + l + 114 - (7 * m)) % 31
	q := 31 + 28 + p + 1
	if isLeap(year) {
		q++
	}
	if n == 4 {
		q += 31
	}
	return q
}
