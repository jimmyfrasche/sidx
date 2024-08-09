package sidx_test

import (
	"fmt"

	"github.com/jimmyfrasche/sidx"
)

func ExampleIndex() {
	// s is a list of indices into vs
	// for each we print * if it's invalid and vs[j] if it's valid.
	s := []int{-100, -50, -3, -1, 1, 3, 50, 100}
	vs := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	for _, i := range s {
		i = sidx.Index(len(s), i)
		if i < 0 {
			fmt.Print("*")
		} else {
			fmt.Print(vs[i])
		}
	}
	// Output:
	// **FHBD**
}

func ExampleN() {
	vs := []string{"A", "B", "C", "D"}
	vs = vs[1:sidx.N(vs, -1)]
	for _, v := range vs {
		fmt.Print(v)
	}
	// Output:
	// BC
}

func ExampleAtOk() {
	vs := []string{"A", "B", "C", "D", "E"}
	for _, i := range []int{2, -2, 40} {
		fmt.Printf("%2d: ", i)
		fmt.Println(sidx.AtOk(vs, i))
	}
	// Output:
	//  2: C true
	// -2: D true
	// 40:  false
}

func ExampleAtOr() {
	vs := []string{"A", "B", "C", "D", "E"}
	for _, i := range []int{2, -2, 40} {
		fmt.Printf("%2d: ", i)
		fmt.Println(sidx.AtOr(vs, i, "!"))
	}
	// Output:
	//  2: C
	// -2: D
	// 40: !
}

func ExampleAt() {
	vs := []string{"A", "B", "C", "D", "E"}
	for _, i := range []int{2, -2} {
		fmt.Printf("%2d: ", i)
		fmt.Println(sidx.At(vs, i))
	}
	// Output:
	//  2: C
	// -2: D
}

func ExampleEnd() {
	v := sidx.End([]string{"A", "B", "C", "D"})
	fmt.Println(v)
	// Output:
	// D
}

func ExampleSlice() {
	vs := []string{"A", "B", "C", "D"}
	vs = sidx.Slice(vs, 1, -1)
	fmt.Println(vs)
	// Output:
	// [B C]
}

func ExamplePop() {
	var v string
	vs := []string{"A", "B", "C", "D"}
	for len(vs) > 0 {
		v, vs = sidx.Pop(vs)
		fmt.Print(v)
	}
	// Output:
	// DCBA
}
