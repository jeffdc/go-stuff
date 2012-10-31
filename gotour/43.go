package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    z := 1.0
    prev := 1.0
    delta := 0.000000000000001
    i := 0
    for ; ;i++ {
        z = z - (z*z - x)/(2*z)
        if (z < prev + delta) && (z > prev - delta) {
            break
        }
        prev = z
            
    }
    fmt.Println("took", i)
    return z
    
}

func main() {
    data := []float64{1, 2, 3, 4}
    
    for _, n := range data {
        fmt.Println(Sqrt(n), math.Sqrt(n))
    }
}