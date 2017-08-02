package main 

import (
	"fmt"
	"os"
	sr "github.com/httpreserve/simplerequest"
)

//var dnz = "https://api.digitalnz.org/v3/records.json?api_key=iW-eypBNx_yo1JzKqqRY&per_page=0&facets_per_page=500&page=1&facets=content_partner"
var dnz = "https://api.digitalnz.org/"

func main() {

	fmt.Println(sr.Version())

	fmt.Println("xxx")

	sreq, err := sr.Create("GET", dnz)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}	

	sresp, err := sreq.Do()
	if err != nil {
		fmt.Println(dnz)
		fmt.Println(err)		
		os.Exit(1)
	}

	fmt.Println(sresp.Data)

}