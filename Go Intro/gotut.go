package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Sitemapindex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
  Titles []string `xml:"url>news>title"`
  Keywords []string `xml:"url>news>Keywords"`
  Locations [string] `xml:sitemap>loc`
}

func main() {
  var s Sitemapindex
  var n News

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)

	xml.Unmarshal(bytes, &s)
	//fmt.Println(s.Locations)

	for _, Location := range s.Locations {
    resp, _ := http.Get("Location")
  	bytes, _ := ioutil.ReadAll(resp.Body)

  	xml.Unmarshal(bytes, &n)
	}
}
