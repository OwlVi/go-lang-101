// package
package main

// import
import (
	"fmt"

	array_101 "github.com/OwlVi/go-lang-101/array"
	func_101 "github.com/OwlVi/go-lang-101/func"
	struct_101 "github.com/OwlVi/go-lang-101/struct"
)

// func

// create many var by var()
// var (
// 	x = 1
// 	y = 2
// 	z = 3
// )

func oldmain() {
	fmt.Println("Hello world")
	// create variable use var
	var x int = 10
	// create variable use :=
	y := 15
	// create variable use const
	const z = 20
	fmt.Println(x, y, z)
	// do condition
	if x > 10 {
		fmt.Println("x less than 10")
	} else if x < 10 {
		fmt.Println("x more than 10")
	} else {
		fmt.Println("x equal 10")
	}

	//do switch
	switch x {
	case 5:
		fmt.Println("x = 5")

	default:
		fmt.Println(x)
	}

	// while loop
	i := 10
	for i < 15 {
		fmt.Print(i, " ")
		i++
	}
	fmt.Println()

	// for loop
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// foreach
	for i := range 5 {
		fmt.Print(i, " ")
	}
	fmt.Println()

	fmt.Println(func_101.Foofoo())
	fmt.Println(func_101.Foo2(x, y))
	fmt.Println(func_101.Foo3(float32(x), y))
	fmt.Println(func_101.Foo4(x, y))

	// _ skip var dont use
	a, _ := func_101.Foo4(x, y)
	fmt.Println(a)
	array_101.Array()

	g := struct_101.Set_struct(10, "test")

	fmt.Println(g.Get_x(), g.Get_y())
}
