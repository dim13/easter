package calendar

// return year day for Orthodox Easter using Gauss formula
// (new style result); subtract 13 for old style

func Paskha(year int) int {
	var a, b, c, d, e int
	const (
		x = 15
		y = 6
	)
	var cumdays int

	a = year % 19
	b = year % 4
	c = year % 7
	d = (19*a + x) % 30
	e = (2*b + 4*c + 6*d + y) % 7
	cumdays = 31 + 28
	if isleap(year) {
		cumdays++
	}
	return (cumdays + 22) + (d + e) + 13
}
