package calculator

import "testing"

func TestCalculator(t *testing.T) {

	t.Run("Test add", func(t *testing.T) {
		want := 2 + 3
		got := Add(2, 3)

		if got != want {
			t.Errorf("want '%d' but got '%d'", want, got)
		}
	})

	t.Run("Test sub", func(t *testing.T) {
		want := 2 - 3
		got := Subtract(2, 3)

		if got != want {
			t.Errorf("want '%d' but got '%d'", want, got)
		}
	})

	t.Run("Test multi", func(t *testing.T) {
		want := float64(2 * 3)
		got := Multiply(2, 3)

		if got != want {
			t.Errorf("want '%f' but got '%f'", want, got)
		}
	})

	t.Run("Test division", func(t *testing.T) {
		want := float64(2 / 3)
		got := Divide(2, 3)

		if got != want {
			t.Errorf("want '%f' but got '%f'", want, got)
		}
	})
}
