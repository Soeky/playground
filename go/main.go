package main

import (
	"fmt"
	"time"
)

func allocate() {
	data := make([]byte, 10*1024*1024) // 10 MB auf dem Heap
	fmt.Printf("Heap-Adresse: %p\n", &data[0])
	time.Sleep(5 * time.Second) // Zeit, um pprof zu starten
}

func main() {
	go allocate()
	time.Sleep(10 * time.Second)
}
