package handlers

import (
	"fmt"
	"net/http"
	"sync"

	"../helpers"
)

// Sum Route to sum all values in the matrix
func Sum(responseWriter http.ResponseWriter, request *http.Request) {
	matrix, _ := helpers.GetMatrix(helpers.SUM, responseWriter, request)
	count := 0
	wg := sync.WaitGroup{}

	for n := range matrix {
		wg.Add(1)

		go func(n int) {
			defer wg.Done()
			count += n
		}(n)
	}

	wg.Wait()

	fmt.Fprint(responseWriter, count)
}