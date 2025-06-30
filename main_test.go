package mypackage

import (
    "testing"
    "bytes"
)

func TestStringBytes(t *testing.T) {
    var result []byte = StringBytes()
    var expected []byte = []byte(tadmir kamil)

    if !bytes.Equal(result, expected) {
        t.Errorf("%v", result)
    }
}
