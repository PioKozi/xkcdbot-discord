package main

import (
	"testing"
)

// func TestMain(t *testing.M) {
// 	os.Exit(t.Run())
// }

type Table struct {
	Using string
	Want  string
}

// add more as needs get figured out
func TestBuildGoogleUrl(t *testing.T) {
	tables := []Table{
		{
			"xkcd",
			"https://www.google.com/search?q=xkcd&num=1",
		},
		{
			"xkcd security",
			"https://www.google.com/search?q=xkcd+security&num=1",
		},
		{
			"some long text  that has    oble sp ace",
			"https://www.google.com/search?q=some+long+text+that+has+oble+sp+ace&num=1",
		},
	}

	for _, table := range tables {
		got := buildGoogleUrl(table.Using)
		if got != table.Want {
			t.Errorf("\nGot \"%s\"\nWant \"%s\"\n", got, table.Want)
			t.Fail()
		}
	}
}

// can't really make this have a failed state, because it's using a searchengine
// this test can only check if results are as could be expected/predictable

// as this function is also doing http query, it will take longer, so don't put
// too many tests in
func TestGoogleScrape(t *testing.T) {
	tables := []Table{
		{
			"site:xkcd.com AND inurl:https://xkcd.com/ security",
			"https://xkcd.com/538/",
		},
		{
			"site:xkcd.com AND inurl:https://xkcd.com/ dynamic entropy",
			"https://xkcd.com/2318/",
		},
	}

	for _, table := range tables {
		got, _ := GoogleScrape(table.Using)
		if got.Url != table.Want {
			t.Errorf("\nGot \"%s\"\nWant \"%s\"\n", got.Url, table.Want)
			t.Fail()
		}
	}
}
