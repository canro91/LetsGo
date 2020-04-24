package reflection

import (
	"reflect"
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
			"Single property",
			struct {
				Name string
			}{"Alice"},
			[]string{"Alice"},
		},
		{
			"Two properties",
			struct {
				Name string
				City string
			}{"Alice", "Wonderland"},
			[]string{"Alice", "Wonderland"},
		},
		{
			"A property isn't string",
			struct {
				Name string
				Age  int
			}{"Alice", 16},
			[]string{"Alice"},
		},
		{
			"Nested fields",
			Person{
				"Alice",
				Profile{16, "Wonderland"},
			},
			[]string{"Alice", "Wonderland"},
		},
		{
			"Pointers to things",
			&Person{
				"Alice",
				Profile{16, "Wonderland"},
			},
			[]string{"Alice", "Wonderland"},
		},
		{
			"Slices",
			[]Profile{
				{1, "Wonderland"},
				{2, "Howards"},
			},
			[]string{"Wonderland", "Howards"},
		},
		{
			"Arrays",
			[2]Profile{
				{1, "Wonderland"},
				{2, "Howards"},
			},
			[]string{"Wonderland", "Howards"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var actual []string

			walk(test.Input, func(input string) {
				actual = append(actual, input)
			})

			if !reflect.DeepEqual(actual, test.ExpectedCalls) {
				t.Errorf("Expected %s, but was %s", test.ExpectedCalls, actual)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{34, "Katowice"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "Katowice"}
		}

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t *testing.T, haystack []string, needle string) {
	t.Helper()

	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
