package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"os"
	"regexp"
)

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Printf("%s environment variable not set.", k)
	}
	return v
}

func line_scraping() (title string, err error) {
	scraping_url := mustGetenv("SCRAPING_URL")
	doc, err := goquery.NewDocument(scraping_url)
	if err != nil {
		return "", err
	}
	doc.Find(".NewsList_header").First().Each(func(n int, s *goquery.Selection) {
		title = s.Text()
	})
	re := regexp.MustCompile(`\s+`)
	title = re.ReplaceAllString(title, `\s`)
	return title, nil
}

// func fb_scraping() (title string, err error) {
// 	scraping_url := mustGetenv("SCRAPING_URL")
// 	doc, err := goquery.NewDocument(scraping_url)
// 	if err != nil {
// 		return "", err
// 	}
// 	doc.Find("._4-u3 ._588p").First().Each(func(n int, s *goquery.Selection) {
// 		title = s.Text()
// 	})
// 	re := regexp.MustCompile(`\s+`)
// 	title = re.ReplaceAllString(title, `\s`)
// 	return title, nil
// }
