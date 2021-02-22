package main


import (
	"fmt"
	"os"
	"log"
	"net/http"
	"net/url"
	"github.com/PuerkitoBio/goquery"
	"time"
)

var (
	foundPaths []string 
	startingUrl *url.URL 
	timeout = time.Duration(8*time.Second)
)

func crawlUrl(path string){
	var targetUrl url.URL
	targetUrl.Scheme = startingUrl.Scheme
	targetUrl.Host = startingUrl.Host
	targetUrl.Path = path 
	httpClient := http.CLient{Timeout: timeout}
	response, err := httpClient.Get(targetUrl.String())
	if err != nil{
		return 
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil { return }

	doc.Find("a").Each(func(i int, s *goquery.Selection){
		href, exists := s.Attr("href")
		if !exists {return}

		parsedUrl, err := url.Parse(href)
		if err != nil {return}


		if urlIsInScope(parsedUrl){
			foundPaths = append(foundPaths, parsedUrl.Path)
			log.Println("Found new path to crawl: " + parsedUrl.String())
			crawlUrl(parsedUrl)
		}
	})
}

func main(){

}