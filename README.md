## GoInt128

Simple library for 128-bit integer arithmetic.
Pretty simple code to handle 128-bit unsigned integers:
addition, subtraction, shifts, multiplication, division, parsing and formatting.

Main code is licensed by LGPL 2.1 license,
code to support Go1.9 is licensed by LICENSE_BSD license.

Available functions:

* UInt128.Add - add two integers
* UInt128.AddC - add two integers including carry, return sum and carry
* UInt128.Add64 - add 64-bit unsigned integer to 128-bit unsigned integer
* UInt128.Sub - subtract integer
* UInt128.SubB - subtract integer including borrow, return difference and borrow
* UInt128.Sub64 - subtract 64-bit unsigned integer from 128-bit unsigned integer
* UInt128.Mul - multiply two integers and return low 128 bits of product
* UInt128.MulFull - multiply two integers and return full product (first high, second low)
* UInt128.Shl - shift left integer
* UInt128.Shr - logival shift right integer
* UInt128.Div64 - divide 128-bit unsigned integer by 64-bit value, return 128-bit quotient and 64-bit remainder
* UInt128DivFull - divide 256-bit unsigned integer by 128-bit value, return 128-bit quotient and remainder
* UInt128.Format - format integer to decimal string
* ParseUInt128 - parse integer from string, return value and error (will be nil if no error)
