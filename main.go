package main

import (
	"encoding/xml"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/prometheus/common/log"
)

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type NewsMap struct {
	Loc            []string `xml:"url>loc"`
	UpdateSchedule []string `xml:"url>changefreq"`
}
type News struct {
	Loc            string
	UpdateSchedule string
}
type NewsAggPAge struct {
	Title string
	News  map[int]News
}

func main() {
	http.HandleFunc("/news", newsAggHandler)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	var siteMap SitemapIndex
	var newMap NewsMap
	n1 := make(map[int]News)
	res, err := http.Get("https://timesofindia.indiatimes.com/travel/staticsitemap/htpoi/sitemap-index.xml")

	if err != nil {
		panic(err)
	}
	bytes, _ := ioutil.ReadAll(res.Body)
	//fmt.Fprint(w, string(bytes))
	xml.Unmarshal(bytes, &siteMap)
	//fmt.Fprintf(w,"%v",siteMap )

	for _, x := range siteMap.Locations {
		//fmt.Print(x)
		resp, err := http.Get(x)
		if err != nil {
			panic(err)
		}
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &newMap)
		//fmt.Fprint(w, string(bytes))

		for i, _ := range newMap.Loc {
			//fmt.Print("idhr aya")
			n1[i] = News{newMap.Loc[i], newMap.UpdateSchedule[i]}
		}
	}
	//fmt.Fprintf(w, "%v", n1)
	news := NewsAggPAge{"Awesome NewsAggPAge", n1}

	t := template.Must(template.ParseFiles("index.html"))
	err = t.Execute(w, news)
	if err != nil {
		panic(err)
	}

}
