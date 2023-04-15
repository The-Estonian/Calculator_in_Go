package main

import (
    "fmt"
    "github.com/The-Estonian/Calculator_in_Go/basic"
    "github.com/The-Estonian/Calculator_in_Go/advanced"
)

func main() {
    var num1 int = 1
    var num2 int = 2
    
    product := basic.Multiplication(num1, num2)
    square := advanced.Square(num2)

    fmt.Printf("The product of %d and %d is %d\n", num1, num2, product)
    fmt.Printf("The square of %d is %d\n", num2, square)
}
