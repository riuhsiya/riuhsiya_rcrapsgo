var counter int
for i := 0; i < 1000; i++ {
    go func() {
        counter++
    }()
}
