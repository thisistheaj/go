package main

import "fmt"
import "os"

func main() {
    var Args []string = os.Args
    if len(Args) < 2 {
         fmt.Println("Saluton, homo. Bonvene!")
    } else {
         fmt.Println("Saluton, " + Args[1] + "!")
    }
}
