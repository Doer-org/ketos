package docker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecompressTarGzToImage(t *testing.T) {
	err := DecompressTarGzToImage()
	assert.NoError(t, err)
}