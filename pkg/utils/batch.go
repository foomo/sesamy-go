package utils

import (
	"context"
	"fmt"
)

// Batch reads from a channel and calls fn with a slice of batchSize.
func Batch[T any](ctx context.Context, ch <-chan T, batchSize int, fn func([]T)) {
	if batchSize <= 1 { // sanity check,
		for v := range ch {
			fmt.Println("<< 1")
			fn([]T{v})
		}
		return
	}

	// batchSize > 1
	var batch = make([]T, 0, batchSize)
	for {
		select {
		case <-ctx.Done():
			if len(batch) > 0 {
				fmt.Println("<< 2")
				fn(batch)
			}
			return
		case v, ok := <-ch:
			if !ok { // closed
				fmt.Println("<< 3")
				fn(batch)
				return
			}

			batch = append(batch, v)
			if len(batch) == batchSize { // full
				fmt.Println("<< 4")
				fn(batch)
				batch = make([]T, 0, batchSize) // reset
			}
		}
	}
}
