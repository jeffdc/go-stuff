package main

import (
    "code.google.com/p/go-tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
    vals := make([][]uint8, dx)
    for x := 0; x<dx; x++ {
        inner := make([]uint8, dy)
        for y := 0; y<dy; y++ {
            inner[y] = uint8(100*x*(y*y)/(x+7))
        }
        vals[x] = inner
    }
    return vals
}

func main() {
    pic.Show(Pic)   
}