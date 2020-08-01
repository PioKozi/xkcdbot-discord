package common

import (
	"testing"
)

type table struct {
	using string
	want  string
}

func TestCleanInput(t *testing.T) {
	tables := []table{
		{
			"this is the standard case",
			"this is the standard case",
		},
		{
			"THIS is ThE UpperCase Case",
			"this is the uppercase case",
		},
		{
			"     there   are  too many   spaces",
			"there are too many spaces",
		},
		{
			".xkcd    Abc    Def gh",
			"abc def gh",
		},
	}

	for _, test := range tables {
		got := PrepareInput(test.using, ".xkcd")
		if got != test.want {
			t.Errorf("\nGot: \"%s\"\nWant: \"%s\"\n", got, test.want)
			t.Fail()
		}
	}
}

func TestUpdateStringsMap(t *testing.T) {
	oldmap := map[string]string{
		"key": "value",
		"foo": "foo",
	}
	newmap := map[string]string{
		"foo":   "bar",
		"hello": "world",
	}

	wantmap := map[string]string{
		"key":   "value",
		"foo":   "bar",
		"hello": "world",
	}

	UpdateStringsMap(oldmap, newmap)

	if stringMapsDiffer(oldmap, wantmap) {
		t.Error("Got: ")
		t.Error(oldmap)
		t.Error("Want: ")
		t.Error(wantmap)
	}
}

func stringMapsDiffer(map1, map2 map[string]string) bool {
	for key, val1 := range map1 {
		if val2, exists := map2[key]; !exists || val1 != val2 {
			return true
		}
	}
	return false
}
