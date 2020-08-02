package cache

type searchPair struct {
	search string
	id     string
}

var (
	Cache = map[string]string{
		"security": "https://xkcd.com/538",
	}

	lastSearches = []searchPair{}
)

func UpdateLastSearches(search, id string) {
	lastSearches = append(lastSearches, searchPair{search, id})
}

func UpdateCache() {
	if len(lastSearches) == 5 {
		for _, pair := range lastSearches {
			Cache[pair.search] = pair.id
		}
		lastSearches = []searchPair{}
	}
}
