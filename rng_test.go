package main

import (
	"fmt"
	"testing"
)

func TestRng(t *testing.T) {
	rndSetSeed(0x8000000000000000)
	fmt.Println(bRND(1))
	fmt.Println(bRND(1))
	fmt.Println(bRND(1))
	fmt.Println(bRND(1))
}
