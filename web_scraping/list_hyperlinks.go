
package main

import (
  "fmt"
  "log"
  "os"
  "github.com/PuerkitoBio/goquery"
  "net/http"
)

func main(){
  if len(os.Args) != 2{
    fmt.Println("Find all links on a web page.")
    fmt.Println("Usage: " + os.Args[0] + " <url>")
    os.Exit(1)
  }

  url := os.Args[1]

  response, err := http.Get(url)
  if err != nil{
    log.Fatal("Error fetching URL. ", err)
  }

  doc, err := goquery.NewDocumentFromReader(response.Body)
  if err != nil{
    log.Fatal("Error loading HTTP response body")
  }

  doc.Find("a").Each(func(i int, s *goquery.Selection){
    href, exists := s.Attr("href")
    if exists{
      fmt.Println(href)
    }
  })
}
