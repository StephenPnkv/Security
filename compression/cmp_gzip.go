
package main

import (
	"compress/gzip"
	"log"
	"os"
	"io/ioutil"
)

func main(){
	if len(os.Args) != 3{
		log.Fatalf("Usage: %s <src file> <compressed file>", os.Args[0])
	}
	
	dst := os.Args[2]

	//Open file to read into buffer
	srcFile, err := os.Open(os.Args[1])
	if err != nil{
		log.Fatalln(err)
	}
	//Read data
	data, err := ioutil.ReadAll(srcFile)
	if err != nil{
		log.Fatalln("Error reading data from file: ", err)
	}

	//Create compressed file
	gzipFile, err := os.Create(dst)
	if err != nil{
		log.Fatalln("Error creating compressed file.")
	}

	//Create gzip writer 
	gzipWriter := gzip.NewWriter(gzipFile)
	defer gzipWriter.Close()

	_, err = gzipWriter.Write(data)
	if err != nil{
		log.Fatalln("Error writing data to compressed file.")
	}
}