myinteger64_main.go
```golang
go build -gcflags="-N -l" -o myinteger64_main myinteger64_main.go
```

myinteger64_main
```golang
package main

import (
    "fmt"
    "math"
)

// Integer mengembalikan nilai maksimum dari tipe int64
func MyInteger64() int64 {
    return math.MaxInt64
}

func main() {
    // Memanggil fungsi Integer dan mencetak nilainya
    // 9223372036854775807
    fmt.Println("Nilai maksimum int64:", MyInteger64())
}

```
