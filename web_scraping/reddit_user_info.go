
package main

import (
	"fmt"
	"os"
	"log"
	"encoding/json"
	"net/http"
	"time"
	"io/ioutil"
)

type redditUserJSONResponse struct{
	Data struct {
		Posts []struct {
			Data struct{
				Subreddit string `json:"subreddit"`
				Title string `json:"link_title"`
				PostedTime float32 `json:"created_utc"`
				Body string `json:"body"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

func printUsage(){
	fmt.Println(os.Args[0] + ` - Print recent Reddit posts by a user
	Usage: ` + os.Args[0] + ` <username>
	Example: ` + os.Args[0] + ` nanodano`)

}


func main(){
	if len(os.Args) != 2{
		printUsage()
		os.Exit(1)
	}
	url := "https://www.reddit.com/user/" + os.Args[1] + ".json"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Error making HTTP request.", err)
	}

	defer response.Body.Close() 
	body, err := ioutil.ReadAll(response.Body)
	if err != nil{
		log.Fatal("Error reading response body.")
	}

	if len(reddit)
}








