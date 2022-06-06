package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	expected := "Hello pengu!"
	got := SayHelloTo("pengu")

	assert.Equal(t, expected, got)
}
