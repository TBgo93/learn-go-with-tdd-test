package maps

import (
	"testing"
)

// map[tipo de key]tipo de valor

func TestSearch(t *testing.T) {
	assertStrings := func(t testing.TB, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	assertError := func(t testing.TB, got, want error) {
		t.Helper()

		if got != want {
			t.Errorf("got error %q want %q", got, want)
		}
	}

	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known key", func(t *testing.T) {

		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})
	t.Run("unknown key", func(t *testing.T) {

		_, err := dictionary.Search("unknown")
		want := ErrNotFound

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	assertError := func(t testing.TB, got, want error) {
		t.Helper()

		if got != want {
			t.Errorf("got error %q want %q", got, want)
		}
	}
	t.Run("new key", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "test"
		value := "this is just a test"

		err := dictionary.Add(key, value)
		assertError(t, err, nil)

		assertDefinition(t, dictionary, key, value)
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, key, value string) {
	t.Helper()
	got, err := dictionary.Search(key)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	if value != got {
		t.Errorf("got %q want %q", got, value)
	}
}

func TestUpdate(t *testing.T) {
	assertError := func(t testing.TB, got, want error) {
		t.Helper()

		if got != want {
			t.Errorf("got error %q want %q", got, want)
		}
	}
	t.Run("existing key", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		dictionary := Dictionary{key: value}
		newDefinition := "new definition"

		err := dictionary.Update(key, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, newDefinition)
	})
	t.Run("new key", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(key, value)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	key := "test"
	dictionary := Dictionary{key: "test definition"}

	dictionary.Delete(key)

	_, err := dictionary.Search(key)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", key)
	}
}
