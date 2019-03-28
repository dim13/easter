package calendar

import (
	"testing"
	"time"
)

func TestPaskha(t *testing.T) {
	testCases := []struct {
		year  int
		month time.Month
		day   int
	}{
		{1999, time.April, 11},
		{2000, time.April, 30},
		{2001, time.April, 15},
		{2002, time.May, 5},
		{2003, time.April, 27},
		{2004, time.April, 11},
		{2005, time.May, 1},
		{2006, time.April, 23},
		{2007, time.April, 8},
		{2008, time.April, 27},
		{2009, time.April, 19},
		{2010, time.April, 4},
		{2011, time.April, 24},
		{2012, time.April, 15},
		{2013, time.May, 5},
		{2014, time.April, 20},
		{2015, time.April, 12},
		{2016, time.May, 1},
		{2017, time.April, 16},
		{2018, time.April, 8},
		{2019, time.April, 28},
		{2020, time.April, 19},
		{2021, time.May, 2},
		{2022, time.April, 24},
		{2023, time.April, 16},
		{2024, time.May, 5},
		{2025, time.April, 20},
		{2026, time.April, 12},
		{2027, time.May, 2},
		{2028, time.April, 16},
		{2029, time.April, 8},
		{2030, time.April, 28},
		{2031, time.April, 13},
		{2032, time.May, 2},
		{2033, time.April, 24},
		{2034, time.April, 9},
		{2035, time.April, 29},
		{2036, time.April, 20},
		{2037, time.April, 5},
		{2038, time.April, 25},
		{2039, time.April, 17},
	}
	for _, tc := range testCases {
		got := Paskha(tc.year)
		if got.Day() != tc.day || got.Month() != tc.month {
			t.Logf("got %v, want %v", got.Format("02 Jan 2006"), tc)
		}
	}
}
