package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
	"runtime"
	//"sync"
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

func GetFirstPage(url string, crawledPages map[string]int){
	//http getでresponse codeを見る
	fmt.Println("start crawl" + url)
	res, _ := http.Get(url)
	//defer res.Body.Close()
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
				}
			}
		})
	}else{
		fmt.Println(res.StatusCode)
	}
}

//func GoGetPages(urls []string) /*[]string*/{
//	pageurls := []string{}
//	var wg sync.WaitGroup
//	for _, url := range urls  {
//		wg.Add(1)
//		go func(url string){
//			defer wg.
//			Done()
//			fmt.Println(url)
//			pageurls = GetFirstPage(url)
//		}(url)
//	}
//	wg.Wait()
//	//return pageurls
//}

func isMatchDomain(url string) int{
    return strings.Index(url, "://hair.cm")
}

func main() {
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())

	url := "https://hair.cm/tu/article-27066/"
	crawledPages := map[string]int{}
	GetFirstPage(url, crawledPages)

	for url, _ := range crawledPages {
		GetFirstPage(url, crawledPages)
	}
	fmt.Println(crawledPages)
}
