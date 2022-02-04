package httpserver_test

import (
	"bytes"
	"encoding/json"
	"github.com/matryer/is"
	"github.com/mikejeuga/go-exercises/src/adapters/httpserver"
	add "github.com/mikejeuga/go-exercises/src/domain"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestServer(t *testing.T) {
    t.Parallel()
	srv := httpserver.NewServer(add.Default)
	for _, tt := range []struct {
		name, expected string
		res *httptest.ResponseRecorder
		req *http.Request


	}{
		{
			name: "the MATH server is healthy",
			res: httptest.NewRecorder(),
			req: httptest.NewRequest(http.MethodGet,"/", nil),
			expected: "The math server is healthy",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			srv.Handler.ServeHTTP(tt.res, tt.req)
			is.Equal(tt.res.Code, http.StatusOK)
			is.Equal(tt.res.Body.String(), tt.expected)

		})
	}
 }

func TestAdd_Handler(t *testing.T) {
    t.Parallel()
	srv := httpserver.NewServer(add.Default)
	for _, tt := range []struct {
		name, expected, ctype string
		res *httptest.ResponseRecorder
		req *http.Request
	}{
		{
			name: "The math handler shows the addition of no numbers in the query",
			res: httptest.NewRecorder(),
			req: httptest.NewRequest(http.MethodPost, "/math", nil),
			expected: "Total: 0",
		},{
			name: "The math handler shows the addition of the 2 numbers in the query",
			res: httptest.NewRecorder(),
			req: httptest.NewRequest(http.MethodPost, "/math?num=4&num=5", nil),
			expected: "Total: 9",
		},{
			name: "The math handler shows the addition of the 3 numbers in the query",
			res: httptest.NewRecorder(),
			req: httptest.NewRequest(http.MethodPost, "/math?num=4&num=5&num=32", nil),
			expected: "Total: 41",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			srv.Handler.ServeHTTP(tt.res, tt.req)
			is.Equal(tt.res.Code, http.StatusOK)
			is.Equal(tt.res.Body.String(), tt.expected)
		})
	}
 }

func TestAdd(t *testing.T) {
	is := is.New(t)
	srv := httpserver.NewServer(add.Default)


	jsonifym, _ := json.Marshal(httpserver.MathPayload{Numbers: []string{"4", "5", "32"}})

	req := httptest.NewRequest(http.MethodPost, "/math", bytes.NewReader(jsonifym))
	req.Header.Set("Content-Type", "application/json")
	resp:= httptest.NewRecorder()


	srv.Handler.ServeHTTP(resp, req)
	is.Equal(resp.Code, http.StatusOK)
	is.Equal(resp.Body.String(), "Total: 41")

}

func TestAdd_URLencoded(t *testing.T) {
	is := is.New(t)
	srv := httpserver.NewServer(add.Default)

	form := url.Values{}
	form.Add("num", "4")
	form.Add("num", "5")
	form.Add("num", "32")

	req := httptest.NewRequest(http.MethodPost, "/math", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp:= httptest.NewRecorder()


	srv.Handler.ServeHTTP(resp, req)
	is.Equal(resp.Code, http.StatusOK)
	is.Equal(resp.Body.String(), "Total: 41")

}

