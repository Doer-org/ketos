package docker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecompressTarGzToImage(t *testing.T) {
	tarGzFileName := "../../../testdata/test-tar/ketos-tmp-image.tar.gz"
	err := DecompressTarGzToImage(tarGzFileName)
	assert.NoError(t, err)
}
