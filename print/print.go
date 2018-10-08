package print

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

type Settings struct {
	StartPage     int
	EndPage       int
	InputFileName string
	LinesPerPage  int
	FormFeed      bool
	Destination   string
}

type printer struct {
	Settings
	reader *bufio.Reader
	data []byte
}

func Execute(s Settings) error {
	var p = printer{ Settings: s }
	if err := p.scan(); err != io.EOF && err != nil {
		return err
	}
	return p.print()
}

func (p *printer) scan() error {
	if p.InputFileName == "" {
		p.reader = bufio.NewReader(os.Stdin)
	} else {
		f, err := os.Open(p.InputFileName)
		defer f.Close()
		if err != nil {
			return err
		}
		p.reader = bufio.NewReader(f)
	}
	for page := 1; page < p.StartPage; page++ {
		if p.FormFeed {
			if _, err :=  p.reader.ReadSlice('\f'); err != nil {
				return err
			}
		} else {
			for i := 0; i < p.LinesPerPage; i++ {
				if _, err := p.reader.ReadSlice('\n'); err != nil {
					return err
				}
			}
		}
	}
	for page := p.StartPage; p.EndPage == 0 || page <= p.EndPage; page++ {
		if p.FormFeed {
			str, err := p.reader.ReadSlice('\f')
			p.data = append(p.data, str...)
			if err != nil {
				return err
			}
		} else {
			for i := 0; i < p.LinesPerPage; i++ {
				str, err := p.reader.ReadSlice('\n')
				p.data = append(p.data, str...)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (p *printer) print() error {
	if p.Destination == "" {
		fmt.Fprint(os.Stdout, string(p.data))
	} else {
		cmd := exec.Command("lp", "-d"+p.Destination)
		stdin, err := cmd.StdinPipe()
		if err != nil {
			return err
		}
		go func() {
			defer stdin.Close()
			fmt.Fprint(stdin, string(p.data))
		}()
		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", out)
	}
	return nil
}
