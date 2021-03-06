package tusd_test

import (
	"net/http"
	"testing"

	. "github.com/tus/tusd"
)

func TestCORS(t *testing.T) {
	store := NewStoreComposer()
	store.UseCore(zeroStore{})
	handler, _ := NewHandler(Config{
		StoreComposer: store,
	})

	(&httpTest{
		Name:   "Preflight request",
		Method: "OPTIONS",
		ReqHeader: map[string]string{
			"Origin": "tus.io",
		},
		Code: http.StatusOK,
		ResHeader: map[string]string{
			"Access-Control-Allow-Headers": "Origin, X-Requested-With, Content-Type, Upload-Length, Upload-Offset, Tus-Resumable, Upload-Metadata",
			"Access-Control-Allow-Methods": "POST, GET, HEAD, PATCH, DELETE, OPTIONS",
			"Access-Control-Max-Age":       "86400",
			"Access-Control-Allow-Origin":  "tus.io",
		},
	}).Run(handler, t)

	(&httpTest{
		Name:   "Actual request",
		Method: "GET",
		ReqHeader: map[string]string{
			"Origin": "tus.io",
		},
		Code: http.StatusMethodNotAllowed,
		ResHeader: map[string]string{
			"Access-Control-Expose-Headers": "Upload-Offset, Location, Upload-Length, Tus-Version, Tus-Resumable, Tus-Max-Size, Tus-Extension, Upload-Metadata",
			"Access-Control-Allow-Origin":   "tus.io",
		},
	}).Run(handler, t)
}
