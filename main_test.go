package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
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

func TestMain(m *testing.M) {
	vintbasbin := os.Getenv("VINTBAS_PATH")
	if vintbasbin == "" {
		vintbasbin = "./vintbas"
	}
	fmt.Printf("Comparing against vintbas at [%s]\n", vintbasbin)

	seed, err := strconv.Atoi(os.Getenv("RAND_SEED"))
	if err != nil {
		seed = -42
	}

	rand.Seed(int64(seed))

	numbercount := 8096
	fmt.Printf("Testing seed %d...\n", seed)

	fmt.Printf("Generating %d random numbers...\n", numbercount)
	testfile := "1 PRINT RND(" + strconv.Itoa(seed) + ")\n"
	for i := 0; i < numbercount; i++ {
		testfile += fmt.Sprintf("%d PRINT RND(1)\n", i+2)
	}
	// we have to create a file, it can not read from stdin
	must(os.WriteFile("testing.bas", []byte(testfile), 0644))
	fmt.Println("Running vintbas to generate numbers...")
	cmd := exec.Command(vintbasbin, "testing.bas")
	fmt.Println("Getting output from vintbas...")
	rng := noerr(cmd.Output())
	fmt.Println("Dividing vintbas output...")
	rngs := strings.Split(string(rng), "\n")
	fmt.Println("Parsing numbers...")
	for i, r := range rngs {
		if i == 0 {
			// first rng call omitted
			continue
		}
		if r == "" || r == " " {
			continue
		}
		add := float32(0.0)
		n := noerr(fmt.Sscanf(r, "%f", &add))
		if n == 1 {
			rand_override = append(rand_override, add)
		}
	}

	fmt.Println("Getting og program output with set rng...")
	input_batch = GenInputs()
	fmt.Printf("Inputs are: %#v\n", input_batch)
	targetfile := noerr(os.ReadFile("target.bas"))
	testfile = "1 XXXXXXXXXXXX=RND(" + strconv.Itoa(seed) + ")\n" + string(targetfile)
	must(os.WriteFile("testing.bas", []byte(testfile), 0644))
	cmd = exec.Command(vintbasbin, "testing.bas")
	in := noerr(cmd.StdinPipe())
	for _, inp := range input_batch {
		in.Write([]byte(inp + "\n"))
	}
	ideal_output := string(noerr(cmd.Output()))

	fmt.Println("Running actual program...")
	iointerface = "tests"
	exitfunc = func(code int) {
		fmt.Println("Comparing outputs...")
		actual_output := strings.Join(output_batch, "")

		ideal_lined := strings.Split(ideal_output, "\n")
		actual_lined := strings.Split(actual_output, "\n")

		if len(ideal_lined) != len(actual_lined) {
			defer os.Exit(1)
			fmt.Printf("Not equal line count: have %d expected %d\n", len(actual_lined), len(ideal_lined))
		} else {
			fmt.Printf("Comparing %d lines...\n", len(ideal_lined))
		}

		maxlen := 0
		for i := 0; i < len(ideal_lined) && i < len(actual_lined); i++ {
			if len(actual_lined[i]) > maxlen {
				maxlen = len(actual_lined[i])
			}
			if len(ideal_lined[i]) > maxlen {
				maxlen = len(actual_lined[i])
			}
		}
		matching := 0
		total := int(math.Max(float64(len(ideal_lined)), float64(len(actual_lined))))
		for i := 0; i < len(ideal_lined) && i < len(actual_lined); i++ {
			fmt.Printf("%- *s | %- *s\n", maxlen, actual_lined[i], maxlen, ideal_lined[i])
			if ideal_lined[i] != actual_lined[i] {
				fmt.Printf("Line %- 8d missmatch: has [%s] want [%s]\n", i, actual_lined[i], ideal_lined[i])
				defer os.Exit(1)
			} else {
				matching++
			}
		}

		fmt.Printf("Matching %d out of %d total.\n", matching, total)
		if matching == total {
			defer os.Exit(0)
		} else {
			defer os.Exit(1)
		}
	}
	main()

}

func noerr[T any](ret T, err error) T {
	if err != nil {
		panic(err)
	}
	return ret
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
