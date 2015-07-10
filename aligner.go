package main

import (
	"bufio"
	"fmt"
	"io"
)

type Aligner struct {
	Delimiter *Delimiter
	Padding   *Padding
	Margin    *Margin
	Space     *Space
	lines     [][]string
}

func NewAligner() *Aligner {
	return &Aligner{
		Delimiter: NewDelimiter(),
		Padding:   NewPadding(),
		Margin:    NewMargin(),
		Space:     NewSpace(),
	}
}

func NewAlignerWithModules(d *Delimiter, p *Padding, m *Margin, s *Space) *Aligner {
	a := &Aligner{
		Delimiter: d,
		Padding:   p,
		Margin:    m,
		Space:     s,
	}
	if d == nil {
		a.Delimiter = NewDelimiter()
	}
	if p == nil {
		a.Padding = NewPadding()
	}
	if m == nil {
		a.Margin = NewMargin()
	}
	if s == nil {
		a.Space = NewSpace()
	}
	return a
}

func (a *Aligner) AppendLine(s string) {
	sp := a.Delimiter.Split(a.Space.Trim(s))
	a.lines = append(a.lines, sp)

	if len(sp) > 1 {
		a.Space.UpdateHeadWidth(s)
		a.Padding.UpdateWidth(sp)
	}
}

func (a *Aligner) ReadAll(r io.Reader) error {
	s := bufio.NewScanner(r)
	for s.Scan() {
		a.AppendLine(s.Text())
	}
	return s.Err()
}

func (a *Aligner) format(sp []string) string {
	if len(sp) == 1 {
		return sp[0]
	}
	return a.Space.Adjust(a.Margin.Join(a.Padding.Format(sp)))
}

func (a *Aligner) Flush(out io.Writer) error {
	for _, sp := range a.lines {
		_, err := fmt.Fprintln(out, a.format(sp))
		if err != nil {
			return err
		}
	}
	return nil
}
