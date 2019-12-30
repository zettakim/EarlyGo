package main

import "fmt"

func main() {
	a := make([]int, 5, 10)

	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println()

	b := []int{1, 2, 3}
	b = append(b, 4, 5, 6)
	fmt.Println(b)
	fmt.Println()

	c := []int{1, 2, 3, 4}
	d := []int{5, 6, 7, 8}

	c = append(c, d...)
	fmt.Println(c)
	fmt.Println()

	// 배열 대입
	e := [3]int{1, 2, 3}
	var f [3]int

	f = e
	f[0] = 9

	fmt.Println(e)
	fmt.Println(f)
	fmt.Println()

	// 슬라이스 대입
	g := []int{1, 2, 3}
	var h []int

	h = g
	h[0] = 9

	fmt.Println(g)
	fmt.Println(h)
	fmt.Println()

	// 슬라이스 복사
	i := []int{1, 2, 3, 4, 5}
	j := make([]int, 3)

	copy(j, i)

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println()

	k := []int{1, 2, 3}
	l := make([]int, 3)

	copy(l, k)
	l[0] = 9

	fmt.Println(k)
	fmt.Println(l)
	fmt.Println()

	// 슬라이스와 용량
	m := []int{1, 2, 3, 4, 5}
	fmt.Println(len(m), cap(m))

	m = append(m, 6, 7)
	fmt.Println(len(m), cap(m))
	fmt.Println()

	n := []int{1, 2, 3, 4, 5}
	o := n[0:5]

	fmt.Println(n)
	fmt.Println(o)
	fmt.Println()

	p := []int{1, 2, 3, 4, 5}

	fmt.Println(p[0:3])
	fmt.Println(p[1:3])
	fmt.Println(p[2:5])
	fmt.Println()

	q := []int{1, 2, 3, 4, 5}

	fmt.Println(q[:])
	fmt.Println(q[0:])
	fmt.Println(q[:5])
	fmt.Println(q[0:len(q)])

	fmt.Println(q[3:])
	fmt.Println(q[:3])
	fmt.Println(q[1:3])
	fmt.Println()

	r := [5]int{1, 2, 3, 4, 5}
	s := r[:2]
	s[0] = 9

	fmt.Println(r)
	fmt.Println(s)
	fmt.Println()

	t := []int{1, 2, 3, 4, 5, 6, 7, 8}
	u := t[0:6:8]

	fmt.Println(len(u), cap(u))
	fmt.Println(u)
}
