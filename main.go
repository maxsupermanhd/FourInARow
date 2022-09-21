package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	scanner = bufio.NewScanner(os.Stdin)

	v_arr = []float64{1, 100, 500, 1e20, 1, 800, 4000, 1e20,
		1, 75, 900, 1e18, 1, 450, 3000, 1e18}
	b_str_arr                                   = [8][8]string{}
	l_arr                                       = [8]int{}
	s_arr                                       = [4]int{}
	f_arr                                       = [4]int{}
	n_arr                                       = [4]int{}
	p_str, q_str                                string
	x_str                                       = "X"
	o_str                                       = "O"
	l, m, z, d1, d2, d, s, t, c, m9, v1, n1, m4 int
)

func main() {
	fmt.Println("       FOUR IN A ROW")
	fmt.Println("     CREATIVE COMPUTING")
	fmt.Println("   MORRISTOWN, NEW JERSEY")
	fmt.Println()
	fmt.Println()
	fmt.Println("THE GAME OF FOUR IN A ROW")
	for {
		a := bINPUTs("DO YOU WANT INSTRUCTIONS ")
		if a == "YES" {
			fmt.Println("THE GAME CONSISTS OF STACKING X'S")
			fmt.Println("AND O'S (THE COMPUTER HAS O) UNTIL")
			fmt.Println("ONE OF THE PLAYERS GETS FOUR IN A")
			fmt.Println("ROW VERTICALLY, HORIZONTALLY, OR ")
			fmt.Println("DIAGONALLY.")
			fmt.Println()
			fmt.Println()
			break
		} else if a == "NO" {
			break
		} else {
			fmt.Println("YES OR NO")
		}
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			b_str_arr[i][j] = "-"
		}
	}
	for z1 := 0; z1 < 8; z1++ {
		l_arr[z1] = 0
	}
	if bINPUTs("DO YOU WANT TO GO FIRST ") == "NO" {
		// 610
	}
	sub340()

	for {
		m = bINPUTi("A NUMBER BETWEEN 1 AND 8")
		if m < 1 || m > 8 || l_arr[m-1] > 7 {
			fmt.Println("ILLEGAL MOVE, TRY AGAIN.")
			continue
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
			m = m4
			sub1240()
			for z1 = 1; z1 <= 4; z1++ {
				n_arr[z1-1] = 0
			}
			for z = 1; z <= 4; z++ {
				s = s_arr[z-1]
				if s-w > 3 {
					// 1130
				}
			}
		}
	}

}

// prints field
func sub340() {
	for i := 8; i >= 1; i-- {
		for j := 1; j <= 8; j++ {
			fmt.Printf("  %v", b_str_arr[i-1][j-1])
		}
		fmt.Println()
	}
	fmt.Println()
	for i := 1; i <= 8; i++ {
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
	c := 0
	for k := 1; k <= 3; k++ {
		m5 := m + k*d1
		l1 := l_arr[m-1] - 1 + k*d2
		if m5 < 1 || l1 < 1 || m5 > 8 || l1 > 8 {
			continue
		}
		sb := b_str_arr[l1-1][m5-1]
		if c == 0 {
			if sb == p_str {
				s++
				continue
			}
			c = 1
		}
		if sb == q_str {
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
