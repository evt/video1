package server

import (
	"net/http"
	"time"
)

// HTTPClient is HTTP client with timeout set
var HTTPClient = &http.Client{
	Timeout: time.Second * 10,
}
