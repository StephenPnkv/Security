
package main

import (
  "fmt"
  "os"
  //"os/exec"
  "log"
  "strconv"
)

func main(){
  if len(os.Args) != 2{
    log.Fatal("Usage: kill <PID>")
  }
  pid, _ := strconv.Atoi(os.Args[1])
  process, err := os.FindProcess(pid)
  if err != nil{
    log.Fatal(err)
  }
  fmt.Println("Killing process ", process.Pid)
  process.Kill()
}
