package main

import (
	"fmt"
	"math"
)

type person struct {
	firstName string
	secondName string
}

type manager struct {
	person
	position string
}

func (p person) String() string{
	return fmt.Sprintf(`
	Имя:     %s  
	Фамилия: %s
	`, p.firstName, p.secondName)
}



func main() {
	var p person = person{
		firstName: "vadim",
		secondName: "ivlev",
	}
	fmt.Println("Привет", math.Pi, p)
}


func consoleInputOutput(){
	var num0 float64
	var num1 float64
	fmt.Print("Enter the numbers:")
	fmt.Scanf("%f %f",&num0, &num1)
	fmt.Printf("Numbers %.5f %.5f", num0,  num1)
}