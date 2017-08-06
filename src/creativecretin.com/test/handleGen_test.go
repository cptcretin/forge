package test

import (
	//"fmt"
	"creativecretin.com/forge/1.0/app"
	"testing"
)

const (
	bytes    = 7
	keyCount = 10000
	testPage = 100
)

func TestTagGen(t *testing.T) {
	var keys [keyCount]string

	for i := 0; i < keyCount; i++ {
		keys[i] = app.NewHandle(bytes)
	}
}
