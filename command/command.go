package command

import (
	"errors"
	"fmt"
	"github.com/MegaShow/goselpg/print"
	"github.com/spf13/pflag"
	"os"
)

// define variables
var d struct {
	c print.Settings // const
	isHelp  bool
	isParse bool
	args	[]string
	visit   map[string]bool
}

func init() {
	d.c = print.Settings{
		StartPage:     1,
		EndPage:       0,
		InputFileName: "",
		LinesPerPage:  72,
		FormFeed: false,
		Destination:   "",
	}
	pflag.Usage = usage
}

func Parse() (s print.Settings, err error) {
	if d.isParse {
		return s, errors.New("has already parsed")
	}
	d.isParse = true

	// unexported settings
	pflag.BoolVarP(&d.isHelp, "help", "h", false, "help")

	// exported settings
	pflag.IntVarP(&s.StartPage, "start", "s", d.c.StartPage, "start page number")
	pflag.IntVarP(&s.EndPage, "end", "e", d.c.EndPage, "end page number (default max-page)")
	pflag.IntVarP(&s.LinesPerPage, "lines", "l", d.c.LinesPerPage, "lines per page")
	pflag.BoolVarP(&s.FormFeed, "formFeed", "f", d.c.FormFeed, "fixed lines per page")
	pflag.StringVarP(&s.Destination, "dest", "d", d.c.Destination, "destination")

	pflag.Parse()
	d.args = pflag.Args()
	d.visit = make(map[string]bool)
	pflag.Visit(func(f *pflag.Flag) { d.visit[f.Name] = true })
	if len(d.args) >= 1 {
		s.InputFileName = d.args[0]
	}
	if d.isHelp {
		pflag.Usage()
		os.Exit(1)
	}
	return s, check(s)
}

func usage() {
	fmt.Printf("selpg version: selpg/0.0.1\n" +
		"Usage: selpg [-s startPage] [-e endPage] [-l linesPerPage | -f] filename\n\n" +
		"Options:\n")
	pflag.PrintDefaults()
	fmt.Println()
}
