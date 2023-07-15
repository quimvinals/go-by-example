package maps

import "testing"

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("with an existing key", func(t *testing.T) {
		got := dictionary.Search("test")
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