package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)

	v_arr = []float64{1, 100, 500, 1e20, 1, 800, 4000, 1e20,
		1, 75, 900, 1e18, 1, 450, 3000, 1e18}
	b_str_arr                  = [8][8]string{}
	l_arr                      = [8]int{}
	s_arr                      = [4]int{}
	f_arr                      = [4]int{}
	n_arr                      = [4]int{}
	p_str, q_str, a_str, b_str string
	x_str                      = "X"
	o_str                      = "O"
	l, m, z, d1, d2, d         int
	s, t, c, m9, v1, n1        int
	m4, z1, v, w, i, n         int
	i1, k, m5, l1, j           int
)

func main() {
	fmt.Println("       FOUR IN A ROW")
	fmt.Println("     CREATIVE COMPUTING")
	fmt.Println("   MORRISTOWN, NEW JERSEY")
	fmt.Println()
	fmt.Println()
	fmt.Println("THE GAME OF FOUR IN A ROW")
	for {
		a_str = bINPUTs("DO YOU WANT INSTRUCTIONS ")
		if a_str == "YES" {
			fmt.Println("THE GAME CONSISTS OF STACKING X'S")
			fmt.Println("AND O'S (THE COMPUTER HAS O) UNTIL")
			fmt.Println("ONE OF THE PLAYERS GETS FOUR IN A")
			fmt.Println("ROW VERTICALLY, HORIZONTALLY, OR ")
			fmt.Println("DIAGONALLY.")
			fmt.Println()
			fmt.Println()
			break
		} else if a_str == "NO" {
			break
		} else {
			fmt.Println("YES OR NO")
		}
	}
	for i = 0; i < 8; i++ {
		for j = 0; j < 8; j++ {
			b_str_arr[i][j] = "-"
		}
	}
	for z1 = 0; z1 < 8; z1++ {
		l_arr[z1] = 0
	}
	if bINPUTs("DO YOU WANT TO GO FIRST ") == "NO" {
		goto l610
	}
	sub340()

l450:
	m = bINPUTi("A NUMBER BETWEEN 1 AND 8 ")
	if m < 1 || m > 8 || l_arr[m-1] > 7 {
		fmt.Println("ILLEGAL MOVE, TRY AGAIN.")
		goto l450
	}
	l_arr[m-1]++
	l = l_arr[m-1]
	b_str_arr[l_arr[m-1]-1][m-1] = "X"
	fmt.Println()
	sub340()
	p_str = x_str
	sub1240()
	for z = 1; z <= 4; z++ {
		if s_arr[z-1] < 4 {
			continue
		}
		fmt.Print("Y O U   W I N !!!")
		return
	}
l610:
	m9 = 0
	v1 = 0
	n1 = 1
	for m4 = 1; m4 <= 8; m4++ {
		l = l_arr[m4-1] + 1
		if l > 8 {
			continue
		}
		v = 1
		p_str = o_str
		w = 0
	l690:
		m = m4
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
				v = v + int(v_arr[i1-1]) + n*int(v_arr[8*w+i-1])
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
		if v == v1 {
			n1 = n1 + 1
			if rand.Float32() > float32(1)/float32(n1) {
				continue
			}
			v1 = v
			m9 = m4
		} else {
			n1 = 1
		}
	}
	if m9 == 0 {
		fmt.Println("T I E   G A M E ...")
		return
	}
	m = m9
l1130:
	fmt.Println("COMPUTER PICKS COLUMN ", m)
	l = l_arr[m-1] + 1
	l_arr[m-1] = l_arr[m-1] + 1
	b_str_arr[l-1][m-1] = o_str
	p_str = o_str
	sub340()
	sub1240()
	for z = 1; z <= 4; z++ {
		if s_arr[z-1] >= 4 {
			fmt.Println("C O M P U T E R   W I N S !!!")
			return
		}
	}
	goto l450

}

// prints field
func sub340() {
	for i = 8; i >= 1; i-- {
		for j = 1; j <= 8; j++ {
			fmt.Printf("  %v", b_str_arr[i-1][j-1])
		}
		fmt.Println()
	}
	fmt.Println()
	for i = 1; i <= 8; i++ {
		fmt.Printf("  %d", i)
	}
	fmt.Println()
	fmt.Println()
}

func sub1240() {
	q_str = "X"
	if p_str == "X" {
		q_str = "O"
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
		l1 = l_arr[m-1] - 1 + k*d2
		if m5 < 1 || l1 < 1 || m5 > 8 || l1 > 8 {
			continue
		}
		b_str = b_str_arr[l1-1][m5-1]
		if c == 0 {
			if b_str == p_str {
				s++
				continue
			}
			c = 1
		}
		if b_str == q_str {
			k = 3
		} else {
			t++
		}
		continue
	}
	if d != 0 {
		d = 0
		d1 = -d1
		d2 = -d2
		goto l1390 // I hate this
	}
	s_arr[z-1] = s
	f_arr[z-1] = t
}

func bINPUTs(q string) string {
	fmt.Print(q)
	if !scanner.Scan() {
		if scanner.Err() != nil {
			panic(scanner.Err())
		}
	}
	return scanner.Text()
}

func bINPUTi(q string) int {
	ret, err := strconv.Atoi(bINPUTs(q))
	if err != nil {
		os.Exit(1)
	}
	return ret
}

func bSGNi(v int) int {
	if v == 0 {
		return 0
	} else if v > 0 {
		return 1
	}
	return -1
}
