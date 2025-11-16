package neat

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTrim(t *testing.T) {
	// success - under length
	text := trim("text", 5)
	assert.Equal(t, "text", text)

	// success - at length
	text = trim("text", 4)
	assert.Equal(t, "text", text)

	// success - over length
	text = trim("text", 3)
	assert.Equal(t, "tex", text)
}

func TestBody(t *testing.T) {
	// success
	body := Body("\tBody.\n", 5)
	assert.Equal(t, "Body.", body)
}

func TestName(t *testing.T) {
	// success
	name := Name("\tNAME\n", 4)
	assert.Equal(t, "name", name)
}

func TestPath(t *testing.T) {
	// success
	path := Path("\t/././path")
	assert.Equal(t, "/path", path)
}

func TestTime(t *testing.T) {
	// setup
	want := time.Unix(1234567890, 0).Local()

	// success
	tobj := Time(1234567890)
	assert.Equal(t, want, tobj)
}
