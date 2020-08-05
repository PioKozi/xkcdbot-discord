package search

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// returns url for the whole search
func buildGoogleUrl(searchTerm string) string {
	searchTerm = strings.Trim(searchTerm, " ")
	searchTerm = strings.Join(strings.Fields(searchTerm), " ")
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)
	return fmt.Sprintf("https://www.google.com/search?q=%s&num=1", searchTerm)
}

// returns the HTML for that search
func googleRequest(searchUrl string) (*http.Response, error) {
	baseClient := &http.Client{}

	req, _ := http.NewRequest("GET", searchUrl, nil)
	req.Header.Set("User-Agent", randomUserAgent())

	res, err := baseClient.Do(req)
	log.Println(res)
	if err != nil {
		return nil, err
	} else {
		return res, nil
	}
}

func googleResultParser(response *http.Response) (string, error) {

	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		return "", err
	}

	sel := doc.Find("div.g")
	item := sel.Eq(0)

	linkTag := item.Find("a")
	link, _ := linkTag.Attr("href")
	link = strings.Trim(link, " ")

	if link != "" && link != "#" {
		return link, err
	}

	return "", err
}

func GoogleScrape(searchTerm string) (string, error) {
	// hardcoding queries because sometimes Google isn't very nice (which in
	// this scenario is fair, really)
	log.Println(searchTerm)
	switch strings.Join(strings.Fields(searchTerm)[3:], " ") {
	case "dynamic entropy":
		return "https://xkcd.com/2318/", nil
	case "sudo":
		return "https://xkcd.com/149/", nil
	case "network":
		return "https://xkcd.com/350/", nil
	case "networking":
		return "https://xkcd.com/2259/", nil
	}
	googleUrl := buildGoogleUrl(searchTerm)
	res, err := googleRequest(googleUrl)
	if err != nil {
		return "", err
	}
	link, err := googleResultParser(res)
	if err != nil {
		return "", err
	} else {
		return link, nil
	}
}
