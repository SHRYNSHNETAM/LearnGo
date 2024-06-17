package main
import (
		"errors" // package for setting error datatypes
		"fmt"
)

func main(){
	temp()
	printMe("Hello from PrintMe")
	fmt.Println(Divide(2,3))
	var d, r, erro = MultipleReturn(0,4)
	if erro!=nil{
		fmt.Println(erro.Error()) // MEthod to print error
	}
	fmt.Println(d,r)

	switchimp(2,3)
}

func temp(){
	fmt.Println("Hello from temp")
}

func printMe(Var string){ //if not returning any 
	fmt.Println(Var)
}

func Divide(var1 int, var2 int) int { // You Gotta define what datatype your function is returning 
	return var2/var1
}

func MultipleReturn(var1 int, var2 int) (int, int, error){
	var err error
	if var1==0{
		err = errors.New("Cannot Divide Zero") // Method to set error datatypes 
		return 0,0,err
	}
	return var2/var1, var2%var1, err;
}

func switchimp (var1 int, var2 int){
	switch{ // similar works as if else statements
	case var1==var2:
		fmt.Println("Equal")
	case var1>var2:
		fmt.Printf("%v is greater than %v\n", var1, var2)
	default:
		fmt.Printf("%v is greater than %v\n", var2, var1)
	}

	switch var1{ //switch with a variable for this variable now we can write conditions
	case 0: // if var1 == 0 this this will execute
		fmt.Printf("%v is Zero\n",var1)
	default:
		fmt.Printf("%v is Greater than Zero\n",var1)
	}
}