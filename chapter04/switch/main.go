package main

import "fmt"

func main() {
    fmt.Print("Ingrese un numero entre 0 y 10: ")
    var number uint8
    fmt.Scanf("%d", &number)
    switch number {
        case 0: fmt.Println("Cero")
        case 1: fmt.Println("Uno")
        case 2: fmt.Println("Dos")
        case 3: fmt.Println("Tres")
        case 4: fmt.Println("Cuatro")
        case 5: fmt.Println("Cinco")
        case 6: fmt.Println("Seis")
        case 7: fmt.Println("Siete")
        case 8: fmt.Println("Ocho")
        case 9: fmt.Println("Nueve")
        case 10: fmt.Println("Diez")
        default: fmt.Println("Unknow number")
    }
}
