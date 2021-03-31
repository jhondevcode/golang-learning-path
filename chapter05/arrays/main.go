package main

import "fmt"

func main() {
    var x [5]int
    x[4] = 100
    fmt.Println(x)
    fmt.Println(x[4])

    fmt.Println("")

    notas := [5]uint16{6, 78, 69, 35, 46}
    fmt.Println(notas)
    for _, value := range notas {
        fmt.Println(value)
    }
}
