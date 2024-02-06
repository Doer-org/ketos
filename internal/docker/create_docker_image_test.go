package docker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateImageWithDockerFile(t *testing.T) {
	// dockerfileを使ってイメージを作成
	err := createImageWithDockerFile("../../examples/go", "Dockerfile")
	assert.NoError(t, err)
}
