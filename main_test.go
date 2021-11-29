package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_helloWorld(t *testing.T) {
	assert := assert.New(t)
	expected := "Hello World"
	assert.Equal(expected, helloWorld(), "Failed test HelloWorld")
}
