package main

import (
	"encoding/json"
	"fmt"
	sr "github.com/httpreserve/simplerequest"
	"os"
	"reflect"
)

//var dnz = "https://api.digitalnz.org/v3/records.json?api_key=iW-eypBNx_yo1JzKqqRY&per_page=0&facets_per_page=500&page=1&facets=content_partner"

var dnz_base = "https://api.digitalnz.org/v3/records"
var dnz_format = ".json?"
var dnz_api_key = "api_key=iW-eypBNx_yo1JzKqqRY"

var our_search = "&text=kaka&category=Images&geo_bbox=-41,174,-42,175&per_page=2&page=2"

func getdata(f map[string]interface{}) {

	for k, v := range f {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
			if k == "large_thumbnail_url" {
				fmt.Println("HERE IS A SMALL CLUE FOR WHAT NEXT!! :D HAVE FUN FRAN")
			}
		case int:
			//fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array.")
			for _, u := range vv {
				switch u.(type) {
				case map[string]interface{}:
					getdata(u.(map[string]interface{}))
				}
			}
		case map[string]interface{}:
			fmt.Println(k, "is an map")
			getdata(v.(map[string]interface{}))
		default:
			if v != nil {
				fmt.Println(k, "is of a type I don't know how to handle", reflect.TypeOf(v))
			}
		}
	}
}

func main() {

	fmt.Fprintf(os.Stderr, sr.Version()+"\n")

	dnz_combined := dnz_base + dnz_format + dnz_api_key + our_search

	// do the request (you can look up the documentation if you like too)
	// documentation: https://digitalnz.org/developers/api-docs-v3
	sreq, err := sr.Create("GET", dnz_combined)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// get the response...
	sresp, err := sreq.Do()
	if err != nil {
		//fmt.Println(dnz)
		fmt.Println(err)
		os.Exit(1)
	}

	// put the response data into here
	data := sresp.Data

	f := make(map[string]interface{})
	err = json.Unmarshal([]byte(data), &f)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error:", err, "\n")
	}

	getdata(f)
}
