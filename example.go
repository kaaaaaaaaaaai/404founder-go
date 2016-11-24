package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
	//"sync"
	"sync"
)

type Page struct{
	url string
	statusCode int
}

//extend page
type ErrorPages struct {
	originUrl string
	page Page
}

func GetFirstPage(url string) []string{
	urls := []string{}
	crawledPages := map[string]int{}
	//http getでresponse codeを見る
	res, _ := http.Get(url)
	defer res.Body.Close()
	//200なら
	if res.StatusCode == 200{
		//中身のhrefを取得する。
		doc, _ := goquery.NewDocumentFromResponse(res)
		doc.Find("a").Each(func(_ int, s *goquery.Selection) {
			nextUrl, _ := s.Attr("href")
			if isMatchDomain(nextUrl) > 0 {
				if _, ok:= crawledPages[nextUrl]; ok{
				}else{
					crawledPages[nextUrl] = 200
					fmt.Println(url + " --->" + nextUrl)
					var wg sync.WaitGroup
					wg.Add(1)
					crawledPages[url] = 200
					go GetFirstPage(nextUrl)
				}
			}
		})
	}else{
		fmt.Println(res.StatusCode)
	}
	return urls
}
//
//func GoGetPages(urls []string) /*[]string*/{
//	pageurls := []string{}
//	var wg sync.WaitGroup
//	for _, url := range urls  {
//		wg.Add(1)
//		go func(url string){
//			defer wg.Done()
//			fmt.Println(url)
//			pageurls = GetFirstPage(url)
//		}(url)
//	}
//	wg.Wait()
//	//return pageurls
//}

func isMatchDomain(url string) int{
    return strings.Index(url, "://09362f6a.ngrok")
}

func main() {
	url := "https://09362f6a.ngrok.io"
	GetFirstPage(url)
}
