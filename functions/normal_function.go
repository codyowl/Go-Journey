package main

import "fmt"

func just_some_normal_function(text string, number int) (string, int) {
	return text, number
}


func main(){
	text, number := just_some_normal_function("hi", 5)
	fmt.Println("Entered number is : ", number)
	fmt.Println("Entered text is : ", text)
}

