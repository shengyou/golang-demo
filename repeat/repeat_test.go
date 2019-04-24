package repeat

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("Should repeat 'a' 5 times", func(t *testing.T) {
		want := "aaaaa"
		got := Repeat(5, "a")

		if got != want {
			t.Errorf("want '%s' but got '%s'", want, got)
		}
	})

	t.Run("Should repeat 'b' 3 times", func(t *testing.T) {
		want := "bbb"
		got := Repeat(3, "b")

		if got != want {
			t.Errorf("want '%s' but got '%s'", want, got)
		}
	})
}
