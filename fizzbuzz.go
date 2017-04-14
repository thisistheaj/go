package main

import "fmt"
import "os"
import "strconv"

func main(){
    Args := os.Args
    var max int

    if len(Args) < 2 {
        max = 100
    } else {
        max, _ = strconv.Atoi(Args[1])
    }

    for i := 0; i < max; i++ {

        isFizz := i % 3 == 0
        isBuzz := i % 5 == 0

        if isFizz && isBuzz {
            fmt.Print("fizzbuzz")
        } else if isFizz {
            fmt.Print("fizz")
        fmt.Print("\t")
        } else if isBuzz {
            fmt.Print("buzz")
        fmt.Print("\t")
        } else {
            fmt.Print(i)
        fmt.Print("\t")
        }
        if i % 8 == 0 {
            fmt.Print("\n")
        }
    }
}
