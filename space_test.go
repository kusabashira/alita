package alita

import (
	"reflect"
	"testing"
)

func TestSpaceDefault(t *testing.T) {
	h := NewSpace()
	actual := h.tabWidth
	expect := 8
	if actual != expect {
		t.Errorf("got %d; want %d", actual, expect)
	}
}

type SpaceUpdateWidthTest struct {
	src    string
	before *Space
	after  *Space
}

var indexTestsSpaceUpdateWidth = []SpaceUpdateWidthTest{
	// read only headspace
	{"  abc ",
		&Space{8, 2, "  "},
		&Space{8, 2, "  "}},
	{"\t abc\t",
		&Space{8, 9, "\t"},
		&Space{8, 9, "\t"}},

	// update
	{"abc",
		&Space{8, 1, " "},
		&Space{8, 0, ""}},
	{"\tabc",
		&Space{8, 10, "\t  "},
		&Space{8, 8, "\t"}},
	{"\t abc",
		&Space{8, 16, "\t\t"},
		&Space{8, 9, "\t "}},

	// no update
	{" abc",
		&Space{8, 1, " "},
		&Space{8, 1, " "}},
	{"\t   abc",
		&Space{8, 8, "\t"},
		&Space{8, 8, "\t"}},
	{"\t\tabc",
		&Space{8, 9, "\t "},
		&Space{8, 9, "\t "}},
}

func TestSpaceUpdateWidth(t *testing.T) {
	for _, test := range indexTestsSpaceUpdateWidth {
		s := test.before
		s.UpdateHeadWidth(test.src)
		actual := s
		expect := test.after
		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("%v.UpdateWidth(%q) got %v; want %v",
				test.before, test.src, actual, expect)
		}
	}
}

type SpaceStrip struct {
	src string
	dst string
}

var indexTestsSpaceStrip = []SpaceStrip{
	{"abc", "abc"},

	{" abc", "abc"},
	{"   abc", "abc"},
	{"abc ", "abc"},
	{"abc  ", "abc"},
	{" abc ", "abc"},
	{"   abc   ", "abc"},

	{"\tabc", "abc"},
	{"\t abc\t", "abc"},
	{"\t\tabc\t", "abc"},
}

func TestSpaceStrip(t *testing.T) {
	s := NewSpace()
	for _, test := range indexTestsSpaceStrip {
		actual := s.Strip(test.src)
		expect := test.dst
		if actual != expect {
			t.Errorf("got %q; want %q", actual, expect)
		}
	}
}

type SpaceAdjust struct {
	space *Space
	src   string
	dst   string
}

var indexTestsSpaceAdjust = []SpaceAdjust{
	// no change
	{&Space{8, 0, ""}, "abc", "abc"},
	{&Space{8, 0, ""}, "  abc", "  abc"},

	// remove tail space
	{&Space{8, 0, ""}, "abc      ", "abc"},
	{&Space{8, 0, ""}, "abc\t    ", "abc"},

	// insert head space
	{&Space{8, 1, " "}, "abc", " abc"},
	{&Space{8, 8, "\t"}, "abc", "\tabc"},
	{&Space{8, 9, "\t "}, "abc", "\t abc"},

	// both change
	{&Space{8, 9, "\t "}, "abc   ", "\t abc"},
	{&Space{8, 9, "\t "}, "abc\t ", "\t abc"},
	{&Space{8, 9, "\t "}, "  abc\t ", "\t   abc"},
	{&Space{8, 9, "\t "}, "\tabc\t ", "\t \tabc"},
}

func TestSpaceAdjust(t *testing.T) {
	for _, test := range indexTestsSpaceAdjust {
		s := test.space
		actual := s.Adjust(test.src)
		expect := test.dst
		if actual != expect {
			t.Error("%v.Adjust(%q) = %q; want %q",
				s, test.src, actual, expect)
		}
	}
}