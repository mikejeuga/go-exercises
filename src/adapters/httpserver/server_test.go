package httpserver_test

import (
	"github.com/matryer/is"
	"github.com/mikejeuga/go-exercises/src/adapters/httpserver"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
    t.Parallel()
	srv := httpserver.NewServer()
	for _, tt := range []struct {
		name, expected string
		handler http.HandlerFunc
		res *httptest.ResponseRecorder
		req *http.Request


	}{
		{
			name: "the MATH server is healthy",
			handler: srv.Home,
			res: httptest.NewRecorder(),
			req: httptest.NewRequest(http.MethodGet,"/", nil),
			expected: "The math server is healthy",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			tt.handler.ServeHTTP(tt.res, tt.req)
			is.Equal(tt.res.Code, http.StatusOK)
			is.Equal(tt.res.Body.String(), tt.expected)

		})
	}
 }

func TestAdd_Handler(t *testing.T) {
	t.Skip()
    t.Parallel()
	srv := httpserver.NewServer()
	for _, tt := range []struct {
		name, expected string
		handler http.HandlerFunc
		res *httptest.ResponseRecorder
		req *http.Request
	}{
		{
			name: "The math handler shows the addition of the numbers in the query",
			handler: srv.Add,
			res: httptest.NewRecorder(),
			req: httptest.NewRequest(http.MethodPost, "/math", nil),
			expected: "Total: 0",
		},{
			name: "The math handler shows the addition of the numbers in the query",
			handler: srv.Add,
			res: httptest.NewRecorder(),
			req: httptest.NewRequest(http.MethodPost, "/math?num=4&num=5", nil),
			expected: "Total: 9",
		},{
			name: "The math handler shows the addition of the numbers in the query",
			handler: srv.Add,
			res: httptest.NewRecorder(),
			req: httptest.NewRequest(http.MethodPost, "/math?num=4&num=5&num=32", nil),
			expected: "Total: 41",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			tt.handler.ServeHTTP(tt.res, tt.req)
			is.Equal(tt.res.Code, http.StatusOK)
			is.Equal(tt.res.Body.String(), tt.expected)

		})
	}
 }