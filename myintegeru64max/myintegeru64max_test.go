package myintegeru64max

import (
    "math"
    "testing"
)

func TestMyIntegerU64Max(t *testing.T) {
    // Memanggil fungsi MyIntegerU64Max()
    var result uint64 = MyIntegerU64Max()

    // Nilai yang diharapkan adalah math.MaxUint64
    var expected uint64 = uint64(math.MaxUint64)

    // Membandingkan hasil dengan nilai yang diharapkan
    if result != expected {
        t.Errorf("MyIntegerU64Max() = %d; want %d", result, expected)
    }
}
