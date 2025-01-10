package main

import "fmt"

func main() {
	var name string ="alain"
	fmt.Println(name)
	fmt.Print("Hello")
	
	var age int =20
	fmt.Println(age)

	fmt.Printf("Hello %s, you are %d years old\n", name, age)
	var score float64 = 25500

    fmt.Printf("Total score is %f \n",score)
	score= 3000
	fmt.Printf("Total score is %0.1f \n",score)
}