package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	digitOnly            = regexp.MustCompile(`^\d+$`)
	colonSeparatedDigits = regexp.MustCompile(`^(\d+):(\d+)$`)
)

type Margin struct {
	left  int
	right int
}

func NewMargin(format string) (*Margin, error) {
	m := &Margin{}
	switch {
	case format == "":
		m.left, m.right = 1, 1
	case digitOnly.MatchString(format):
		n, err := strconv.Atoi(format)
		if err != nil {
			return nil, err
		}
		m.left, m.right = n, n
	case colonSeparatedDigits.MatchString(format):
		a := colonSeparatedDigits.FindAllStringSubmatch(format, -1)[0]
		l, err := strconv.Atoi(a[1])
		if err != nil {
			return nil, err
		}
		r, err := strconv.Atoi(a[2])
		if err != nil {
			return nil, err
		}
		m.left, m.right = l, r
	default:
		return nil, fmt.Errorf("margin: invalid format: %s", format)
	}
	return m, nil
}

func NewMarginWithNumber(left, right int) *Margin {
	return &Margin{
		left:  left,
		right: right,
	}
}

func (m *Margin) Join(a []string) string {
	switch len(a) {
	case 0:
		return ""
	case 1:
		return a[0]
	}

	l, r := m.left, m.right
	if l < 0 {
		l = 0
	}
	if r < 0 {
		r = 0
	}

	buflen := (l + r) * (len(a) / 2)
	if len(a)%2 == 0 {
		buflen -= r
	}
	for i := 0; i < len(a); i++ {
		buflen += len(a[i])
	}
	lm, rm := strings.Repeat(" ", l), strings.Repeat(" ", r)

	b := make([]byte, buflen)
	bp := copy(b, a[0])
	for i := 2; i <= len(a); i += 2 {
		bp += copy(b[bp:], lm)
		bp += copy(b[bp:], a[i-1])
		if i != len(a) {
			bp += copy(b[bp:], rm)
			bp += copy(b[bp:], a[i])
		}
	}
	return string(b)
}
