package maps

import "testing"

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error ) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, key, value string) {
	t.Helper()

	got, err := dictionary.Search(key)

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertStrings(t, got, value)
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("with an existing key", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		expected := "this is just a test"
	
		assertStrings(t, got, expected)
	})

	t.Run("with a non existing key", func(t *testing.T) {
		_, err := dictionary.Search("nope")
		expected := "not found"

		if err == nil {
			t.Fatal("it was expected to get an error")
		}
	
		assertStrings(t, err.Error(), expected)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}
		err := dictionary.Add("test", "this is just a test")

		assertError(t, err, nil)
		
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func (t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}