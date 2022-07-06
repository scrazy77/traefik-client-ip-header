package traefik-client-ip-header

import (
	"context"
	"net/http"
	"testing"

	plugin "github.com/scrazy77/traefik-client-ip-header"
)

func TestClientIP(t *testing.T) {
	cfg := plugin.CreateConfig()
	cfg.HeaderName = "CLIENT-IP"
	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := plugin.New(ctx, next, cfg, "traefik-client-ip-header")
	if err != nil {
		t.Fatal(err)
	}

}
