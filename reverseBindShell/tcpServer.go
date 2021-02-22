//This is a simple TCP server that reads data from a client. 

package main

import (
  "bufio"
  "net"
  "log"
  "fmt"
  "errors"
  "strings"
)

var protocal = "tcp"
var listenAddress = "localhost:4545"

var (
  ErrCreatingServer = errors.New("Error creating listen server: ")
  ErrAcceptingConn = errors.New("Error accepting connection: ")
)

func logError(err error){
  if err != nil{
    log.Fatal(err)
  }
}

func handleConn(conn net.Conn){
  defer conn.Close()
  clientReader := bufio.NewReader(conn)
  fmt.Printf("Established a connection from %s.\n", conn.RemoteAddr().String())
  //Create a 4MB buffer to read data from client
  for {
    clientRequest, err := clientReader.ReadString('\n')
    clientRequest = strings.TrimSpace(clientRequest)
    logError(err)
    fmt.Println("[" + conn.RemoteAddr().String() + "] - " + clientRequest)

    responseString := fmt.Sprintf("Read %d bytes.\n", len(clientRequest))
    if _, err := conn.Write([]byte(responseString)); err != nil{
      log.Panicln("Error responding to server.")
    }
  }

}

func main(){
  listener, err := net.Listen(protocal, listenAddress)
  logError(err)
  defer listener.Close()

  log.Println("Listening for connections . . .")
  for {
    conn, err := listener.Accept()
    logError(err)
    //defer conn.Close()
    go handleConn(conn)
  }

}
