package main

import (
    "fmt"
    "github.com/The-Estonian/Calculator_in_Go/basic"
    // "github.com/The-Estonian/Calculator_in_Go/advanced"
)

func main() {
    var input1 int = 0
    var input2 int = 0
    command := ""
    output := 0
    for {
        for {
            fmt.Print("NUM: ")
            fmt.Scan(&input1)
            if input1 != 0 {
                break
            }
        }
        fmt.Print(`+ - / *`)
        fmt.Scan(&command)
        fmt.Print("NUM: ")
        fmt.Scan(&input2)
        if command == "+" {
            output = basic.Addition(input1,input2)
        } else if command == "-" {
           output = basic.Substraction(input1,input2)
        } else if command == "/" {
           output = basic.Division(input1,input2)
        } else if command == "*" {
           output = basic.Multiplication(input1,input2) 
        }
        fmt.Println(output)
    }
}
