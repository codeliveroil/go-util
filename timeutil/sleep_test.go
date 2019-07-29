package timeutil

import (
	"context"
	"testing"
	"time"

	"github.com/codeliveroil/go-util/pkg/test"
)

func TestSleep(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())

	// Test regular sleep
	start := time.Now()
	d := 50 * time.Millisecond
	Sleep(ctx, d, 0)
	elapsed := time.Now().Sub(start)
	test.Compare(t, elapsed > d, true)

	// Test jitter - no real unflaky way to test this, but it'll report as exercised in code coverage
	start = time.Now()
	d = 50 * time.Millisecond
	Sleep(ctx, d, d)
	elapsed = time.Now().Sub(start)
	test.Compare(t, elapsed > d, true)

	// Test cancellation
	start = time.Now()
	d = 10 * time.Second
	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel()
	}()
	Sleep(ctx, d, 0)
	elapsed = time.Now().Sub(start)
	test.Compare(t, elapsed < d, true)

}
