//This program uses a tor proxy server to forward web requests 
//to a given domain. The output is printed and saved to a text file.
//Stephen Penkov 12/20/20

package main

import (
  "net/http"
  "net/url"
  "log"
  "os"
  "io/ioutil"
  "time"
)

func main(){
  if len(os.Args) != 2{
    log.Fatal("Usage: %s <target url>\nMake sure tor proxy server is running.\n", os.Args[0])
  }

  targetUrl := os.Args[1]
  torProxy := "socks5://localhost:9150"

  torProxyUrl,err := url.Parse(torProxy)
  if err != nil{
    log.Fatal("Error parsing proxy url: ", torProxy, ". ", err)
  }
  torTransport := &http.Transport{
    Proxy: http.ProxyURL(torProxyUrl),
  }
  client := &http.Client{
    Transport: torTransport,
    Timeout: time.Second * 5,
  }
  res, err := client.Get(targetUrl)
  if err != nil{
    log.Fatal("Error making GET request to ", targetUrl, ". ")
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil{
    log.Fatal("Error reading response body.")
  }

  htmlFile, err := os.Create("tor_req_output.txt")
  if err != nil{
    log.Fatal("Failed to create file.")
  }
  defer htmlFile.Close()
  bytesWritten := []byte(body)
  n, err := htmlFile.Write(bytesWritten)
  if err != nil{
    log.Fatal("Error writing to file.")
  }
  log.Printf("%d bytes written:\n%s", n, body)

}
