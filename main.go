package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
)

func getCpuUsage(usage int) (string, string) {
	switch {
	case usage < 50:
		return Green, "HEALTHY"
	case usage <= 80:
		return Yellow, "WARNING"
	default:
		return Red, "CRITICAL"
	}
}

func main() {
	fmt.Println("======= CPU MONITOR =======")

	for i := 0; i < 5; i++ {
		cpu := rand.Intn(100)
		color, status := getCpuUsage(cpu)

		fmt.Printf("%s[%s] CPU Usage : %d%%%s\n", color, status, cpu, Reset)

		time.Sleep(time.Second)
	}

}
