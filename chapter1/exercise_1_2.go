package main

import (
    "fmt"
    "os"
)

func main() {
    for index := 1; index < len(os.Args); index++ {
        fmt.Printf("%v - %s\n", index, os.Args[index])
    }
}
