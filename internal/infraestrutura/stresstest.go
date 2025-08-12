package infraestrutura

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func StressTest(url string, requests int, concurrency int) {
	results := make([]int, requests)
	var wg sync.WaitGroup

	start := time.Now()

	wg.Add(requests)
	ch := make(chan int, concurrency)

	for j := 0; j < requests; j++ {
		ch <- 1 // Controla a concorrência
		go func(j int) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				results[j] = 400
			} else {
				results[j] = resp.StatusCode
			}
			<-ch // Libera um slot de concorrência
		}(j)
	}

	wg.Wait()
	close(ch)

	elapsed := time.Since(start)

	fmt.Printf("Tempo total gasto: %v\n", elapsed)
	fmt.Printf("Total de requests realizados: %d\n", requests)

	statusCounts := make(map[int]int)
	for _, code := range results {
		statusCounts[code]++
	}

	for code, count := range statusCounts {
		fmt.Printf("Quantidade de status HTTP %d: %d\n", code, count)
	}

}
