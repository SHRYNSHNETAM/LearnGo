package main
import "fmt"

type myStruct struct{ //struct lets you define custom/own datatypes which includes containing different datatypes
	myVar1 uint
	myVar2 uint
}

type myStruct2 struct{ //We can even store a struct into a sturct 
	myVar1 myStruct
}

func (e myStruct) myMethod() uint{ // Method attached to my myStruct struct using (e myStruct) means Now myStruct has a method called myMethod
	return e.myVar1*e.myVar2 //using e. we can access data memebers of attached structs Here we are returning myVar1*myVar2
}

type myInterface interface{
	myMthod() uint
}

func main(){
	var myVar1 myStruct
	fmt.Println(myVar1.myVar1,myVar1.myVar2)

	myVar2 := myStruct{myVar1: 45, myVar2: 56} //This way we can initialize the variable
	fmt.Println(myVar2.myVar1,myVar2.myVar2)

	myVae3 := myStruct2{myVar1: myStruct{56,23}} 
	fmt.Println(myVae3)

	myVar4 := struct{ //Anonymous Struct: Initializing and defining it in the same position
		myVar1 int
		myVar2 uint
	}{3,6}

	fmt.Println(myVar4)
	
	fmt.Println(myVar2.myMethod()) // see now myVar2 is of myStruct datatype now has access to myMethod() method

	fmt.Println()
}