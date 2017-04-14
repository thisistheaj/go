package main

import "fmt"
import "os"

func main() {
    var Args []string = os.Args
    if len(Args) < 2 {
         fmt.Println("Why hello there, stranger!")
    } else {
         fmt.Println("Hello, " + Args[1] + "!")
    }
}
