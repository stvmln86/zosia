// Package test implements unit testing data and functions.
package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stvmln86/zosia/zosia/tools/sqls"
)

// MockData is additional database data for unit testing.
const MockData = `
	insert into Users (addr, uuid) values ('1.1.1.1', '1111');
	insert into Pairs (user, name, body) values (1, 'alpha', 'Alpha.');
	insert into Pairs (user, name, body) values (1, 'bravo', 'Bravo.');
`

// DB returns a new in-memory database populated with MockData.
func DB(t *testing.T) *sqlx.DB {
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}

	if _, err := db.Exec(sqls.Pragma + sqls.Schema + MockData); err != nil {
		t.Fatal(err)
	}

	return db
}

// Get returns a value from a database query.
func Get(t *testing.T, db *sqlx.DB, code string, elems ...any) any {
	var data any
	if err := db.Get(&data, code, elems...); err != nil {
		t.Fatal(err)
	}

	return data
}

// Request returns a new mock Request with a body and mux values.
func Request(mthd, path, body string, pairs ...string) *http.Request {
	r := httptest.NewRequest(mthd, path, bytes.NewBufferString(body))
	for _, pair := range pairs {
		if name, data, ok := strings.Cut(pair, "="); ok {
			r.SetPathValue(name, data)
		}
	}

	return r
}

// Response returns a ResponseRecorder's status code and JSON response.
func Response(w *httptest.ResponseRecorder) (int, map[string]any, error) {
	var data map[string]any
	rslt := w.Result()
	err := json.NewDecoder(rslt.Body).Decode(&data)
	return rslt.StatusCode, data, err
}

// Set executes a database query.
func Set(t *testing.T, db *sqlx.DB, code string, elems ...any) {
	if _, err := db.Exec(code, elems...); err != nil {
		t.Fatal(err)
	}
}
