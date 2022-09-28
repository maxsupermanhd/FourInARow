package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

var (
	// begin crossplatform support

	scanner       = bufio.NewScanner(os.Stdin)
	iointerface   = "std"
	input         = make(chan string)
	input_batch   = []string{}
	output_batch  = []string{}
	output        = fmt.Print
	rand_override = []float32{}
	exitfunc      = os.Exit

	// end crossplatform support

	v_arr = []float64{1, 100, 500, 1e20, 1, 800, 4000, 1e20,
		1, 75, 900, 1e18, 1, 450, 3000, 1e18}
	b_str_arr                  = [8][8]string{}
	l_arr                      = [8]int64{}
	s_arr                      = [4]int64{}
	f_arr                      = [4]int64{}
	n_arr                      = [4]int64{}
	p_str, q_str, a_str, b_str string
	x_str                      = "X"
	o_str                      = "O"
	c, d, d1, d2, i, i1        int64
	j, k, l, l1, m, m4         int64
	m5, m9, n, n1, s, t        int64
	v, v1, w, z, z1            int64
)

func main() {
	bPRINT("                      FOUR IN A ROW\n")
	bPRINT("                    CREATIVE COMPUTING\n")
	bPRINT("                  MORRISTOWN, NEW JERSEY\n")
	bPRINT("\n")
	bPRINT("\n")
	bPRINT("\n")
	bPRINT("THE GAME OF FOUR IN A ROW\n")
	for {
		a_str = bINPUTs("DO YOU WANT INSTRUCTIONS? ")
		if a_str == "YES" {
			bPRINT("THE GAME CONSISTS OF STACKING X'S\n")
			bPRINT("AND O'S (THE COMPUTER HAS O) UNTIL\n")
			bPRINT("ONE OF THE PLAYERS GETS FOUR IN A\n")
			bPRINT("ROW VERTICALLY, HORIZONTALLY, OR \n")
			bPRINT("DIAGONALLY.\n")
			bPRINT("\n")
			bPRINT("\n")
			break
		} else if a_str == "NO" {
			break
		} else {
			bPRINT("YES OR NO\n")
		}
	}
	for i = 1; i <= 8; i++ {
		for j = 1; j <= 8; j++ {
			b_str_arr[i-1][j-1] = "-"
		}
	}
	for z1 = 1; z1 <= 8; z1++ {
		l_arr[z1-1] = 0
	}
	if bINPUTs("DO YOU WANT TO GO FIRST? ") == "NO" {
		goto l610
	}
	sub340()
l450:
	m = int64(bINPUTi("A NUMBER BETWEEN 1 AND 8? "))
	if m < 1 || m > 8 {
		bPRINT("ILLEGAL MOVE, TRY AGAIN.\n")
		goto l450
	}
	l = l_arr[m-1]
	if l > 7 {
		bPRINT("ILLEGAL MOVE, TRY AGAIN.\n")
		goto l450
	}
	l_arr[m-1] = l + 1
	l = l + 1
	b_str_arr[l-1][m-1] = "X"
	bPRINT("\n")
	sub340()
	p_str = x_str
	sub1240()
	for z = 1; z <= 4; z++ {
		if s_arr[z-1] >= 4 {
			bPRINT("Y O U   W I N !!!\n")
			exitfunc(0)
		}
	}
l610:
	m9 = 0
	v1 = 0
	n1 = 1
	for m4 = 1; m4 <= 8; m4++ {
		// bPRINT("TRACE ", c, d, d1, d2, i, i1, j, k, l, l1, m, m4, m5, m9, n, n1, s, t, v, v1, w, z, z\n1)
		// 631 PRINT"TRACE ";C;D;D1;D2;I;I1;J;K;L;L1;M;M4;M5;M9;N;N1;S;T;V;V1;W;Z;Z1
		l = l_arr[m4-1] + 1
		if l > 8 {
			continue
		}
		v = 1
		p_str = o_str
		w = 0
		m = m4
	l690:
		sub1240()
		for z1 = 1; z1 <= 4; z1++ {
			n_arr[z1-1] = 0
		}
		for z = 1; z <= 4; z++ {
			s = s_arr[z-1]
			if s-w > 3 {
				goto l1130
			}
			t = s + f_arr[z-1]
			if t >= 4 {
				v = v + 4
				n_arr[s-1] = n_arr[s-1] + 1
			}
		}
		for i = 1; i <= 4; i++ {
			n = n_arr[i-1] - 1
			if n != -1 {
				i1 = 8*w + 4*bSGNi(n) + i
				v = int64(float64(v) + v_arr[i1-1] + float64(n)*v_arr[8*w+i-1])
			}
		}
		if w != 1 {
			w = 1
			p_str = x_str
			goto l690
		}
		l = l + 1
		if l <= 8 {
			sub1240()
			for z = 1; z <= 4; z++ {
				if s_arr[z-1] > 3 {
					v = 2
				}
			}
		}
		if v < v1 {
			continue
		}
		if v > v1 {
			n1 = 1
			goto l1060
		}
		n1 = n1 + 1
		if bRND(1) > float32(1)/float32(n1) {
			continue
		}
	l1060:
		v1 = v
		m9 = m4
	}
	if m9 == 0 {
		bPRINT("T I E   G A M E ...\n")
		return
	}
	m = m9
l1130:
	bPRINT(fmt.Sprintf("COMPUTER PICKS COLUMN  %d \n\n", m))
	l = l_arr[m-1] + 1
	l_arr[m-1] = l_arr[m-1] + 1
	b_str_arr[l-1][m-1] = o_str
	p_str = o_str
	sub340()
	sub1240()
	for z = 1; z <= 4; z++ {
		if s_arr[z-1] >= 4 {
			bPRINT("C O M P U T E R   W I N S !!!\n")
			exitfunc(0)
		}
	}
	goto l450

}

