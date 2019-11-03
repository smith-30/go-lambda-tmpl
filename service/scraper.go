package service

import (
	"bufio"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func ExampleScrape() {
	// Request the HTML page.
	res, err := http.Get("https://www.vlcm.net/rc/pc/index.php?action_CRA01_01do=true&cid=00131")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	// shift-JIS to utf
	utfBody := transform.NewReader(bufio.NewReader(res.Body), japanese.ShiftJIS.NewDecoder())
	doc, err := goquery.NewDocumentFromReader(utfBody)
	if err != nil {
		log.Fatal(err)
	}

	txt, _ := doc.Html()
	fmt.Printf("%v\n", txt)

	// Find the review items
	doc.Find(".class0006").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		a := s.Find("a")
		band, ok := a.Attr("href")
		if ok {
			fmt.Printf("Review %d: %s\n", i, band)
			fmt.Printf("Review %d: %s %s\n", i, band, a.Text())
		}
		// title := a.HasClass()
		// fmt.Printf("%v\n", s.Text())

	})
}
