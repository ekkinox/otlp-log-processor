package internal_test

import (
	"sync"
	"testing"

	"github.com/ekkinox/otlp-log-processor/internal"
)

func TestStorage(t *testing.T) {
	t.Parallel()

	t.Run("test sequential storage", func(t *testing.T) {
		t.Parallel()

		str := internal.NewStorage()

		str.Store("foo")
		str.Store("bar")
		str.Store("foo")

		res := str.Dump()

		if res["foo"] != 2 {
			t.Errorf("expected foo to have 2 occurences")
		}
		if res["bar"] != 1 {
			t.Errorf("expected bar to have 1 occurences")
		}
	})

	t.Run("test concurrent storage", func(t *testing.T) {
		t.Parallel()

		str := internal.NewStorage()

		var wg sync.WaitGroup

		keys := []string{
			"foo",
			"bar",
			"baz",
		}

		expected := 100 / len(keys)

		for i := 0; i < 99; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()

				key := keys[i%len(keys)]
				str.Store(key)
			}(i)
		}

		wg.Wait()

		res := str.Dump()

		for _, key := range keys {
			if res[key] != int64(expected) {
				t.Errorf("expected key %s to have count %d, but got %d", key, expected, res[key])
			}
		}
	})
}
