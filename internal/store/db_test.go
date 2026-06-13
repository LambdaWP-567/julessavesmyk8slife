package store

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestNewCluster(t *testing.T) {
    c := NewCluster(1, "test-cluster")
    assert.Equal(t, 1, c.ID)
    assert.Equal(t, "test-cluster", c.Name)
}
