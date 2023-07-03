package test

import "errors"

var ErrAPIResponse = errors.New("couldn't get a response from the API, status code: 500")
var ErrRateNotFound = errors.New("failed to retrieve the rate for the bitcoin")