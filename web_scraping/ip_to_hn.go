//This program takes an IP address as input and returns the name of the host

package main

import (
  "fmt"
  "net"
  "log"
  "os"
)

func main(){
  if len(os.Args) != 2 {
      log.Fatalf("No IP address argument provided.\nUsage: %s <IP address>", os.Args[0])
   }
   arg := os.Args[1]

  ip := net.ParseIP(arg)
  if ip == nil{
    log.Fatal(`IP (${ip}) is not valid.`)
  }

  fmt.Println("Looking up hostnames for ip: " + arg)
  hostnames, err := net.LookupAddr(ip.String())
  if err != nil{
    log.Fatal(err)
  }
  for _, hostnames := range hostnames {
    fmt.Println(hostnames)
  }
}
