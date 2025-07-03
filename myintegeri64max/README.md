myintegeri64max.go
myintegeri64max_test.go
```bash
go test -v
```
```bash
go run myintegeri64max_main.go
```

myintegeri64max.go
```bash
go build -gcflags="-N -l" -o myintegeri64max myintegeri64max.go
```

myintegeri64max.go
```golang
package myintegeri64max

import (
    "math"
)

// Integer mengembalikan nilai maksimum dari tipe int64
func MyIntegerI64Max() int64 {
    // 9223372036854775807
    return math.MaxInt64
}

```

myintegeri64max_test.go
```golang
package myintegeri64max

import (
    "math"
    "testing"
)

func TestMyIntegerI64Max(t *testing.T) {
    // Memanggil fungsi MyIntegerI64Max
    var result int64 = MyIntegerI64Max()

    // Nilai yang diharapkan adalah math.MaxInt64
    var expected int64 = int64(math.MaxInt64)

    // Membandingkan hasil dengan nilai yang diharapkan
    if result != expected {
        t.Errorf("MyIntegerI64Max() = %d; want %d", result, expected)
    }
}

```
