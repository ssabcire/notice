package main

import (
	"cloud.google.com/go/datastore"
	"context"
	"github.com/PuerkitoBio/goquery"
	"log"
	"os"
	"regexp"
)

type Entity struct {
	Aaa string
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Printf("%s environment variable not set.", k)
	}
	return v
}

func line_scraping() (title string, err error) {
	scraping_url := mustGetenv("LINE_URL")
	doc, err := goquery.NewDocument(scraping_url)
	if err != nil {
		return "", err
	}
	doc.Find(".NewsList_header").First().Each(func(n int, s *goquery.Selection) {
		title = s.Text()
	})
	re := regexp.MustCompile(`\s+`)
	title = re.ReplaceAllString(title, " ")
	return title, nil
}


func initialize() (ctx context.Context, client *datastore.Client, err error) {
	ctx = context.Background()
	client, err = datastore.NewClient(ctx, mustGetenv("GOOGLE_CLOUD_PROJECT"))
	if err != nil {
		return nil, nil, err
	}
	return ctx, client, nil
}

func dsGet(kind string, name string) (string, error) {
	ctx, client, err := initialize()
	if err != nil {
		return "", err
	}
	k := datastore.NameKey(kind, name, nil)
	e := new(Entity)
	if err := client.Get(ctx, k, e); err != nil {
		return "", err
	}
	return e.Aaa, nil
}

func dsPut(kind string, name string, s string) error {
	ctx, client, err := initialize()
	if err != nil {
		return err
	}
	k := datastore.NameKey(kind, name, nil)
	e := &Entity{Aaa: s}
	if _, err := client.Put(ctx, k, e); err != nil {
		return err
	}
	return nil
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
