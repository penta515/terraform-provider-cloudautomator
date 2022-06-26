package client

import (
	"net/http"
	"testing"
)

func TestActive_ActivateJob(t *testing.T) {
	setup()
	defer server.Close()

	mux.HandleFunc("/jobs/1000/active", func(w http.ResponseWriter, r *http.Request) {
		testHttpMethod(t, r, "POST")
		_, _ = w.Write([]byte(`
		{
			"data": {
				"id": "1000",
				"type": "trigger_jobs",
				"attributes": {
					"active": true
				}
			}
		}`))
	})

	c, _ := New("example-token", WithAPIEndpoint(server.URL))

	a, _, err := c.ActivateJob("1000")

	active := true
	expect := &Active{
		Id:     "1000",
		Active: &active,
	}

	if err != nil {
		t.Fatal(err)
	}
	testEqual(t, expect, a)
}

func TestActive_DeactivateJob(t *testing.T) {
	setup()
	defer server.Close()

	mux.HandleFunc("/jobs/1000/active", func(w http.ResponseWriter, r *http.Request) {
		testHttpMethod(t, r, "DELETE")
		_, _ = w.Write([]byte(`
		{
			"data": {
				"id": "1000",
				"type": "trigger_jobs",
				"attributes": {
					"active": false
				}
			}
		}`))
	})

	c, _ := New("example-token", WithAPIEndpoint(server.URL))

	a, _, err := c.DeactivateJob("1000")

	active := false
	expect := &Active{
		Id:     "1000",
		Active: &active,
	}

	if err != nil {
		t.Fatal(err)
	}
	testEqual(t, expect, a)
}
