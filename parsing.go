package main

import (
    "fmt"
    "log"
    "bufio"
    "os"
//    "strconv"
    "strings"

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

    file, err := os.Open("head.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    s := []string{}
    for scanner.Scan() {
      sline := scanner.Text()
      if len(sline) > 1 {
	s = append(s, sline)
      } else {
	sput(s)
	s = []string{}
      }
    }
    if len(s) > 0 {
      sput(s)
      s = []string{}
    }

}

func sput(s1 []string){
  root := strings.Split(s1[0], " | ")
//  fmt.Printf("%s --- %d\n", s1[0], len(s1))
  for i, str := range s1 {
    strsplit := strings.Split(strings.TrimLeft(str, " "), " | ")
    fmt.Printf("%d --- %s --- %s --- %s\n", i, root[0], strsplit[0], strsplit[1])
  }
}

