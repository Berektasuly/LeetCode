package main

import "fmt"

func isPowerOfTwo(n int) bool {
    power := 0
    i := 1
    for i < n {
        i *= 2
        power++
        fmt.Println(power)
    }
    if i == n {
        return true
    } else {
        return false
    }
}