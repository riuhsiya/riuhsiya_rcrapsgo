package myintegeri64max

import (
    "math"
    "testing"
)

func TestMyIntegerI64Max(t *testing.T) {
    // Memanggil fungsi MyIntegerI64
    var result int64 = MyIntegerI64Max()

    // Nilai yang diharapkan adalah math.MaxInt64
    var expected int64 = int64(math.MaxInt64)

    // Membandingkan hasil dengan nilai yang diharapkan
    if result != expected {
        t.Errorf("MyIntegerI64Max() = %d; want %d", result, expected)
    }
}

