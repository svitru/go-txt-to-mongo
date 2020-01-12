package main

import (
    "fmt"
//    "strconv"
//    "strings"

)

type Dicts struct {
    Word string
    Sign string
}

type Words struct {
    Index int
    Root string
    Types []Dicts
}

func main() {
  b := Words{3, "boo", []Dicts{{"one", "two"}, {"three", "four"}}}
  fmt.Println(b)
}
