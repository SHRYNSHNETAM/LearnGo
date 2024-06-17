package main
import "fmt"

func main(){
	fmt.Println("Hello Go!") // For Running the Code go run file_Name

	var intN int = 1 //all int datatypes are int8,int16,int32,int64 and int is defined according to system architechture for 32-bit int is int 32 and similar for 64-bit system
	fmt.Println(intN)

	var floatN float64 = 23.66 // all float datatypes are float32 and float64
	fmt.Println(floatN)

	var boolean bool = true; // bool is also supported, if not initialized its false
	fmt.Println(boolean)

	var stringN string
	fmt.Println(stringN)

	var string2 string = `Hello
GO`  //Using BackTicks we can even Print Afte line breaks
	fmt.Println(string2)

	fmt.Println(len(stringN)) // length function for size ACC TO utf-8 encoding
	fmt.Println(len(string2))

	var int2 int //Default value if not intialized for int/int16/int32/* is 0, for string its "" empty string
	fmt.Println(int2)

	var string4 = "Temp" // Variable can also be declared like this, But you have to initialize it with some value so that variable takes its types 
	fmt.Println(string4)

	var String3 = 'T'  //THis is Rune (you can say as char) in Golang
	fmt.Println(String3)

	myVar := "Not" // variable can also be declared like this 
	fmt.Println(myVar)

	var int6, int7 int = 1, 2 // initializing multiple variable like this 
	fmt.Println(int6,int7) // Printing multiple variable like this 

	int8, int9 := 3, 4 //initializing multiple variables also like this
	fmt.Println(int8,int9)

	const myVar2 string = "Temp" //Constants also exists in GO, cant change the const later and also does not need to be used for zero error unlike variable
    // dont need to print myVar2 here for zero error but we need to use variable if intialized!!

	
}
