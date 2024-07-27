package services

import (
	"fmt"
	"time"
)

type StressTestsReporterService struct {
	Results     []int
	ElapsedTime time.Duration
}

func NewStressTestsReporterService(results []int, elapsedTime time.Duration) *StressTestsReporterService {
	return &StressTestsReporterService{
		Results:     results,
		ElapsedTime: elapsedTime,
	}
}

func (s *StressTestsReporterService) Run() {
	resultMap := map[int]int{}

	for _, status := range s.Results {
		resultMap[status] = resultMap[status] + 1
	}

	fmt.Printf("\n\nTotal runtime: %dms", s.ElapsedTime.Milliseconds())
	fmt.Printf("\nTotal requests made: %d", len(s.Results))

	for key, val := range resultMap {
		percentageForStatus := float64(val) / float64(len(s.Results)) * 100
		fmt.Printf("\nNumber of %d status requests: %d - %.2f%s\n", key, val, percentageForStatus, "%")
	}
}
