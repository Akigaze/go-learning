package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("Search a existing word", func(t *testing.T) {

		got, _ := dictionary.Search("test")
		wanted := "this is just a test"

		assertString(t, got, wanted)
	})

	t.Run("Search an unknown word", func(t *testing.T) {
		word := "unknown"
		_, err := dictionary.Search(word)

		assertError(t, err, ErrNotFound)
	})

}

func assertString(t *testing.T, got, wanted string) {
	t.Helper()

	if got != wanted {
		t.Errorf("expect '%s' but got '%s'", wanted, got)
	}

}

func assertError(t *testing.T, got, wanted error) {
	t.Helper()
	if got == nil {
		t.Fatal("expect to have an error but didn't")
	}

	if got != wanted {
		t.Errorf("expect '%s' but got '%s'", wanted, got)
	}
}
