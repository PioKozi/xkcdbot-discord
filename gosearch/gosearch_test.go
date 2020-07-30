package gosearch

import (
	"os"
	"testing"
)

func TestMain(t *testing.M) {
	os.Exit(t.Run())
}

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

// TODO:  unit tests for GoogleScrape
