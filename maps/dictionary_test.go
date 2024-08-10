package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add("test", "this is just a test")

		assertNoError(t, err)

		want := "this is just a test"
		got, err := dictionary.Search("test")

		assertNoError(t, err)
		assertStrings(t, got, want)
	})

	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		err := dictionary.Add("test", "something else")

		assertError(t, err, ErrExists)

		want := "this is just a test"
		got, err := dictionary.Search("test")

		assertNoError(t, err)
		assertStrings(t, got, want)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}

		err := dictionary.Update("test", "this is just a test")
		assertError(t, err, ErrDoesNotExist)
	})

	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		err := dictionary.Update("test", "updated")
		assertNoError(t, err)

		want := "updated"
		got, err := dictionary.Search("test")

		assertNoError(t, err)
		assertStrings(t, got, want)
	})
}

func TestDelete(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	dictionary.Delete("test")

	_, err := dictionary.Search("test")

	assertError(t, err, ErrNotFound)
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("wanted no error, but got %q", err)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("wanted an error, but didn't get one")
	}

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
