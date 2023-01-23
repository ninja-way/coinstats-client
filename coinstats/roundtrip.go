package coinstats

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type loggingRoundTripper struct {
	logger io.Writer
	next   http.RoundTripper
}

// RoundTrip print every request logs
func (l loggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	_, _ = fmt.Fprintf(l.logger, "%s : %s %s\n", time.Now().Format(time.RFC3339), r.Method, r.URL)
	return l.next.RoundTrip(r)
}
