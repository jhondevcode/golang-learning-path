package main

import "fmt"

func main() {
    const array_size uint32 = 10
    //fmt.Print("Array size: ")
    //fmt.Scanf("%d", &array_size)
    var array [array_size]uint32
    for counter := 0; counter < int(array_size); counter++ {
        fmt.Print("Numero ", counter + 1, ": ")
        fmt.Scanf("%d", &array[counter])
    }
    var addition uint32
    for _, value := range array {
        addition += value
    }
    var promedio float32 = float32(addition) / float32(array_size)
    fmt.Println("Promedio:", promedio)
}
