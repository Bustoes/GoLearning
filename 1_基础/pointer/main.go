package main

import "fmt"

func main() {
	/*
	a := 10
	pa := &a

	fmt.Printf("a : %v, *pa : %v, &a : %p, &&a : %p\n", &a, *pa, pa, &pa)
	*/
	a := 10
	mod1(a)
	fmt.Println(a)
	mod2(&a)
	fmt.Println(a)

	//new	不常用
	p1 := new(int)
	fmt.Printf("%T\n", p1)
}

func mod1(x int) {
	x = 100
}

func mod2(y *int) {
	*y = 100
}




