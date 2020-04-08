package main

// <sitemapindex>
//    <sitemap>
//       <loc>http://www.washingtonpost.com/news-politics-sitemap.xml</loc>
//    </sitemap>
//    <sitemap>
//       <loc>http://www.washingtonpost.com/news-blogs-politics-sitemap.xml</loc>
//    </sitemap>
//    <sitemap>
//       <loc>http://www.washingtonpost.com/news-opinions-sitemap.xml</loc>
//    </sitemap>
// </sitemapindex>

import ("fmt"
        "net/http"
        "io/ioutil"
        "encoding/xml")

type SitemapIndex struct {
  Locations []Location `xml:"sitemap"`
}

type Location struct {
  Loc string `xml:"loc"`
}

func (L Location) String() string {
  return fmt.Sprintf(l.Loc)
}

func main() {
  resp, _ := http.Get("http://www.washingtonpost.com/news-sitemap-index.xml")
  bytes, _ := ioutil.ReadAll(resp.Body)
  resp.Body.Close()

  var s SitemapIndex
  xml.Unmarshal(bytes, &s)

  fmt.Println(s.Locations)
}
