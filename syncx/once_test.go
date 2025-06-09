package syncx

import (
	"errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOnce_Do(t *testing.T) {
	testCases := []struct {
		name     string
		resource map[string]string
		makeErr  bool

		wantErr      error
		wantResource map[string]string
	}{
		{
			name:     "no error",
			resource: nil,
			makeErr:  false,

			wantErr:      nil,
			wantResource: map[string]string{"k": "v"},
		},
		{
			name:     "error",
			resource: nil,
			makeErr:  true,

			wantErr:      errors.New("error"),
			wantResource: nil,
		},
	}
	for _, tc := range testCases {
		once := Once{}
		t.Run(tc.name, func(t *testing.T) {
			var wg sync.WaitGroup
			var mu sync.Mutex
			var firstErr error

			f := func() error {
				if tc.makeErr {
					return errors.New("error")
				}

				mu.Lock()
				defer mu.Unlock()

				tc.resource = map[string]string{
					"k": "v",
				}

				return nil
			}

			for i := 0; i < 5; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					err := once.Do(f)
					mu.Lock()
					if err != nil && firstErr == nil {
						firstErr = err
					}
					mu.Unlock()
				}()
			}
			wg.Wait()

			assert.Equal(t, tc.wantErr, firstErr)
			assert.Equal(t, tc.wantResource, tc.resource)
		})
	}
}
