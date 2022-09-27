package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"testing"
)

func GenInputs() []string {
	out := []string{}
	if rand.Int()%2 == 0 {
		out = append(out, "YES")
	} else {
		out = append(out, "NO")
	}
	if rand.Int()%2 == 0 {
		out = append(out, "YES")
	} else {
		out = append(out, "NO")
	}
	// we need to make 4*8 moves
	// extra will be not consumed
	for i := 0; i < 4*8+2; i++ {
		out = append(out, strconv.Itoa(rand.Int()%8+1))
	}
	return out
}

type StringArrReader struct {
	strs []string
}

func NewStringArrReader(s []string) *StringArrReader {
	r := StringArrReader{}
	r.strs = make([]string, len(s))
	copy(r.strs, s)
	return &r
}

func (r *StringArrReader) Read(o []byte) (n int, err error) {
	if len(r.strs) > 0 {
		out := r.strs[0]
		if len(r.strs) > 1 {
			r.strs = r.strs[0:]
		} else {
			r.strs = []string{}
		}
		copy(o, out)
		return len(o), nil
	} else {
		return 0, io.EOF
	}
}

func TestMain(m *testing.M) {
	fmt.Println("Testing main...")
	inputs := GenInputs()
	input_batch = make([]string, len(inputs))
	copy(input_batch, inputs)
	m.Run()
	targetfile, err := os.ReadFile("target.bas")
	if err != nil {
		panic(err)
	}
	testfile := "1 RNG(-" + strconv.Itoa(rand.Int()%31+1) + ")\n" + string(targetfile)
	err = os.WriteFile("testing.bas", []byte(testfile), 0644)
	if err != nil {
		panic(err)
	}
	cmd := exec.Command("./vintbas", "testing.bas")
	cmd.Stdin = NewStringArrReader(inputs)
	os.Exit(0)
}