// prints field
func sub340() {
	for i = 8; i >= 1; i-- {
		for j = 1; j <= 8; j++ {
			bPRINT(fmt.Sprintf("  %v", b_str_arr[i-1][j-1]))
		}
		bPRINT("\n")
	}
	// bPRINT("\n")
	for i = 1; i <= 8; i++ {
		bPRINT(fmt.Sprintf("  %d", i))
	}
	bPRINT(" \n")
	bPRINT("\n")
}

func sub1240() {
	q_str = x_str
	if p_str == x_str {
		q_str = o_str
	}
	d2 = 1
	d1 = 0
	z = 0
	sub1360()
	d2 = 1
	d1 = 1
	sub1360()
	d2 = 0
	d1 = 1
	sub1360()
	d2 = -1
	d1 = 1
	sub1360()
}

func sub1360() {
	d = 1
	s = 1
	t = 0
	z = z + 1
l1390:
	c = 0
	for k = 1; k <= 3; k++ {
		m5 = m + k*d1
		l1 = l + k*d2
		if m5 < 1 || l1 < 1 || m5 > 8 || l1 > 8 {
			continue
		}
		b_str = b_str_arr[l1-1][m5-1]
		if c == 0 {
			goto l1480
		}
	l1450:
		if b_str == q_str {
			k = 3
			goto l1510
		}
		t = t + 1
		goto l1510
	l1480:
		if b_str == p_str {
			s = s + 1
			goto l1510
		}
		c = 1
		goto l1450
	l1510:
		continue
	}
	if d != 0 {
		d = 0
		d1 = -d1
		d2 = -d2
		goto l1390
	}
	s_arr[z-1] = s
	f_arr[z-1] = t
}

// begin crossplatform support

func bINPUTs(q string) string {
	bPRINT(q)
	if iointerface == "std" {
		if !scanner.Scan() {
			if scanner.Err() != nil {
				exitfunc(0)
			}
		}
		return scanner.Text()
	}
	if iointerface == "js" {
		bPRINT("\n")
		return <-input
	}
	if iointerface == "tests" {
		if len(input_batch) > 0 {
			out := input_batch[0]
			if len(input_batch) > 1 {
				input_batch = input_batch[0:]
			} else {
				input_batch = []string{}
			}
			return out
		} else {
			bPRINT("No input!!!\n")
		}
	}
	return ""
}

func bINPUTi(q string) int {
	ret, err := strconv.Atoi(bINPUTs(q))
	if err != nil {
		return 0
	}
	return ret
}

func bPRINT(str string) {
	if iointerface != "tests" {
		output(str)
	} else {
		output_batch = append(output_batch, str)
	}
}

var rand_override_i = 0

func bRND(i int) float32 {
	if len(rand_override) > 0 {
		if len(rand_override) > rand_override_i {
			defer func() { rand_override_i++ }()
			return rand_override[rand_override_i]
		} else {
			rand_override_i = 0
			return rand_override[rand_override_i]
		}
	} else {
		return rand.Float32()
	}
}

// end crossplatform support

func bSGNi(value int64) int64 {
	if value == 0 {
		return 0
	} else if value > 0 {
		return 1
	}
	return -1
}
