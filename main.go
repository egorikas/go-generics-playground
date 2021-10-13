package main

import (
	"fmt"
   )
   
   func print[T any](s []T) {
	for _, v := range s {
	 fmt.Println(v)
	}
   }
   
   func main() {
	print([]string{"tet", "test"})
   }