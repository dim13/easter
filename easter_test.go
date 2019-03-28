package calendar

import (
	"testing"
	"time"
)

/*

 */

func TestEaster(t *testing.T) {
	testCases := []struct {
		year  int
		month time.Month
		day   int
	}{
		{1999, time.April, 4},
		{2000, time.April, 23},
		{2001, time.April, 15},
		{2002, time.March, 31},
		{2003, time.April, 20},
		{2004, time.April, 11},
		{2005, time.March, 27},
		{2006, time.April, 16},
		{2007, time.April, 8},
		{2008, time.March, 23},
		{2009, time.April, 12},
		{2010, time.April, 4},
		{2011, time.April, 24},
		{2012, time.April, 8},
		{2013, time.March, 31},
		{2014, time.April, 20},
		{2015, time.April, 5},
		{2016, time.March, 27},
		{2017, time.April, 16},
		{2018, time.April, 1},
		{2019, time.April, 21},
		{2020, time.April, 12},
		{2021, time.April, 4},
		{2022, time.April, 17},
		{2023, time.April, 9},
		{2024, time.March, 31},
		{2025, time.April, 20},
		{2026, time.April, 5},
		{2027, time.March, 28},
		{2028, time.April, 16},
		{2029, time.April, 1},
		{2030, time.April, 21},
		{2031, time.April, 13},
		{2032, time.March, 28},
		{2033, time.April, 17},
		{2034, time.April, 9},
		{2035, time.March, 25},
		{2036, time.April, 13},
		{2037, time.April, 5},
		{2038, time.April, 25},
		{2039, time.April, 10},
	}
	for _, tc := range testCases {
		got := Easter(tc.year)
		z := time.Date(tc.year, 1, got, 0, 0, 0, 0, time.UTC)
		if z.Day() != tc.day || z.Month() != tc.month {
			t.Logf("got %v, want %v", z.Format("02 Jan 2006"), tc)
		}
	}
}
