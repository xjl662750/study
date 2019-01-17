package main

import "fmt"
import "strconv"
var t int

func init() {
    var err error
    t, err = strconv.Atoi("2")
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Println("init:", t)
}

func main() {
    fmt.Println("main:", t)
}