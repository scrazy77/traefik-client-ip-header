package traefik_client_ip_header

import (
	"context"
	"net/http"
)

// Config configures the middleware.
type Config struct {
	HeaderName string `json:"headerName"`
}

// CreateConfig returns a config instance.
func CreateConfig() *Config {
	return &Config{}
}

type ClientIP struct {
	next       http.Handler
	name       string
	headername string
}

// New created a new Demo plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	var headername string = "CLIENT-IP"
	if len(config.HeaderName) > 0 {
		headername = config.HeaderName
	}
	return &ClientIP{
		next:       next,
		name:       name,
		headername: headername,
	}, nil
}

func (a *ClientIP) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	req.Header.Set(a.headername, req.RemoteAddr)
	a.next.ServeHTTP(rw, req)
}
