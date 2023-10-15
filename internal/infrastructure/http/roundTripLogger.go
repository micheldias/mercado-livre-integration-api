package util

import (
	contexthelper "mercado-livre-integration/internal/infrastructure/contextHelper"
	"net/http"
	"net/http/httputil"
)

type RoundTripperLogger struct {
	Inner http.RoundTripper
}

func (r *RoundTripperLogger) RoundTrip(request *http.Request) (*http.Response, error) {
	logger, ok := contexthelper.GetLogger(request.Context())

	if ok {
		dumpRequest, err := httputil.DumpRequest(request, true)
		if err != nil {
			logger.Warning("error when try to dump request", "error", err.Error())
		}
		logger.Info("sending request", "request", string(dumpRequest))
	}
	response, err := r.Inner.RoundTrip(request)
	if err != nil {
		if ok {
			logger.Error("error in request", "error", err.Error())
		}
		return response, err
	}
	if ok {
		dumpResponse, err := httputil.DumpResponse(response, true)
		if err != nil {
			logger.Warning("error when try to dump response", "error", err.Error())

		}
		logger.Info("received response", "response", string(dumpResponse))
	}

	return response, err

}
