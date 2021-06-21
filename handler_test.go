package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETBizDays(t *testing.T) {
	t.Run("it only accepts GET method", func(t *testing.T) {
		req, recorder := newTestRequest(http.MethodPost, "")

		bizdaysHandler(recorder, req)

		got := recorder.Result().StatusCode
		want := http.StatusMethodNotAllowed

		if got != want {
			t.Errorf("got = %d; want = %d", got, want)
		}
	})

	t.Run("requires 'from' query parameter to be present", func(t *testing.T) {
		req, recorder := newTestRequest(http.MethodGet, "api/bizdays?to=43298")

		bizdaysHandler(recorder, req)

		got := recorder.Result().StatusCode
		want := http.StatusBadRequest

		if got != want {
			t.Errorf("got = %d; want = %d", got, want)
		}
	})

	t.Run("requires 'to' query parameter to be present", func(t *testing.T) {
		req, recorder := newTestRequest(http.MethodGet, "api/bizdays?from=43298")

		bizdaysHandler(recorder, req)

		got := recorder.Result().StatusCode
		want := http.StatusBadRequest

		if got != want {
			t.Errorf("got = %d; want = %d", got, want)
		}
	})

	t.Run("accepts only RFC3339 formated dates", func(t *testing.T) {
		req, recorder := newTestRequest(http.MethodGet, "api/bizdays?from=4234&to=43298")

		bizdaysHandler(recorder, req)

		got := recorder.Result().StatusCode
		want := http.StatusBadRequest

		if got != want {
			t.Errorf("got = %d; want = %d", got, want)
		}
	})

	t.Run("returns number of business days between two dates", func(t *testing.T) {
		req, recorder := newTestRequest(http.MethodGet, "api/bizdays?from=2021-06-21T00:00:00Z&to=2021-06-25T00:00:00Z")

		bizdaysHandler(recorder, req)

		gotStatusCode := recorder.Result().StatusCode
		wantStatusCode := http.StatusOK

		if gotStatusCode != wantStatusCode {
			t.Fatalf("got = %d; want = %d; error = %s", gotStatusCode, wantStatusCode, recorder.Body.String())
		}

		gotDays := recorder.Body.String()
		wantDays := "4"

		if gotDays != wantDays {
			t.Errorf("got = %s; want = %s", gotDays, wantDays)
		}
	})
}

func newTestRequest(method, url string) (*http.Request, *httptest.ResponseRecorder) {
	req, _ := http.NewRequest(method, url, nil)
	recorder := httptest.NewRecorder()

	return req, recorder
}
