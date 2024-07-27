package main

import (
	"flag"

	"github.com/Lucioschenkel/stresser/internal/services"
	"github.com/Lucioschenkel/stresser/internal/validation"
)

var (
	url         *string
	requests    *int
	concurrency *int
)

func init() {
	url = flag.String("url", "", "The endpoint under test")
	requests = flag.Int("requests", 1, "The total number of requests")
	concurrency = flag.Int("concurrency", 1, "The maximum number of concurrent requests")
}

func main() {
	flag.Parse()

	inputIsValid, err := validation.ValidateServiceArgs(*url, *concurrency, *requests)
	if !inputIsValid {
		panic(err)
	}

	stressTester := services.NewStressTesterService(*url, *requests, *concurrency)

	output := stressTester.Run()

	stressTestsReporterService := services.NewStressTestsReporterService(output.Results, output.ElapsedTime)
	stressTestsReporterService.Run()
}
