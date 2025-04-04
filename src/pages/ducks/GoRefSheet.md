---
# don't modify
layout: "../../layouts/LayoutSingle.astro"
title: "GoRefSheet"
---

Certainly! Here's a comprehensive syntax guide for the Go programming language in markdown format. This guide covers many common operations and data structures that are shared among other programming languages:

# Go Programming Language Syntax Guide

## Table of Contents

1. [Basic Structure](#basic-structure)
2. [Variables and Data Types](#variables-and-data-types)
3. [Operators](#operators)
4. [Control Structures](#control-structures)
5. [Functions](#functions)
6. [Arrays and Slices](#arrays-and-slices)
7. [Maps](#maps)
8. [Structs](#structs)
9. [Pointers](#pointers)
10. [Interfaces](#interfaces)
11. [Error Handling](#error-handling)
12. [Goroutines and Channels](#goroutines-and-channels)

## Basic Structure

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

## Variables and Data Types

```go
// Variable declaration
var x int
var y string = "hello"
z := 3.14 // Short variable declaration

// Constants
const PI = 3.14159

// Basic data types
var (
    a bool = true
    b int = 42
    c float64 = 3.14
    d string = "Go"
    e rune = 'A'
)

// Complex types
var f complex128 = 1 + 2i
```

## Operators

```go
// Arithmetic operators
a + b
a - b
a * b
a / b
a % b

// Comparison operators
a == b
a != b
a < b
a <= b
a > b
a >= b

// Logical operators
a && b
a || b
!a

// Bitwise operators
a & b
a | b
a ^ b
a << b
a >> b
```

## Control Structures

```go
// If statement
if x > 0 {
    // code
} else if x < 0 {
    // code
} else {
    // code
}

// For loop
for i := 0; i < 10; i++ {
    // code
}

// While loop (using for)
for condition {
    // code
}

// Infinite loop
for {
    // code
}

// Switch statement
switch x {
case 1:
    // code
case 2, 3:
    // code
default:
    // code
}

// Range
for index, value := range collection {
    // code
}
```

## Functions

```go
// Basic function
func add(a, b int) int {
    return a + b
}

// Multiple return values
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Variadic function
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

// Anonymous function
f := func(x int) int {
    return x * x
}
```

## Arrays and Slices

```go
// Array
var arr [5]int

// Array initialization
arr := [5]int{1, 2, 3, 4, 5}

// Slice
var slice []int

// Slice initialization
slice := []int{1, 2, 3, 4, 5}

// Make a slice
slice := make([]int, 5, 10)

// Append to slice
slice = append(slice, 6, 7, 8)

// Slice operations
subSlice := slice[1:4]
```

## Maps

```go
// Map declaration
var m map[string]int

// Map initialization
m := make(map[string]int)

// Map literal
m := map[string]int{
    "apple": 1,
    "banana": 2,
}

// Map operations
m["cherry"] = 3
value, exists := m["apple"]
delete(m, "banana")
```

## Structs

```go
// Struct definition
type Person struct {
    Name string
    Age  int
}

// Create a struct
p := Person{Name: "Alice", Age: 30}

// Access struct fields
fmt.Println(p.Name)

// Struct method
func (p Person) SayHello() {
    fmt.Printf("Hello, I'm %s\n", p.Name)
}
```

## Pointers

```go
// Pointer declaration
var p *int

// Create a pointer
x := 42
p = &x

// Dereference a pointer
fmt.Println(*p)

// New function
p := new(int)
*p = 42
```

## Interfaces

```go
// Interface definition
type Shape interface {
    Area() float64
}

// Implement interface
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}
```

## Error Handling

```go
// Error handling
result, err := someFunction()
if err != nil {
    log.Fatal(err)
}

// Custom error
type MyError struct {
    message string
}

func (e MyError) Error() string {
    return e.message
}
```

## Goroutines and Channels

```go
// Goroutine
go someFunction()

// Channel declaration
ch := make(chan int)

// Send to channel
ch <- 42

// Receive from channel
value := <-ch

// Buffered channel
ch := make(chan int, 10)

// Close channel
close(ch)

// Select statement
select {
case v1 := <-ch1:
    // code
case v2 := <-ch2:
    // code
default:
    // code
}
```

Certainly! I apologize for not providing more detailed information on arrays in the previous response. I'll expand on arrays in Go, including their declaration, initialization, and common operations. Here's an additional section on arrays in markdown format:

## Arrays

Arrays in Go are fixed-size sequences of elements of the same type. Unlike slices, the size of an array is part of its type, which means that `[5]int` and `[10]int` are distinct types.

### Array Declaration

```go
// Declare an array of 5 integers
var arr [5]int

// Declare an array of 3 strings
var strArr [3]string

// Declare a 2D array (3x4 matrix of integers)
var matrix [3][4]int
```

### Array Initialization

```go
// Initialize with values
arr := [5]int{1, 2, 3, 4, 5}

// Initialize with sparse array
sparseArr := [10]int{1, 5: 4, 6, 8: 100}

// Let compiler count the array elements
countedArr := [...]int{2, 4, 6, 8, 10}

// Initialize 2D array
matrix := [2][3]int{
    {1, 2, 3},
    {4, 5, 6},
}
```

### Accessing and Modifying Array Elements

```go
// Access an element
firstElement := arr[0]

// Modify an element
arr[2] = 10

// Get array length
length := len(arr)
```

### Iterating Over Arrays

```go
// Using traditional for loop
for i := 0; i < len(arr); i++ {
    fmt.Printf("Element %d: %d\n", i, arr[i])
}

// Using range
for index, value := range arr {
    fmt.Printf("Element %d: %d\n", index, value)
}
```

### Array Comparison

```go
// Arrays of the same type can be compared
arr1 := [5]int{1, 2, 3, 4, 5}
arr2 := [5]int{1, 2, 3, 4, 5}
areEqual := arr1 == arr2 // true
```

### Copying Arrays

```go
// Create a copy of an array
arrCopy := arr

// Modify the copy (doesn't affect the original)
arrCopy[0] = 100
```

### Arrays as Function Parameters

```go
// Pass array by value (creates a copy)
func printArray(arr [5]int) {
    for _, v := range arr {
        fmt.Printf("%d ", v)
    }
    fmt.Println()
}

// Pass array by reference
func modifyArray(arr *[5]int) {
    arr[0] = 100
}

// Usage
arr := [5]int{1, 2, 3, 4, 5}
printArray(arr)
modifyArray(&arr)
```

### Converting Arrays to Slices

```go
// Create a slice from an array
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[:]

// Create a slice from part of an array
partialSlice := arr[1:4]
```

Remember that while arrays in Go are useful in certain scenarios, slices are more commonly used due to their flexibility and dynamic nature. Arrays are particularly useful when you need a fixed-size collection or when working with very large data sets where you want to avoid the overhead of slice capacity management.
