# Pascal Triangle

Pascal's Triangle


Must Know

To successfully complete this project, you should revise the following Go concepts:


Slices:
        Understand how to create, access, modify, and iterate over slices.
        Use slices to generate and store rows of Pascal’s Triangle.

Functions:
        Know how to define and call functions in Go.
        Pass parameters and return values, particularly how to return a slice of slices representing Pascal’s Triangle.
Loops:
        Utilize for loops to iterate through sequences.
        Nested loops may be necessary for generating each row and calculating the values of Pascal’s Triangle.

Conditional Statements:
        Apply if statements to implement logic based on the position within Pascal’s Triangle (e.g., the edges of the triangle always being 1).

Recursion (Optional):
        While not strictly necessary, understanding recursion can provide an alternative approach to generating Pascal’s Triangle.
        Recognize base cases and recursive cases for a function that generates the triangle’s rows.

Arithmetic Operations:
        Perform addition, a fundamental operation for calculating each element of Pascal’s Triangle as the sum of the two elements directly above it.

Indexing and Slicing:
        Access elements and slices, crucial for identifying and summing the correct elements when constructing each row of the triangle.

Memory Management:
        Be mindful of how slices are stored and copied, especially when creating new rows based on the values of the previous row.

Error Handling (Optional):
        Use Go's error handling mechanisms to manage potential errors, such as invalid input types or values.

Efficiency and Optimization:
        Consider the time and space complexity of different approaches to generating Pascal’s Triangle.
        Evaluate and apply optimizations to improve the performance of the solution.
    

By revisiting these concepts, you will be well-prepared to tackle the challenges of implementing Pascal’s Triangle in Go, applying both your mathematical understanding and programming skills to develop an efficient and effective solution.
Tasks

    Pascal's Triangle
        Mandatory
        Score: 45.45% (Checks completed: 90.91%)

Create a function func pascalTriangle(n int) [][]int that returns a slice of slices of integers representing Pascal’s triangle for n:

    Returns an empty slice if n <= 0.
    You can assume n will always be an integer.

Example Usage:
```go
package main

import (
    "fmt"
)

func pascalTriangle(n int) [][]int {
    // Implementation here
}

func printTriangle(triangle [][]int) {
    for _, row := range triangle {
        fmt.Printf("[%s]\n", joinInts(row))
    }
}

func joinInts(arr []int) string {
    // Join integers as strings for printing
}

func main() {
    printTriangle(pascalTriangle(5))
}

Expected Output:

[1]
[1,1]
[1,2,1]
[1,3,3,1]
[1,4,6,4,1]
```
