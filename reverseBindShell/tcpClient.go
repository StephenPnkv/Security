
package main

import (
  "log"
  "net"
  "fmt"
  "bufio"
  "os"
)

var protocal = "tcp"
var remoteHostAddress = "localhost:4545"

func logError(err error){
  if err != nil{
    log.Fatal(err)
  }
}

func main(){
  conn, err := net.Dial(protocal, remoteHostAddress)
  logError(err)
  defer conn.Close()
  clientReader := bufio.NewReader(os.Stdin)
  serverReader := bufio.NewReader(conn)

  fmt.Println("Enter a message to send to remote server. Input :q to quit.")
  for{

    clientRequest, _ := clientReader.ReadString('\n')
    if clientRequest == ":q"{
      break
    }
    fmt.Fprintf(conn, clientRequest + "\n")
    responseMsg, _ := serverReader.ReadString('\n')
    fmt.Println("Message from the server: " + responseMsg)

  }
}
