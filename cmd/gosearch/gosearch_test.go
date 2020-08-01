package gosearch

import (
	"os"
	"testing"
)

func TestMain(t *testing.M) {
	os.Exit(t.Run())
}

type table struct {
	using string
	want  string
}

// add more as needs get figured out
func TestBuildGoogleUrl(t *testing.T) {
	tables := []table{
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

	for _, test := range tables {
		got := buildGoogleUrl(test.using)
		if got != test.want {
			t.Errorf("\nGot \"%s\"\nWant \"%s\"\n", got, test.want)
			t.Fail()
		}
	}
}
