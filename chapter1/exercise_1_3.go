package main

import (
    "fmt"
    "time"
    "strings"
)

func main() {
    data := []string { "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k" }
    start := time.Now()
    experimentSize := 1234567 
   
    for i := 1; i < experimentSize; i++ {
	s := ""   
	for j := 1; j < len(data); j++ {
            s += data[j] + " "
    	}
    }

    fmt.Printf("%.2fs elapsed running with +\n", time.Since(start).Seconds())

    start = time.Now()
    
    for i := 1; i < experimentSize; i++ {
	strings.Join(data, " ")
    }
    
    fmt.Printf("%.2fs elapsed running with join\n", time.Since(start).Seconds())
}
