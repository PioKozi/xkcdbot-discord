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
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36")

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
	// hardocodinig shit because sometimes Google isn't very nice
	switch searchTerm {
	case "dynamic entropy":
		return "https://xkcd.com/2318/", nil
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
