package main

import (
    "context"
    "fmt"
    "log"
    "bufio"
    "os"
    "strconv"
    "strings"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// You will be using this Trainer type later in the program
type Words struct {
    Word string
    Sign string
    Group int
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

    collection := client.Database("test").Collection("words")

    file, err := os.Open("data.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
      res := strings.Split(scanner.Text(), ";")
      i, _ := strconv.Atoi(res[2])
      arr := Words{res[0], res[1], i}
      _ , err := collection.InsertOne(context.TODO(), arr)
      if err != nil {
        log.Fatal(err)
      }
    }

    err = client.Disconnect(context.TODO())
    if err != nil {
      log.Fatal(err)
    }
    fmt.Println("Connection to MongoDB closed.")

}