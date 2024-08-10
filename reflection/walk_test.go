package reflection

import (
	"slices"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Dogwood"},
			[]string{"Dogwood"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Dogwood", "Atlanta"},
			[]string{"Dogwood", "Atlanta"},
		},
		{
			"struct with non-string field",
			struct {
				Name string
				Age  int
			}{"Dogwood", 4},
			[]string{"Dogwood"},
		},
		{
			"struct with nested fields",
			Person{"Dogwood", Profile{4, "Atlanta"}},
			[]string{"Dogwood", "Atlanta"},
		},
		{
			"pointers to things",
			&Person{"Dogwood", Profile{4, "Atlanta"}},
			[]string{"Dogwood", "Atlanta"},
		},
		{
			"slices",
			[]Profile{
				{4, "Atlanta"},
				{2, "New York City"},
			},
			[]string{"Atlanta", "New York City"},
		},
		{
			"arrays",
			[2]Profile{
				{4, "Atlanta"},
				{2, "New York City"},
			},
			[]string{"Atlanta", "New York City"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			assertSlice(t, got, test.ExpectedCalls)
		})
	}

	t.Run("with maps", func(t *testing.T) {
		m := map[string]string{
			"Name": "Dogwood",
			"City": "Atlanta",
		}

		var got []string
		walk(m, func(input string) { got = append(got, input) })

		if len(got) != len(m) {
			t.Errorf("Expected %v to have %d elements, but it had %d", got, len(m), len(got))
		}

		assertContains(t, got, "Dogwood")
		assertContains(t, got, "Atlanta")
	})

	t.Run("with channels", func(t *testing.T) {
		c := make(chan Profile)

		go func() {
			c <- Profile{4, "Atlanta"}
			c <- Profile{2, "New York City"}
			close(c)
		}()

		var got []string
		want := []string{"Atlanta", "New York City"}

		walk(c, func(input string) { got = append(got, input) })

		assertSlice(t, got, want)
	})

	t.Run("with function", func(t *testing.T) {
		f := func() (Profile, Profile) {
			return Profile{4, "Atlanta"}, Profile{2, "New York City"}
		}

		var got []string
		want := []string{"Atlanta", "New York City"}

		walk(f, func(input string) { got = append(got, input) })

		assertSlice(t, got, want)
	})
}

func assertSlice(t testing.TB, got []string, want []string) {
	t.Helper()

	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()

	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}
