package main

import "fmt"

func main() {
	var str string = "Résumé"
	fmt.Println(str)

	fmt.Printf("%v,%T\n",str[2],str[2]) //Here it is printing 195 instead of 233 this is happening bcoz it is reading only half binary
										//Represntation of é which is stored in str[1] and the other half is stored in str[2]

	for i,v := range str{
		fmt.Println(i,v)
	}

	fmt.Println(len(str)) //Here it returns length of bytes the length str is 8 bcoz é takes 2 byte space as per utf-8 encoding
						  //It does not return number of characters
						  
	//In go when we are dealing with strings basically we are dealing with values of arrays of bytes

	var str1 = []rune("Résume") // Runes are just unicode point numbers which represents the characters, HEre we are dealing with array of unicode numbers
	fmt.Println(str1)

	for i,v := range str1{ 
		fmt.Println(i,v)
	}

	fmt.Println(len(str1)) // Here the length/size is going to be 6 

	var myRune = 'A' //similar to char in cpp but not same

	var str3 = ""

	str3+=str // String Concatenation and strings are immutable in go (cannot change value later) str3[0]='V' not possible
			  // Here every time we concatenate a string a new string is created then it is assigned to str3 which is pretty inefficient
			  // To handle there a inbuilt function for this strings.Builder

	fmt.Println(myRune)
}