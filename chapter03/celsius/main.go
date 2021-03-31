package main

import "fmt"

// Programa para comvertir grados fahrenheit a celsius
func main() {
    fmt.Print("fahrenheit grades: ")
    var far float64
    fmt.Scanf("%f", &far)
    // formula: c = (f-32) * 5/9
    celsius := ((far - 32) * 5 ) / 9
    fmt.Println("Celsius:",  celsius)
}
