//ff:func feature=pack type=helper control=iteration dimension=1
//ff:what parallelMap — f over 0…n-1 on up to workers goroutines, results in order, the first error by index
package pack

import "sync"

// parallelMap runs f(i) for every i; the output order does not depend on scheduling.
func parallelMap[T any](n, workers int, f func(int) (T, error)) ([]T, error) {
	out, errs := make([]T, n), make([]error, n)
	idx := make(chan int, n)
	for i := 0; i < n; i++ {
		idx <- i
	}
	close(idx)
	var wg sync.WaitGroup
	for k := 0; k < min(workers, n); k++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range idx {
				out[i], errs[i] = f(i)
			}
		}()
	}
	wg.Wait()
	for _, err := range errs {
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}
