package maps

import (
	"testing"
)

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

func TestAdd(t *testing.T) {
	t.Run("add a new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}
		err := dictionary.Add(word, definition)

		assertNoError(t, err)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("add an existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "another definition")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func assertNoError(t *testing.T, got error) {
	t.Helper()

	if got != nil {
		t.Fatal("got unexpected error:", got)
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("got unexpected error:", err)
	}

	assertString(t, got, definition)
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
