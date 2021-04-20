package main

import "fmt"

const ( //we can write list of constans here
	pi = 3.1415

	o = iota
	t = iota
	s = iota + 6
	ti
	g = 2 << iota
)

func main() {
	fmt.Println("Hello World from moduleg")

	var i int
	i = 42
	//float-float32,float64
	var f float32 = 3.14
	var f1 float64 = 3.14
	fmt.Println(i, f, f1)
	//string
	firstName := "sri"
	fmt.Println(firstName)
	//boolean
	b := true
	fmt.Println(b)
	//complex
	c := complex(3, 4)
	fmt.Println(c)
	//pull real and imaginary part from complex number
	r, im := real(c), imag(c)
	fmt.Println(r, im)
	//pointer data type
	var l *string = new(string)
	*l = "sai"
	fmt.Println(*l)
	//address of operator
	m := "prem"
	fmt.Println(m)

	ptr := &m
	fmt.Println(ptr, *ptr)
	//creating constants
	const p = 3.14

	fmt.Println(p)

	fmt.Println(pi)

	fmt.Println(o, t)
	fmt.Println(s)
	fmt.Println(ti)
	fmt.Println(g)

}
