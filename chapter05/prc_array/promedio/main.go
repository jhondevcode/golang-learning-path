package main

import "fmt"

func main() {
    var numbers [5]uint16
    numbers[0] = 65
    numbers[1] = 21
    numbers[2] = 78
    numbers[3] = 15
    numbers[4] = 34
    var result uint16 = 0
    for i := 0; i < len(numbers); i++ {
        result += numbers[i]
    }
    fmt.Println("Promedio:", float32(result) / float32(len(numbers)))
}
