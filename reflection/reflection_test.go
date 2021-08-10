package main

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
			"Struct with one string field",
			struct {
				Name string
			}{"Gamera"},
			[]string{"Gamera"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Gamera", "Rodan"},
			[]string{"Gamera", "Rodan"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				Age  int
			}{"Gamera", 1231},
			[]string{"Gamera"},
		},
		{
			"Nested fields",
			Person{
				"Gamera",
				Profile{
					1231,
					"Tokyo",
				},
			},
			[]string{"Gamera", "Tokyo"},
		},
		{
			"Pointers to things",
			&Person{
				"Godzilla",
				Profile{9999, "Kyoto"},
			},
			[]string{"Godzilla", "Kyoto"},
		},
		{
			"Slices",
			[]Profile{
				{33, "Tokyo"},
				{34, "Kyoto"},
			},
			[]string{"Tokyo", "Kyoto"},
		},
		{
			"Arrays",
			[2]Profile{
				{33, "Tokyo"},
				{34, "Kyoto"},
			},
			[]string{"Tokyo", "Kyoto"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
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

func assertContains(t testing.TB, haystack []string, needle string) {
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
