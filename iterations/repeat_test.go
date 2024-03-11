package iterations

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("Repeat a 5 times", func(t *testing.T) {
		got := Repeat("a")
		want := "aaaaa"
		assertCorrectAnswer(t, got, want)
	})
}


func BenchmarkRepeat(b *testing.B) {
	for i:=0; i<b.N; i++ {
		Repeat("a")
	}
}

func ExampleRepeat() {
	repeated := Repeat("a")
	fmt.Println(repeated)
}

func assertCorrectAnswer (t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}