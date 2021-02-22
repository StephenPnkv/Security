
package main

import (
  "fmt"
  "os"
  "log"
  "bufio"
  "net/http"
  "net/url"
  "strconv"
)

func checkIfUrlExists(baseUrl, filepath string, doneChannel chan bool){
  targetUrl, err := url.Parse(baseUrl)
  if err != nil{
    log.Println("Error parsing base url.", err)
  }

  targetUrl.Path = filepath

  res, err := http.Head(targetUrl.String())
  if err != nil{
    log.Println("Error fetching ", targetUrl.String())
  }
  if res.StatusCode == 200 {
    log.Println(targetUrl.String())
  }
  doneChannel <- true

}

func main(){
  if len(os.Args) != 4{
    fmt.Println("Usage: " + os.Args[0] + "<wordlist file> <url> <maxThreads>")
    os.Exit(1)
  }

  wListFilename := os.Args[1]
  baseUrl := os.Args[2]
  maxThreads, err := strconv.Atoi(os.Args[3])
  if err != nil{
    log.Fatal("Error converting maxThread value to an integer.", err)
  }

  //Keep track of active threads
  activeThreads := 0
  doneChannel := make(chan bool)
  //Open words file and begin reading
  wListFile, err := os.Open(wListFilename)
  if err != nil{
    log.Fatal("Error opening the words list file!")
  }

  //Scan list line by line
  scanner := bufio.NewScanner(wListFile)
  for scanner.Scan(){
    //Send request to server using specified filename
    go checkIfUrlExists(baseUrl, scanner.Text(), doneChannel)
    activeThreads++
    if activeThreads >= maxThreads{
      <-doneChannel
      activeThreads--
    }

    for activeThreads > 0 {
      <-doneChannel
      activeThreads--
    }

    if err := scanner.Err(); err != nil{
      log.Fatal("Error reading the words list file!", err)
    }
  }
}
