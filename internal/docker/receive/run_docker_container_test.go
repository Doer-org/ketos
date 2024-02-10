package docker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunContainer(t *testing.T){
	respId, err := CreateContainer()
	assert.NotEmpty(t, respId)
	t.Logf("respId: %s", respId)
	assert.NoError(t, err)
	err = RunContainer(respId)
	assert.NoError(t, err)
}