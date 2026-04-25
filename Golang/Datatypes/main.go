package main

import (
	"fmt"
)

// `func function_name(parameters) return_type {...}`
// `int` is user defined type and use system dependent size (32bit or 64bit)
func Default_Integer_Type(i int) int {
	return i
}

// `func function_name(parameters) (return_type1, return_type2, ...) {...}`
// int8 = 1byte and -128 to 127
// int16 = 2bytes and -32768 to 32767
// int32 = 4bytes and -2147483648 to 2147483647
// int64 = 8bytes and -9223372036854775808 to 9223372036854775807
// rune = alias for int32 and represents a Unicode code point
func Others_Integer_Type(i8 int8, i16 int16, i32 int32, i64 int64, r rune) (int8, int16, int32, int64, rune) {
	return i8, i16, i32, i64, r
}

// float64 = 64bit and 8bytes float
func Default_Float_Type(f64 float64) float64 {
	return f64
}

// float32 = 32bit and 4bytes float
// complex64 = 32bit real and 32bit
//
//	imaginary
//
// complex128 = 64bit real and 64bit imaginary
func Others_Float_Type(f32 float32, c64 complex64, c128 complex128) (float32, complex64, complex128) {
	return f32, c64, c128
}

// `func function_name(parameters) return_type {...}`
// `int` is user defined type and use system dependent size (32bit or 64bit)
func Default_Unsigned_Type(ui uint) uint {
	return ui
}

func Others_Unsigned_Type(u8 uint8, u16 uint16, u32 uint32, u64 uint64) (uint8, uint16, uint32, uint64) {
	return u8, u16, u32, u64
}

func String_Type(s string) string {
	return s
}

func main() {
	fmt.Println(Default_Integer_Type(10))                  // Output : 10
	fmt.Println(Others_Integer_Type(1, 2, 3, 4, 'A'))      // Output : 1 2 3 4 65
	fmt.Println(Default_Float_Type(1.23456789))            // Output : 1.23456789
	fmt.Println(Others_Float_Type(1.23456789, 1+2i, 1+2i)) // Output : 1.2345679  (1+2i) (1+2i)
	fmt.Println(Default_Unsigned_Type(10))                 // Output : 10
	fmt.Println(Others_Unsigned_Type(65, 66, 67, 68))      // Output : 65 66 67 68

	fmt.Println(String_Type("Hello Go!")) // Output : Hello Go!
}
