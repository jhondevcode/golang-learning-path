package main

import "fmt"

func main() {
    var names [5]string
    names[0] = "Jhon"
    names[1] = "Jeff"
    names[2] = "Elon"
    names[3] = "Linus"
    names[4] = "Ada"
    fmt.Println(names)
    fmt.Println("Recorriendo con for...")

    for index, element := range names {
        fmt.Println(index, element)
    }
}
