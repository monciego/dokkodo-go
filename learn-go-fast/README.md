# Learn GO Fast

## Six main points about GO

1. Statically Typed Language
2. Strongly Typed Language
3. Go is Compiled
4. Fast Compile Time
5. Built in Concurrency
6. Simplicity

## Concepts of module and packages

A package is jus a folder that contains a collection of go files. Then a collection of these packages is known as a module. And when we're initializing our project we're initializing a new module.

## Constants, Variables, and Basic Data Types

### Declaring Variables of `int` Type

To declare a variable, we use the `var` keyword followed by the variable name and its type. Similar to imports, every declared variable should be used; otherwise, it will result in an error, preventing unused variables from cluttering the code.

```go
var intNum int
fmt.Println(intNum)
```

In addition to the `int` type, we have `int8`, `int16`, `int32`, and `int64`, which specify the amount of memory or bits used to store the number. For instance, 64-bit integers can store much larger numbers than 16-bit integers but consume more memory. Trying to initialize a variable with a number larger than the maximum value for its type will result in a compiler error, known as an overflow error. The compiler doesn't throw overflow errors at runtime, leading to unexpected results without any error message.

```go
var intNum int16 = 32767
intNum = intNum + 1
fmt.Println(intNum) // -32768
```

Additionally, there are unsigned integers (`uint8`, `uint16`, `uint32`, `uint64`), which store only positive integers.

### Floating-Point Types: `float32` and `float64`

- `float32`: 32-bit decimal number
- `float64`: 64-bit decimal number, capable of storing larger and more precise decimal numbers but requiring more memory.

It's advisable to consider data types carefully rather than defaulting to `float64` or `int64`.

### Arithmetic Operations: Two Key Points

1. Arithmetic operations cannot be performed with **mixed types**, but one number can be cast to match the other type.
2. Integer division truncates the result toward zero, resulting in a rounded-down value.

### Strings

- Strings can be initialized using double quotes or backticks.
- Concatenation of strings is achieved by using the `+` operator.

```go
var greeting string = "Hello" + " " + "World"
```

- The length of a string can be obtained using the built-in `len` function, but note that it returns the number of bytes, not characters. Go uses UTF-8 encoding.
- To obtain the length of characters in a string, import the `unicode/utf8` package and use the `RuneCountInString` function.

```go
fmt.Println(utf8.RuneCountInString("y")) // 1
```

### Default Types

The default values for types are as follows:

- For integers and floats: `0`
- For strings: `""` (empty string)
- For booleans: `false`

Specifying types explicitly, especially when it's not obvious, is a good practice that enhances code readability.

### Constants

Constants are similar to variables, but their values cannot be changed after initialization. They must be explicitly initialized; there's no option for declaring a constant without an initial value.
