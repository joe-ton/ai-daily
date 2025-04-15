package scraper

import (
	"errors"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var (
	ErrEmptyValues = errors.New("Empty values found in struct")
)

type BookInfo struct {
	Title string
	UPC   string
	URL   string
}

func (b BookInfo) Find() ([]BookInfo, error) {
	if b.Title == "" || b.UPC == "" || b.URL == "" {
		return nil, ErrEmptyValues
	}
	resp, err := http.Get(b.URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	doc.Find(".product_pod").Each(func(i int, s *goquery))

}
