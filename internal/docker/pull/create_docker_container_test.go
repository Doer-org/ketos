package docker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateContainer(t *testing.T){
	resp_id, err := CreateContainer()
	assert.NotEmpty(t, resp_id)
	t.Logf("resp_id: %s", resp_id)
	assert.NoError(t, err)
}