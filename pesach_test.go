package calendar

import (
	"testing"
	"time"
)

func TestPesach(t *testing.T) {
	testCases := []struct {
		year  int
		month time.Month
		day   int
	}{
		{1999, time.April, 1},
		{2000, time.April, 20},
		{2001, time.April, 8},
		{2002, time.March, 28},
		{2003, time.April, 17},
		{2004, time.April, 5},
		{2005, time.April, 24},
		{2006, time.April, 13},
		{2007, time.April, 2},
		{2008, time.April, 20},
		{2009, time.April, 9},
		{2010, time.March, 30},
		{2011, time.April, 19},
		{2012, time.April, 6},
		{2013, time.March, 26},
		{2014, time.April, 15},
		{2015, time.April, 4},
		{2016, time.April, 23},
		{2017, time.April, 11},
		{2018, time.March, 31},
		{2019, time.April, 20},
		{2020, time.April, 8},
		{2021, time.March, 28},
		{2022, time.April, 17},
		{2023, time.April, 5},
		{2024, time.April, 23},
		{2025, time.April, 13},
		{2026, time.April, 2},
		{2027, time.April, 22},
		{2028, time.April, 11},
		{2029, time.March, 30},
		{2030, time.April, 18},
		{2031, time.April, 7},
		{2032, time.March, 26},
		{2033, time.April, 14},
		{2034, time.April, 4},
		{2035, time.April, 24},
		{2036, time.April, 11},
		{2037, time.March, 31},
		{2038, time.April, 20},
		{2039, time.April, 8},
	}
	for _, tc := range testCases {
		got := Pesach(tc.year)
		if got.Day() != tc.day || got.Month() != tc.month {
			t.Logf("got %v, want %v", got.Format("02 Jan 2006"), tc)
		}
	}
}
