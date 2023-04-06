package main

import (
	"flag"
	"fmt"

	"p0hScan/match"
)

var url string = ""

func init() {
	flag.StringVar(&url, "url", "", "input a url")
}

func main() {
	flag.Parse()
	if url != "" {
		match.Match(url)
	} else {
		fmt.Println("please input a url")
	}
	/*
		if(is urls){
			urls_read, err1 := ioutil.ReadFile("urllist.txt")
			if err1 != nil {
				fmt.Println("Open file error:", err1)
			}
			url := string(urls_read)
			//match.Match(url)
			fmt.Println(url)
		}
	*/
}
