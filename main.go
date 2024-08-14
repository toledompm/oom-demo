package main

import (
	"fmt"
	"os"
	"time"
)

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
		oom()
	default:
		fmt.Println("Invalid command")
		os.Exit(1)
	}
}

// Allocates and uses a lot of memory
func oom() {
	var usedMemory [][]byte
	var totalUsedBytes int = 0

	for j := 0; j < 1000; j++ {
		memory := make([]byte, 1024*1024*256)
		for i := range memory {
			memory[i] = 1
		}
		usedMemory = append(usedMemory, memory)
		totalUsedBytes += 1024 * 1024 * 256

		fmt.Printf("Used %.2f GB\n", float64(totalUsedBytes)/1024/1024/1024)
		time.Sleep(1 * time.Second)
	}
}

func overcommit() {
	var usedMemory [][]byte
	var totalUsedBytes int = 0
	for j := 0; j < 1000; j++ {
		memory := make([]byte, 1024*1024*512)
		usedMemory = append(usedMemory, memory)
		totalUsedBytes += 1024 * 1024 * 512
		fmt.Printf("Used %.2f GB\n", float64(totalUsedBytes)/1024/1024/1024)
		time.Sleep(1 * time.Second)
	}
}
