package calculator_test

import (
	"github.com/mocamocaland/tdd-in-go-sample/calculator"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// setup statements
	setup()
	// run the tests
	e := m.Run()

	teardown()
	// report the exit code
	os.Exit(e)
}

func setup() {
	log.Println("Setting up.")
}

func teardown() {
	log.Println("Tearing down.")
}

func init() {
	log.Println("Init setup.")
}

func TestAdd(t *testing.T) {
	//defer func() {
	//	log.Println("Deferred tearig down.")
	//}()
	// Arrange
	e := calculator.Engine{}

	actAssert := func(x, y, want float64) {

		// Act
		got := e.Add(x, y)

		// Assert
		if got != want {
			t.Errorf("Add(%.2f,%.2f) incorrect, got: %.2f, want:%.2f", x, y, got, want)
		}
	}
	// テストデータ、期待値、検証
	t.Run("posivite input", func(t *testing.T) {
		x, y := 2.5, 3.5
		want := 6.0
		actAssert(x, y, want)
	})

	t.Run("negative input", func(t *testing.T) {
		x, y := -2.5, -3.5
		want := -6.0
		actAssert(x, y, want)
	})
}

func BenchmarkAdd(b *testing.B) {
	e := calculator.Engine{}

	// run Add func b.N times
	for i := 0; i < b.N; i++ {
		e.Add(2, 3)
	}

}
