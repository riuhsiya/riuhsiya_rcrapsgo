package main

import (
    "fmt"
    "math"
)

// Integer mengembalikan nilai maksimum dari tipe int64
func MyIntegerI64Min() int64 {
    return math.MinInt64
}

func main() {
    // Memanggil fungsi Integer dan mencetak nilainya
    // -9223372036854775808
    fmt.Println("Nilai maksimum int64:", MyIntegerI64Min())
}
