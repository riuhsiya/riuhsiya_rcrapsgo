package main

import (
    "fmt"
    "math"
)

// Integer mengembalikan nilai maksimum dari tipe int64
func MyIntegerI64Max() int64 {
    return math.MaxInt64
}

func main() {
    // Memanggil fungsi Integer dan mencetak nilainya
    // 9223372036854775807
    fmt.Println("Nilai maksimum int64:", MyIntegerI64Max())
}
