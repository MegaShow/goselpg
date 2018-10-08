package command

import (
	"errors"
	"github.com/MegaShow/goselpg/print"
)

func check(s print.Settings) (err error) {
	if d.visit["start"] && isPositiveNumber(s.StartPage) == false {
		return errors.New("start page number cannot be non-positive")
	}
	if d.visit["end"] && isPositiveNumber(s.EndPage) == false {
		return errors.New("end page number cannot be non-positive")
	}
	if d.visit["start"] && d.visit["end"] && s.StartPage > s.EndPage {
		return errors.New("start page number cannot be larger that end page number")
	}
	if d.visit["lines"] && isPositiveNumber(s.LinesPerPage) == false {
		return errors.New("lines cannot be non-positive")
	}
	if d.visit["lines"] && d.visit["formFeed"] {
		return errors.New("lines and formFeed cannot be set again")
	}
	if len(d.args) > 1 {
		return errors.New("more than one file defined")
	}
	return
}

func isPositiveNumber(value int) (bool) {
	return value > 0
}
