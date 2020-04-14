/*
@ Author:       Wang Xiaoqiang <https://github.com/rzjing>
@ File:         user_agent.go
@ Create Time:  2020/4/14 17:33
@ Software:     GoLand
*/

package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

const url = "http://www.useragentstring.com/pages/useragentstring.php?name="

// 获取指定浏览器下常用 User-Agent
func GetUserAgent(browser string) (UA []string) {
	if resp, err := http.Get(url + browser); err != nil {
		log.Println(err.Error())
	} else {
		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			if dom, err := goquery.NewDocumentFromReader(resp.Body); err != nil {
				log.Println(err.Error())
			} else {
				dom.Find("#liste").Find("ul").Each(func(i int, selection *goquery.Selection) {
					selection.Find("li").Each(func(i int, selection *goquery.Selection) {
						UA = append(UA, strings.TrimSpace(selection.Find("a").Text()))
					})
				})
			}
		}
	}
	return
}

func main() {
	var UA = make(map[string][]string)

	for _, b := range []string{"Chrome", "Firefox", "Internet+Explorer", "Opera", "Safari", "Edge", "Android+Webkit+Browser"} {
		UA[b] = GetUserAgent(b)
	}

	for browser, ua := range UA {
		for _, item := range ua {
			fmt.Println(browser, item)
		}
	}
}
