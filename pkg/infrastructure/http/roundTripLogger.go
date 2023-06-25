package util

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

type RoundTripperLogger struct {
	Inner http.RoundTripper
}

func (r *RoundTripperLogger) RoundTrip(request *http.Request) (*http.Response, error) {
	dumpRequest, err := httputil.DumpRequest(request, true)
	if err != nil {
		fmt.Println("error when try to dump request", err.Error())
	}
	fmt.Println("sending request", string(dumpRequest))

	response, err := r.Inner.RoundTrip(request)
	if err != nil {
		fmt.Println("error in request", err.Error())
		return response, err
	}
	dumpResponse, err := httputil.DumpResponse(response, true)
	if err != nil {
		fmt.Println("error when try to dump response", err.Error())

	}
	fmt.Println("sending request", string(dumpResponse))
	return response, err

}
