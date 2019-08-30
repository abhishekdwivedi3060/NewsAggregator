package main

import (
	"encoding/xml"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/prometheus/common/log"
)

type SitemapIndex struct {
	Locations []string `xml: "sitemap>loc"`
}

type NewsMap struct {
	Loc []string `xml: "sitemap>loc"`
}
type News struct {
	Title string
	News  []NewsMap
}

func main() {
	http.HandleFunc("/news", newsAggHandler)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	var siteMap SitemapIndex
	res, err := http.Get("https://timesofindia.indiatimes.com/travel/staticsitemap/htpoi/sitemap-index.xml")
	if err != nil {
		panic(err)
	}
	bytes, _ := ioutil.ReadAll(res.Body)
	_ := xml.Unmarshal(bytes, &siteMap)
	t := template.Must(template.ParseFiles("link {{.}}"))
	//err := t.Execute(w, user1)
	//if err != nil {
	//	panic(err)
	//}

}
