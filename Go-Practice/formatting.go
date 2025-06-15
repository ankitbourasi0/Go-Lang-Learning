package main

import "fmt"

func main() {

	//General Formatting Verb
	/*
		-------------------------------------------------
		Verb        Description
		%v			Prints the value in the default form
		%#v			Print the value in Go syntax format
		%T			Print the TYPE of value
		%%			Print the '%' Sign.
		-------------------------------------------------
	*/
	var i = 15.5
	var txt = "Hello World"
	fmt.Printf("%v \n", i)   //15.5
	fmt.Printf("%#v \n", i)  //15.5
	fmt.Printf("%T \n", i)   //float64
	fmt.Printf("%v%% \n", i) //15.5%

	fmt.Printf("%v \n", txt)  //Hello World
	fmt.Printf("%#v \n", txt) //"Hello World"
	fmt.Printf("%T \n", txt)  //string

	//Integer Formatting Verb
	/*
		-------------------------------------------------
		Verb        Description
		%b 			Prints the Base 2 Value
		%d			Prints the Base 10 Value
		%+d			Prints the Base 10 with Sign
		%o			Prints the Base 8 value
		%O			Prints the Base 8 value, with Prefix 0o
		%x			Prints the Base 16 value, in Lowercase
		%X			Prints the Base 16 value, in uppercase
		%#x			Prints the Base 16 value, with Prefix 0x
		%4d			Prints the left padded(space) value with specified width(e.g 4)
		%-4d		Prints the right padded(space) value with specified width(e.g 4)
		%04d		Prints the padding of 0 and value with specified width(e.g 4)
		-------------------------------------------------
	*/

	var j = 15
	fmt.Printf("Base2: %b \n", j)                      //1111
	fmt.Printf("Base10: %d \n", j)                     //
	fmt.Printf("Base10 with sign: %+d \n", j)          //+15
	fmt.Printf("Base8: %o \n", j)                      //0o
	fmt.Printf("Base8 with Prefix 0o: %O \n", j)       //0o17
	fmt.Printf("Base16 lowercase: %x \n", j)           //f
	fmt.Printf("Base16: uppercase %X \n", j)           //F
	fmt.Printf("Base16 with Predix 0x: %#x \n", j)     //0xf
	fmt.Printf("left Padding of width 4: %4d \n", j)   //   15
	fmt.Printf("right Padding of width 4: %-4d \n", j) //  15
	fmt.Printf("Padding of 0 of width 4: %04d \n", j)  //  0015

	//String Formatting Verb
	/*
		-------------------------------------------------
		Verb        Description
		%s 			Prints the value in Plain String
		%q			Prints the value in double quotes String
		%8s			Prints the value in plain String with Padding(8 width) right
		%-8s		Prints the value in plain String with Padding(8 width) left
		%x			Prints the value in Hex dump of bytes values
		% x			Prints the value in Hex dump with space
		-------------------------------------------------
	*/

	fmt.Printf("Plain string: %s \n", txt)                                        // Hello World
	fmt.Printf("double quoted string: %q \n", txt)                                // "Hello World"
	fmt.Printf("Plain string with Padding(justified 8 width) right: %8s \n", txt) // Hello World
	fmt.Printf("Plain string with Padding(justified 8 width) left: %-8s \n", txt) // Hello World
	fmt.Printf("Hex dump in Bytes value: %x \n", txt)                             // 48656c6c6f20576f726c64
	fmt.Printf("Hex dump in Bytes value with Spaces: % x \n", txt)                // 48 65 6c 6c 6f 20 57 6f 72 6c 64

	//Boolean Formatting Verb
	/*
		-------------------------------------------------
		Verb       Description
		%t			Prints the value in Boolean(true or false same as %v)
		-------------------------------------------------
	*/
	var b = true
	var c = false
	fmt.Printf("Boolean value :%t \n", b) //true
	fmt.Printf("Boolean value :%t \n", c) //false

	//float Formatting Verb
	/*
		-------------------------------------------------
		Verb        Description
		%e 			Prints the scientific notation with "e" as exponent
		%f			Prints the decimal point, no exponent
		%.2f	    Prints the default width , upto 2 decimal(precision)
		%6.2f		Prints the width 6, with precision of 2
		%g			Prints the exponent as needed, only necessary digits.
		-------------------------------------------------
	*/

	floatNo := 3.141
	fmt.Printf("Float value with exponent scientific notation: %e \n", floatNo)             // 3.141000e+00
	fmt.Printf("Float value: %f\n", floatNo)                                                // 3.141000
	fmt.Printf("Float value default width and upto 2 decimal(precision):  %.2f\n", floatNo) // 3.14
	fmt.Printf("Float value width 6, decimal point 2: %6.2f \n", floatNo)                   //   3.14
	fmt.Printf("Float value as Needed no Exponenet: %g\n", floatNo)                         //  3.141

}
