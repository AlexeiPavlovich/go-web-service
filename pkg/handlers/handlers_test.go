package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
)

func TestHandler(t *testing.T) {
	mux := Routs()

	ts := httptest.NewTLSServer(mux)
	defer ts.Close()
	resp, err := ts.Client().Get(ts.URL + "/json")
	if err != nil {
		t.Log(err)
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Error(resp.StatusCode)
	}
}

func Routs() http.Handler {
	/*mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.HomeHandler))
	mux.Get("/about", http.HandlerFunc(handlers.AboutHandler))
	return mux
	*/
	mux := chi.NewRouter()
	mux.Get("/json", Json)

	return mux

}
