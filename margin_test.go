package alita

import (
	"testing"
)

func TestDefaultMargin(t *testing.T) {
	m := NewMargin()
	l, r := 1, 1
	if m.left != l || m.right != r {
		t.Errorf("got %d:%d, want %d:%d", m.left, m.right, l, r)
	}
}

type FormatTest struct {
	format string
	left   int
	right  int
}

var indexTestsFormat = []FormatTest{
	// digit only
	{"1", 1, 1},
	{"2", 2, 2},
	{"10", 10, 10},
	{"00", 0, 0},

	// colon separated digits
	{"1:1", 1, 1},
	{"2:1", 2, 1},
	{"1:4", 1, 4},
	{"10:5", 10, 5},
	{"5:20", 5, 20},
}

func TestParseFormat(t *testing.T) {
	m := NewMargin()
	for _, test := range indexTestsFormat {
		if err := m.Set(test.format); err != nil {
			t.Errorf("Set(%q) returns err; want nil", test.format)
		}
		if m.left != test.left || m.right != test.right {
			t.Errorf("got %d:%d, want %d:%d",
				m.left, m.right, test.left, test.right)
		}
	}
}

var indexTestsErrFormat = []string{
	"abc",
	"100000000000000000000000000000",
	"-1",
	":",
	"1:",
	":1",
	"1:-1",
	"-1:1",
	"3:1:4",
}

func TestParseErrFormat(t *testing.T) {
	m := NewMargin()
	for _, format := range indexTestsErrFormat {
		if err := m.Set(format); err == nil {
			t.Errorf("Set(%q) returns nil; want err", format)
		}
	}
}

type JoinTest struct {
	format string
	src    []string
	dst    string
}

var indexTestJoinStrings = []JoinTest{
	{"1", []string{"n", "=", "100"}, "n = 100"},
	{"2", []string{"n", "=", "100"}, "n  =  100"},
}

func TestJoinStrings(t *testing.T) {
	m := NewMargin()
	for _, test := range indexTestJoinStrings {
		if err := m.Set(test.format); err != nil {
			t.Errorf("Set(%q) returns err; want nil", test.format)
		}
		actual := m.Join(test.src)
		expect := test.dst
		if actual != expect {
			t.Errorf("Join(%q) = %q; want %q",
				test.src, actual, expect)
		}
	}
}