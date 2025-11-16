package test

import (
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

func TestSet(t *testing.T) {
	// setup
	db := DB(t)

	// success
	Set(t, db, "insert into Users (uuid, addr) values (?, ?)", "2222", "2.2.2.2")

	// confirm - query executed
	uuid := Get(t, db, "select uuid from Users where addr=?", "2.2.2.2")
	assert.Equal(t, "2222", uuid)
}
