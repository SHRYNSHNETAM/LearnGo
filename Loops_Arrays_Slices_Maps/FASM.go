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
}