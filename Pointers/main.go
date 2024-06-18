package main

import "fmt"

func main(){
	var p1 *int //Declaring a p1 pointers which is addressing to anyone means it is <nil>
	fmt.Println(p1) // It will print <nil>

	var p2 *int = new(int) //Declaring a p2 pointers and initializing it with a address of random int variable

	fmt.Println(p2) //it will print the address of int variable
	fmt.Println(*p2) //it will print the value of int variable , means it will print where p2 is pointing Here 0

	var i1 int = 1
	var p3 *int = &i1 //& denotes the address of the variable since pointer is just an addresss so we can store any address in it

	fmt.Println(p3) //it will print the address of int variable(i1)
	fmt.Println(*p3) //it will print 1

	var slice = []int{2,3,4}
	var copySlice = slice //This is copying the pointer not creating a new slice where values of original slice is copies

	copySlice[2] = 8 // This will also change the values from original slice bcoz both are pointing to same data

	fmt.Println(slice)
	fmt.Println(copySlice)

	var arr = [5]int{1,2,3,4,5}
	var arr2 = square(arr)
	var arr3 = square2(&arr)
	fmt.Println(arr) 
	fmt.Println(arr2) // since it is returning a new copy of original array with squared value thats why its going to be diffferent
	fmt.Println(arr3) //it will be same as arr
}

func square(arr [5]int) [5]int{//Here we are passing a new array which is a copy of original array
	for i, val := range arr{   
		arr[i]=val*val
	}
	return arr
}

func square2(arr *[5]int) [5]int{ //Here we are passing the pointer means creating a new copy of pointer which points to original array
	for i, val := range arr{     //means we are directly modifying our original array This saves time and memory 
		arr[i]=val*val
	}
	return *arr
}