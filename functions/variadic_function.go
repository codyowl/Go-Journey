package main

import "fmt"

func my_variadic_function(names ...string){
	for _, name := range names {
		fmt.Println("The name is : ", name)
	}
}
// Similar to *args in python

func main() {
	my_variadic_function("Ilayaraja", "Rahman", "Yuvan")
}


