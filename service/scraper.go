package service

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func ListScrape() {
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

	// Find the review items
	doc.Find(".class0006").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		onmouseover, ok := s.Attr("onmouseover")
		if ok {
			if strings.Contains(onmouseover, "個人フットサル(中級)") {
				a := s.Find("a")
				href, ok := a.Attr("href")
				if ok {
					if strings.Contains(href, ThisFriday(time.Now())) {
						fmt.Printf("%#v\n", "https://www.vlcm.net/rc/pc"+href[1:])
						message := CheckLesson("https://www.vlcm.net/rc/pc" + href[1:])
						fmt.Printf("%#v\n", message)
					}
				}
			}
		}
	})
}

func CheckLesson(url string) string {
	// Request the HTML page.
	res, err := http.Get(url)
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

	var message string
	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		s.Find("th").Each(func(i int, ss *goquery.Selection) {
			// fmt.Printf("%v\n", ss.Text())
			switch ss.Text() {
			case "開催日時", "イベント名称", "申込状況":
				// fmt.Printf("%#v\n", ss.Find("td").Text())
				message += ss.Find("td").Text()
			}
		})
	})
	return message
}

func ThisFriday(t time.Time) string {
	for {
		if weekday := t.Weekday(); weekday == time.Friday {
			break
		}
		t = t.Add(24 * time.Hour)
	}

	return t.Format("20060102")
}
