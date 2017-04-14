package main

import "fmt"

func main() {
    i := 0
    sum := 0
    for {
      n := fib(i)
      if (n > 4000000) {
         fmt.Println(sum)
         break
      }
      if n % 2 == 0 {
         sum += n
      }
      i++ 
    }
}

func fib(n int) int {
    if n < 2 { return 1}
    return fib(n - 1) + fib(n - 2)
}
