package main
import "fmt"

func main(){
	var arr [8]int32 //8 is length (fixed Length cannot be changed)
					 // int32 is datatype of entries of array cannot store any other datatype
					 //default value of entries is going to be the default value of datatype
	fmt.Println(arr[0])
	arr[2] = 2
	fmt.Println(arr[1:3]) // get value from 1 to 2 indexes

	fmt.Println(&arr[2]) // address of arr[2]

	var arr1 [3]int  = [3]int{1,3,5} // Can Be initialized like this 
	fmt.Println(arr1)

	arr2 := [5]int64{1,2,3,4,5} // Faster way to declare
	fmt.Println(arr2)

	arr3 := [...]int32{1,2,4,3,4} //no need to write the length/size
	fmt.Println(arr3)

	//Slices are just wrappers around arrays gives more powerful interface to data\
	//Or you can say it is similar to vector in Cpp

	var slc []int = []int{1,2,3}
	fmt.Println(slc)

	fmt.Println(len(slc),cap(slc)) // length = 3, capacity = 3

	slc  = append(slc, 2) // values can be added means slice is not of fixed size
	// when new values are added a new array is formed with new address and then all the values are copied there 
	//Then size of array must not be equal to actual capacity of formed array

	fmt.Println(slc)
	fmt.Println(len(slc),cap(slc))  // length = 4, capacity =6 new array but cannot access it[1,2,3,2,*,*]

	var slc2 = []int{6,7}
	fmt.Println(slc2)

	slc2 = append(slc2, slc...) // can add two slices
	fmt.Println(slc2)

	var slc3 = make([]int, 2, 8) // another way of creating slice where 2 is length and 8 is capacity (not necessary to specify)
	fmt.Println(slc3)
	//Whereever a new element is added if there is no extra capacity then a new array allocation happens
	//which takes some times ..

	//Map

	var mp map[string]int = make(map[string]int) // here "key" is of string datatype and "value" is of int datatype
	fmt.Println(mp)

	var mp2 = map[string]int{"Hello":23,"Go":34,"Golang":55}
	fmt.Println(mp2["Hello"])
	fmt.Println(mp2["Jason"]) // if not present default value  f value datatype is returned here it is Zero
	var age1, ok = mp2["Jason"] // In Go a map always return something luckily it also return a second variable which tells if the key is present or not
	// if ok==true then Present or if ok==false then notPresent
	fmt.Println(age1,ok)

	delete(mp2,"Hello") //For deleting an entry from map

	for name, age:= range mp2{ // for iterating over a map we can use range function 
		fmt.Println(name,age) // here name is the key of each iteration and age is the value of that key
	}
	// looping is not ordered means Running this loop several times will gives different result

	for i,val := range arr3{ // similarly we can iterate over a array or a slice
		fmt.Println(i,val) // here i is index and val is values at that index
	}

	//Nop while loop in Go but can be achieved using for loop 

	i := 0
	for i<5{	
		fmt.Println(i)
		i=i+1
	}
	for j:=0;j<5;j++{// traditinal for loop
		fmt.Println(j)
	}

	// i--, i++, i+=, i-=, i/=, i*= also works in Go
}