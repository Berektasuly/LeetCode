func isPowerOfFour(n int) bool {
    if n <= 0 {
        return false
    }
    if n == 1 {
        return true
    }
    for n % 4 == 0 {
        n /= 4
    }
    if n == 1 {
        return true
    } else {
        return false
    }
}