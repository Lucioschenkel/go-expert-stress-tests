package services

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type StressTesterService struct {
	Url         string
	Requests    int
	Concurrency int
}

func NewStressTesterService(url string, requests, concurrency int) *StressTesterService {
	return &StressTesterService{
		Url:         url,
		Requests:    requests,
		Concurrency: concurrency,
	}
}

type StressTestResults struct {
	Results     []int
	ElapsedTime time.Duration
}

func (s *StressTesterService) Run() *StressTestResults {
	start := time.Now()
	fmt.Printf("Invoked with: \nURL: %s, Requests: %d, Concurrency: %d", s.Url, s.Requests, s.Concurrency)

	concurrent := make(chan int, s.Concurrency)
	wg := sync.WaitGroup{}
	results := make([]int, s.Requests)

	wg.Add(s.Requests)

	for i := 0; i < s.Requests; i++ {
		concurrent <- i
		fmt.Printf("\nExecuting request %d of %d", i+1, s.Requests)
		go s.makeRequest(&wg, concurrent, i, results)
	}

	wg.Wait()

	end := time.Now()

	elapsedTime := end.Sub(start)

	return &StressTestResults{
		Results:     results,
		ElapsedTime: elapsedTime,
	}
}

func (s *StressTesterService) makeRequest(wg *sync.WaitGroup, concurrencyBuffer <-chan int, id int, results []int) {
	resp, err := http.Get(s.Url)
	if err != nil {
		fmt.Printf("Error %v", err)
		wg.Done()
		<-concurrencyBuffer
		return
	}

	results[id] = resp.StatusCode

	<-concurrencyBuffer
	wg.Done()
}
