package integers

import "testing"

func TestAdd(t *testing.T) {
	t.Run("Add two number", func(t *testing.T) {
		got := Add(2, 2)
		want := 4
		assertCorrectAnswer(t, got, want)
	})
}

func assertCorrectAnswer (t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}