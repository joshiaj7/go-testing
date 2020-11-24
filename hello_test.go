package main

import (
	"github.com/stretchr/testify/assert"
	
	"testing"
)

// TestHello function
func TestHello(t *testing.T) {
  t.Run("empty", func(t *testing.T) {
    assert.Equal(t, hello(""), "Hello Dude!")
  })
  
  t.Run("ok", func(t *testing.T) {
    assert.Equal(t, hello("badut"), "Hello badut!")
  })
}