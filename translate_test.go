package translator

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestTranslator_Translate(t *testing.T) {
	srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/translate_a/single":
			resp := sentences{
				Sentences: []sentence{
					{Trans: "Hello World!", Orig: "你好，世界！", Backend: 1},
				},
				Src: "zh-CN",
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)
		default:
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte(`<html><script>tkk:'445921.1498498556'</script></html>`))
		}
	}))
	defer srv.Close()

	host := strings.TrimPrefix(srv.URL, "https://")
	c := Config{ServiceUrls: []string{host}}

	trans := New(c)
	trans.client.Transport = newAddHeaderTransport(srv.Client().Transport, map[string]string{
		"User-Agent": defaultUserAgent,
	})

	result, err := trans.Translate("你好，世界！", "auto", "en")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.Text != "Hello World!" {
		t.Fatalf("got %q, want %q", result.Text, "Hello World!")
	}
	if result.Src != "zh-CN" {
		t.Fatalf("got src %q, want %q", result.Src, "zh-CN")
	}
	if result.Dest != "en" {
		t.Fatalf("got dest %q, want %q", result.Dest, "en")
	}
	if result.Origin != "你好，世界！" {
		t.Fatalf("got origin %q, want %q", result.Origin, "你好，世界！")
	}
}
