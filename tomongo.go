package main

import (
    "context"
    "fmt"
    "log"
    "bufio"
    "os"
//    "strconv"
    "strings"
    "reflect"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
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

    // Create client
    client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
    if err != nil {
      log.Fatal(err)
    }

    // Create connect
    err = client.Connect(context.TODO())
    if err != nil {
      log.Fatal(err)
    }

    // Check the connection
    err = client.Ping(context.TODO(), nil)
    if err != nil {
      log.Fatal(err)
    }

    fmt.Println("Connected to MongoDB!")

    collection := client.Database("test").Collection("structs")

    fmt.Println(reflect.TypeOf(collection))

    file, err := os.Open("tail.txt")
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
//        sput(s)
        root := strings.Split(s[0], " | ")
	doc := Words{0, root[0], []Dicts{}}
	for _, str := range s {
	  strsplit := strings.Split(strings.TrimLeft(str, " "), " | ")
	  doc.Types = append(doc.Types, Dicts{strsplit[0], strsplit[1]})
	  if err != nil {
	    log.Fatal(err)
	  }
        s = []string{}
        }
	_ , err = collection.InsertOne(context.TODO(), doc)
	if err != nil {
          log.Fatal(err)
        }
      }
    }
/*    if len(s) > 0 {
      sput(s)
      s = []string{}
    }
*/
    err = client.Disconnect(context.TODO())
    if err != nil {
      log.Fatal(err)
    }
    fmt.Println("Connection to MongoDB closed.")

}
/*
func sput(s1 []string){
  root := strings.Split(s1[0], " | ")
//  fmt.Printf("%s --- %d\n", s1[0], len(s1))
  for i, str := range s1 {
    strsplit := strings.Split(strings.TrimLeft(str, " "), " | ")
    fmt.Printf("%d --- %s --- %s --- %s\n", i, root[0], strsplit[0], strsplit[1])
  }
}
*/
