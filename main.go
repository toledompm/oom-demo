package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

var chunks [][]byte
var totalUsedBytes int = 0

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: main <command>")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "overcommit":
		overcommit()
	case "oom":
		go func() {
			oom()
		}()
		for {
			time.Sleep(250 * time.Millisecond)
			fmt.Printf("Used %.2f GB\n", float64(totalUsedBytes)/1024/1024/1024)

			var m runtime.MemStats
			runtime.ReadMemStats(&m)
			fmt.Printf("HeapAlloc: %v bytes, Sys: %v bytes\n", m.HeapAlloc, m.Sys)
		}
	default:
		fmt.Println("Invalid command")
		os.Exit(1)
	}
}

func oom() {
	malloc_size := 1024 * 1024 * 256
	chunk := make([]byte, malloc_size)

	for i := 0; i < malloc_size; i++ {
		chunk[i] = byte('x')
	}

	for i := 0; i < 1000; i++ {
		tmp := make([]byte, malloc_size)
		copy(tmp, chunk)
		chunks = append(chunks, tmp)

		totalUsedBytes += malloc_size
	}

}

func overcommit() {
	var usedMemory [][]byte
	var totalUsedBytes int = 0
	for j := 0; j < 1000; j++ {
		memory := make([]byte, 1024*1024*512)
		usedMemory = append(usedMemory, memory)
		totalUsedBytes += 1024 * 1024 * 512

		fmt.Printf("Allocated %.2f GB\n", float64(totalUsedBytes)/1024/1024/1024)

		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		fmt.Printf("HeapAlloc: %v bytes, Sys: %v bytes\n", m.HeapAlloc, m.Sys)

		time.Sleep(1 * time.Second)
	}
}
