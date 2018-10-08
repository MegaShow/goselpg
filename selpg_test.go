package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestSelpg(t *testing.T) {
	t.Log("Execute: go build selpg.go")
	if err := exec.Command("go", "build", "selpg.go").Run(); err != nil {
		t.Error(err)
		return
	}
	t.Log("Create: input.txt")
	createInputFile()

	t.Log("Execute: selpg -s1 -e1 input.txt")
	if data, err := exec.Command("./selpg", "-s1", "-e1", "input.txt").Output(); err != nil {
		t.Error(err)
		return
	} else if !compareWithCmd1(data) {
		t.Error("=== Output fail ===")
		return
	}

	t.Log("Execute: selpg -s1 -e1 < input.txt")
	cmd := exec.Command("./selpg", "-s1", "-e1")
	cmd.Stdin, _ = os.Open("input.txt")
	if data, err := cmd.Output(); err != nil {
		t.Error(err)
		return
	} else if !compareWithCmd1(data) {
		t.Error("=== Output fail ===")
		return
	}
}

func createInputFile() {
	f, err := os.Create("input.txt")
	defer f.Close()
	if err != nil {

	}
	for i := 0; i < 333; i++ {
		f.WriteString(fmt.Sprintf("lines %d\n", i*3+1))
		f.WriteString(fmt.Sprintf("lines %d\n", i*3+2))
		f.WriteString(fmt.Sprintf("lines %d\n\f", i*3+3))
	}
}

// selpg -s1 -e1 input.txt
// selpg -s1 -e1 < input.txt
func compareWithCmd1(data []byte) bool {
	var a, e strings.Builder
	a.Write(data)
	for i := 0; i < 24; i++ {
		e.WriteString(fmt.Sprintf("lines %d\n", i*3+1))
		e.WriteString(fmt.Sprintf("lines %d\n", i*3+2))
		e.WriteString(fmt.Sprintf("lines %d\n", i*3+3))
		if i != 23 {
			e.WriteByte('\f')
		}
	}
	return a.String() == e.String()
}
