package main

import ("fmt")



func main(){
	slc1 := []int{1,2,3,4,5}
	sum := sumSliceInt(slc1)
	fmt.Println(sum)

	var slc2 = []float32{3.5,2.3,5.33}
	fmt.Println(sumSliceFloat(slc2))

	fmt.Println(sumSlice(slc1)) //Passing different datatypes in same function 
	fmt.Println(sumSlice(slc2)) //Passing Float slice
}

func sumSliceInt(slc []int) int{  // for sumation of int datatype slice
	var sum int
	for i:=0;i<len(slc);i++{
		sum+=slc[i]
	}
	return sum
}

func sumSliceFloat(slc []float32) float32{ //For sumation of float32 datatype slice
	var sum float32
	for i:=0;i<len(slc);i++ {
		sum+=slc[i]
	}
	return sum
}
// Instead of writing same function for different differnt datatype we can use generics function

func sumSlice[T int | float32 | float64](slc []T) T{ //Here we need to specify which datatypes inputs can acces this function
	var sum T                                         //We can use any kerword instead of defining the datatypes
	for i:=0;i<len(slc);i++ {                        //any keyword gives access to all datatype inputs
		sum+=slc[i]
	}
	return sum
}
