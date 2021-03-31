package main

import "fmt"

func main() {
    var numbers [5]uint32
    // capturamos 5 numeros de la entrada estandar
    fmt.Println("Ingresa 5 numeros")
    for counter := 0; counter < len(numbers); counter++ {
        fmt.Print("Elemento ", counter, ": ")
        fmt.Scanf("%d", &numbers[counter])
    }
    fmt.Println(numbers)
    var suma uint32
    for _, value := range numbers {
        suma = suma + value
    }
    var promedio float32 = float32(suma) / float32(len(numbers))
    fmt.Println("Promedio:", promedio)
}
