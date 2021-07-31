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

		assertError(t, err, ErrWordExist)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "another definition"
		err := dictionary.Update(word, newDefinition)

		assertNoError(t, err)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("Update new word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Update("test", "this is just a test")

		assertError(t, err, ErrWordNotExisted)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "this is just a test"}
	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err == nil {
		t.Errorf("expect word '%s' not found", word)
	}
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
