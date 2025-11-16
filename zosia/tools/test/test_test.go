package test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestDB(t *testing.T) {
	// success
	db := DB(t)
	assert.NotNil(t, db)
}

func TestGet(t *testing.T) {
	// setup
	db := DB(t)

	// success
	uuid := Get(t, db, "select uuid from Users where addr=?", "1.1.1.1")
	assert.Equal(t, "1111", uuid)
}

func TestRequest(t *testing.T) {
	// success
	r := Request("GET", "/", "Body.", "name=data")
	assert.Equal(t, "GET", r.Method)
	assert.Equal(t, "/", r.URL.Path)
	assert.Equal(t, "data", r.PathValue("name"))

	// confirm - body contents
	bytes, err := io.ReadAll(r.Body)
	assert.Equal(t, "Body.", string(bytes))
	assert.NoError(t, err)
}

func TestResponse(t *testing.T) {
	// setup
	w := httptest.NewRecorder()
	fmt.Fprint(w, `{"name": "data"}`)

	// success
	code, data, err := Response(w)
	assert.Equal(t, http.StatusOK, code)
	assert.Equal(t, map[string]any{"name": "data"}, data)
	assert.NoError(t, err)
}

func TestSet(t *testing.T) {
	// setup
	db := DB(t)

	// success
	Set(t, db, "insert into Users (uuid, addr) values (?, ?)", "2222", "2.2.2.2")

	// confirm - query executed
	uuid := Get(t, db, "select uuid from Users where addr=?", "2.2.2.2")
	assert.Equal(t, "2222", uuid)
}
