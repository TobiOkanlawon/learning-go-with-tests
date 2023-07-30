package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("search for a word that exists", func(t *testing.T) {
		got, err := dictionary.Search("test")

		if err != nil {
			t.Fatal("got an error")
		}
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("check for a word that doesn't exist", func(t *testing.T) {
		_, err := dictionary.Search("non existent phrase")

		assertErrors(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("Adds a new word to the dictionary", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "new"
		definition := "never before seen"

		err := dictionary.Add("new", definition)

		assertNoError(t, err)

		_, err2 := dictionary.Search("new")

		assertNoError(t, err2)
		assertDefinition(t, dictionary, word, definition)

	})

	t.Run("It doesn't add words that already exist", func(t *testing.T) {
		word := "word"
		definition := "definition"
		newDictionary := Dictionary{word: definition}

		err := newDictionary.Add(word, definition)

		assertErrors(t, err, ErrWordExists)
		assertDefinition(t, newDictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("updates an existing word", func(t *testing.T) {
		word := "test"
		definition := "definition"
		newDefinition := "new definition"

		dictionary := Dictionary{word: definition}

		dictionary.Update(word, newDefinition)
		assertDefinition(t, dictionary, word, newDefinition)

	})

	t.Run("doesn't run on new words", func(t *testing.T) {
		word := "new word"
		definition := "new definition"

		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertErrors(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {

	t.Run("deletes an existing word", func(t *testing.T) {
		word := "word"
		definition := "definition"
		dictionary := Dictionary{word: definition}

		dictionary.Delete(word)

		_, want := dictionary.Search(word)

		if want != ErrNotFound {
			t.Errorf("Doesn't delete the word")
		}
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %s , but want %s", got, want)
	}
}

func assertErrors(t testing.TB, got, want error) {

	t.Helper()

	if got != want {
		t.Fatal("Expected an error")
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("Did not expect an error while searching for", word)
	}

	if got != definition {
		t.Errorf("%s is not equal to %s", got, definition)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatal("Should not be error", err)
	}
}
