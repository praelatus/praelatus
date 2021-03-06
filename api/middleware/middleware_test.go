// Copyright 2017 Mathew Robinson <chasinglogic@gmail.com>. All rights reserved.
// Use of this source code is governed by the AGPLv3 license that can be found in
// the LICENSE file.

package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockHandler struct{}

func (m mockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("test"))
}

func TestLogger(t *testing.T) {
	m := mockHandler{}
	log := Logger(m)

	r, e := http.NewRequest("GET", "/", nil)
	if e != nil {
		t.Fatal(e)
	}

	w := httptest.NewRecorder()

	log.ServeHTTP(w, r)
}
