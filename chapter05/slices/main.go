package main

import "fmt"

func main() {
    fmt.Print("Array size: ")
    var array_size uint32
    fmt.Scanf("%d", &array_size)
    x := make([]uint32, array_size)
    // filling the array
    for i := 0; uint32(i) < array_size; i++ {
        fmt.Print("Nunmber ", i + 1, ": ")
        fmt.Scanf("%d", &x[i])
    }
    var suma uint32 = 0
    for _, value := range x {
        suma += value
    }
    promedio := suma / array_size
    fmt.Println("Promedio: ", promedio)
}
