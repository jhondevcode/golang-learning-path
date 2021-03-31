package main;

import "fmt"

func main() {
    var pairs [5]uint8
    pairs[0] = 2
    pairs[1] = 4
    pairs[2] = 6
    pairs[3] = 8
    pairs[4] = 10
    for index, value := range pairs {
        fmt.Println(index, value)
    }

    fmt.Println("")

    impairs := [5]uint8{1, 3, 5, 7, 9}
    for _, value := range impairs {
        fmt.Println(value)
    }
}
