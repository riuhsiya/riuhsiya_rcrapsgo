### git --force
rm -rf .git && git init && git remote add origin git@github.com:riuhsiya/riuhsiya_rcrapsgo.git && git config --global user.name "riuhsiya" && git config --global user.email "riuhsiya@gmail.com" && git checkout -b main && git add . && git commit -m "init" && git push origin main --force
```bash
rm -rf .git && git init && git remote add origin git@github.com:riuhsiya/riuhsiya_rcrapsgo.git && git config --global user.name "riuhsiya" && git config --global user.email "riuhsiya@gmail.com" && git checkout -b main && git add . && git commit -m "init" && git push origin main --force
```

### git pull
rm -rf riuhsiya_rcrapsgo && mkdir riuhsiya_rcrapsgo && cd riuhsiya_rcrapsgo && git init && git remote add origin git@github.com:riuhsiya/riuhsiya_rcrapsgo.git && git config --global user.name "riuhsiya" && git config --global user.email "riuhsiya@gmail.com" && git pull origin main --allow-unrelated-histories && git branch -M main
```bash
rm -rf riuhsiya_rcrapsgo && mkdir riuhsiya_rcrapsgo && cd riuhsiya_rcrapsgo && git init && git remote add origin git@github.com:riuhsiya/riuhsiya_rcrapsgo.git && git config --global user.name "riuhsiya" && git config --global user.email "riuhsiya@gmail.com" && git pull origin main --allow-unrelated-histories && git branch -M main
```

### git --force-with-lease
git add . && git commit -m "up" --amend && git push -u origin main --force-with-lease
```bash
git add . && git commit -m "up" --amend && git push -u origin main --force-with-lease
```

```bash
git --no-pager diff README.md
```
```bash
git --no-pager diff --cached
```
```bash
git config --global core.pager ""
```
```bash
go mod init first_program
```
```bash
go test -v
```

String
```bash
package mypackage

// Fungsi yang akan di-test
func String() string {
    return "Hello"
}
```

TestString
```bash
package mypackage

import (
    "testing"
)

// Fungsi test
func TestString(t *testing.T) {
    var expected string = "Hello"
    var result string = String()

    // Bandingkan hasil dengan yang diharapkan
    if result != expected {
        t.Errorf("Expected %s, but got %s", expected, result)
    }
}
```

StringByte
```bash
package mypackage

// Fungsi yang akan di-test
func StringByte() []byte {
    return []byte("Hello")
}
```

TestStringByte
```bash
package mypackage

import (
    "testing"
)

// Fungsi test
func TestStringByte(t *testing.T) {
    var expected []byte = []byte("Hello")
    var result []byte = StringByte()

    // Bandingkan hasil dengan yang diharapkan
    if string(result) != string(expected) {
        t.Errorf("Expected %s, but got %s", expected, result)
    }
}
```

FloatTolerance
```bash
package mypackage

func FloatTolerance() float64 {
    return 3.14
}
```

TestFloatTolerance
```bash
package mypackage

import (
    "testing"
    "math"
)

func TestFloatTolerance(t *testing.T) {
    // memanggil fungsi yang akan di-test
    var result float64 = FloatTolerance()

    // Nilai yang diharapkan
    var expected float64 = 3.14

    // Toleransi yang diizinkan
    var tolerance float64 = 0.0001

    // Membandingkan hasil dengan nilai yang diharapkan dengan toleransi
    if math.Abs(result-expected) > tolerance {
        t.Errorf("Expected %f, got %f", expected, result)
    }
}
```

raceConditions
```golang
var counter int
for i := 0; i < 1000; i++ {
    go func() {
        counter++
    }()
}
```

raceConditions with sync.Mutex
```golang
package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter)
}

```

raceConditions with sync/atomic
```golang
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter)
}

```

raceConditions with sync/atomic
```golang
var counter int32

for i := 0; i < 1000; i++ {
    go func() {
        atomic.AddInt32(&counter, 1)
    }()
}

```

raceConditions with Channel
```golang
package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var wg sync.WaitGroup
	ch := make(chan int, 1)
	ch <- counter

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			current := <-ch
			current++
			ch <- current
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", <-ch)
}

```

TestRaceCondition
```golang
package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestCounterWithMutex(t *testing.T) {
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	if counter != 1000 {
		t.Errorf("Expected counter to be 1000, got %d", counter)
	}
}

func TestCounterWithAtomic(t *testing.T) {
	var counter int64
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}

	wg.Wait()
	if counter != 1000 {
		t.Errorf("Expected counter to be 1000, got %d", counter)
	}
}

func TestCounterWithChannel(t *testing.T) {
	var counter int
	var wg sync.WaitGroup
	ch := make(chan int, 1)
	ch <- counter

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			current := <-ch
			current++
			ch <- current
		}()
	}

	wg.Wait()
	if finalCounter := <-ch; finalCounter != 1000 {
		t.Errorf("Expected counter to be 1000, got %d", finalCounter)
	}
}

```
