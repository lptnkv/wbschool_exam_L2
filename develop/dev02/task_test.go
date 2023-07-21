package main

import (
	"testing"
)

type test struct {
	id     int
	input  string
	want   string
	hasErr bool
}

func TestRleDecode(t *testing.T) {
	testCases := []test{
		{
			1,
			"abcd",
			"abcd",
			false,
		},
		{
			2,
			"",
			"",
			false,
		},
		{
			3,
			"45",
			"",
			true,
		},
		{
			4,
			"a5",
			"aaaaa",
			false,
		},
		{
			5,
			"a5b3ch1",
			"aaaaabbbch",
			false,
		},
		{
			6,
			"g20",
			"gggggggggggggggggggg",
			false,
		},
	}
	for _, testCase := range testCases {
		got, err := rleDecode(testCase.input)
		if err == nil && testCase.hasErr {
			t.Errorf("%d: expected error, but got none\n", testCase.id)
		}
		if got != testCase.want {
			t.Errorf("%d: want %s, but got %s\n", testCase.id, testCase.want, got)
		}
	}
}
