// N.1
package main

// N.2
import (
	"fmt"
)

// `func function_name(parameters) return_type {...}`
// `int` is default byte for integer type of 64
func Default_Integer_Type(i int) int {
	return i
}

// `func function_name(parameters) (return_type1, return_type2, ...) {...}`
// int8 = 1byte and -128 to 127
// int16 = 2bytes and -32768 to 32767
// int32 = 4bytes and -2147483648 to 2147483647
// int64 = 8bytes and -9223372036854775808 to 9223372036854775807
func Others_Integer_Type(i8 int8, i16 int16, i32 int32, i64 int64) (int8, int16, int32, int64) {
	return i8, i16, i32, i64
}

// float64 = 64bit and 8bytes float
func Default_Float_Type(f float64) float64 {
	return f
}

// float32 = 32bit and 4bytes float
// complex64 = 32bit real and 32bit imaginary
// complex128 = 64bit real and 64bit imaginary
func Others_Float_Type(f float32, c64 complex64, c128 complex128) (float32, complex64, complex128) {
	return f, c64, c128
}

func Default_Byte_Type(b byte) byte {
	return b
}

// N.3
func main() {
	fmt.Println(Default_Integer_Type(10))                  // Output : 10
	fmt.Println(Others_Integer_Type(1, 2, 3, 4))           // Output : 1 2 3 4
	fmt.Println(Default_Float_Type(1.23456789))            // Output : 1.23456789
	fmt.Println(Others_Float_Type(1.23456789, 1+2i, 1+2i)) // Output : 1.2345679  (1+2i) (1+2i)
	fmt.Println(Default_Byte_Type('A'))                    // Output : 65
}
