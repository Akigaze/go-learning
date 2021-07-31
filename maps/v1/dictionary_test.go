package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	got := dictionary.Search("test")
	wanted := "this is just a test"

	assertString(t, got, wanted)
}

func assertString(t *testing.T, got, wanted string) {
	t.Helper()

	if got != wanted {
		t.Errorf("expect '%s' but got '%s'", wanted, got)
	}

}
