

package main 

import (
	"fmt"
	"net/http"
	"os"
	"log"
)

func main(){
	if len(os.Args) != 2{
	fmt.Println(os.Args[0] + " - Perform an HTTP HEAD request to a URL")
      fmt.Println("Usage: " + os.Args[0] + " <url>")
      fmt.Println("Example: " + os.Args[0] + 
         " https://www.devdungeon.com")
      os.Exit(1)
	}

	url := os.Args[1]
	response, err := http.Head(url)
	if err != nil{
		log.Fatalln("Error fetching url", err)
	}

	for key, val := range response.Header{
		fmt.Printf("%s: %s \n", key, val[0])
	}
}