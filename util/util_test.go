package util

import (
	"testing"

	"github.com/garciaolais/gownloader/util"
	"github.com/stretchr/testify/assert"
)

func TestIsURL(t *testing.T) {
	urls := []string{
		"http://www.google.com",
		"https://www.google.com",
	}

	for _, url := range urls {
		err := util.IsURL(url)
		assert.NoError(t, err, "Expected no error")
	}

	err := util.IsURL("")
	assert.Error(t, err, "Expected error")

	err = util.IsURL(":V")
	assert.Error(t, err, "Expected error")
}
