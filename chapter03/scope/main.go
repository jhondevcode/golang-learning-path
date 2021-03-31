package main

import "fmt"

var name string = "Jhon"

func main() {
    fmt.Println(name)
    f()
}

func f()  {
    fmt.Println(name)
}
