package main

import (
    "fmt"
    "math"
)

// Integer mengembalikan nilai maksimum dari tipe uint64
func MyIntegerU64Max() uint64 {
    return math.MaxUint64
}

func main() {
    // Memanggil fungsi Integer dan mencetak nilainya
    // 18446744073709551615
    fmt.Println("Nilai maksimum uint64:", MyIntegerU64Max())
}
