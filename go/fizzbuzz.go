package main
import (
    "fmt"
    "strconv"
)

/**
 * Write a program that prints the numbers from 1 to 100. But for multiples
 * of three print “Fizz” instead of the number and for the multiples of five
 * print “Buzz”. For numbers which are multiples of both three and five print
 * “FizzBuzz”.
*/

func fizzbuzz() {
    for i := 1; i <= 100; i++ {
        if i % 15 == 0 {
            fmt.Println("FizzBuzz")
        } else if i % 3 == 0 {
            fmt.Println("Fizz")
        } else if i % 5 == 0 {
            fmt.Println("Buzz")
        } else {
            fmt.Println(i)
        }
    }
}

func fizzbuzz2() {
    for i := 1; i <= 100; i++ {
        output := ""
        if i % 3 != 0 && i % 5 != 0 {
            output += strconv.Itoa(i)
        }
        if i % 3 == 0 {
            output += "Fizz"
        }
        if i % 5 == 0 {
            output += "Buzz"
        }
        fmt.Println(output)
    }
}
