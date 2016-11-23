package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

type ErrorPages struct {
	ErrPage []string
}

func GetPage(url string) {
    var pages ErrorPages
	//http getでresponse codeを見る
	res, _ := http.Get(url)
	defer res.Body.Close()
	//200なら
	if res.StatusCode == 200{
		//中身のhrefを取得する。
		doc, _ := goquery.NewDocumentFromResponse(res)
		doc.Find("a").Each(func(_ int, s *goquery.Selection) {
			url, _ := s.Attr("href")
	        if isMatchDomain(url) > 0 {
	            fmt.Println(url)
	            _ = append(pages.ErrPage, string(url))
	        }
		})
	}else{
		fmt.Println(res.StatusCode)
	}
}

func isMatchDomain(url string) int{
    return strings.Index(url, "hair.cm")
}

func main() {
	var x ErrorPages
	url := "http://hair.cm"
	GetPage(url)
	fmt.Println(x)
}
