package timeutil

import (
	"context"
	"math/rand"
	"time"
)

// Sleep is an interruptable sleep which will return either if the context is
// done or if the duration has expired.
// If maxJitter is set to a non-zero value, a random jitter of [0,maxJitter is
// introduced to the sleep duration.
// This function intentionally does not return any indication of whether the context
// was cancelled because that may misleadingly be used in the caller to avoid race conditions
// but in most cases it would just be an illusion of protection against race conditions.
// Instead, the onus is on the developer to handle race conditions in the caller
// without relying on a return status from this function.
func Sleep(ctx context.Context, duration time.Duration, maxJitter time.Duration) {

	var j time.Duration

	if maxJitter != 0 {
		j = time.Duration(rand.Int63n(int64(maxJitter)))
	}

	select {
	case <-ctx.Done():
		return
	case <-time.After(duration + j):
		return
	}

}
