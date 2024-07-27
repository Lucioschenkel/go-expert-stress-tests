package validation

import "errors"

var ErrConcurrencyGreaterThanRequests = errors.New("concurrency cannot be greater than the number of requests")
var ErrZeroConcurrencyOrRequests = errors.New("concurrency and requests cannot be less than 1")
var ErrUrlIsEmpty = errors.New("url cannot be empty")

func ValidateServiceArgs(url string, concurrency, requests int) (bool, error) {
	if concurrency > requests {
		return false, ErrConcurrencyGreaterThanRequests
	}

	if concurrency < 1 || requests < 1 {
		return false, ErrZeroConcurrencyOrRequests
	}

	if url == "" {
		return false, ErrUrlIsEmpty
	}

	return true, nil
}
