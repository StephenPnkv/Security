
package main

import (
  "fmt"
  "net"
  "log"
  "os"
)

func main(){
  if len(os.Args) != 2 {
      log.Fatal("No hostname argument provided.")
   }
   arg := os.Args[1]

  fmt.Println("Looking up IP for host: " + arg)
  ips, err := net.LookupHost(arg)
  if err != nil{
    log.Fatal(err)
  }
  for _, ip := range ips {
    fmt.Println(ip)
  }
}
