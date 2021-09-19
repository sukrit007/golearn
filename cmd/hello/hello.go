package main

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/sukrit007/golearn/internal/util"
	"rsc.io/quote"
)

var lst2 []int

var z = 2 + 3i

func main() {
	lst := []int{5, 1, 2}
	lst2 = lst
	lst2[0] = 6

	sort.Ints(lst)
	fmt.Println(util.Hello("sukrit"), lst, reflect.TypeOf(lst))
	util.DoInternal()

	fmt.Println("Complex Number", z)

	for counter := 0; counter < 2; counter++ {
		fmt.Println(quote.Go(), counter)
	}
}
