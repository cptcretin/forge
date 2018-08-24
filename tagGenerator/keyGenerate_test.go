package tagGenerator

import (
	"testing"

	"app"
)

const (
	bytes = 22
	count = 12
)

func TestKeyGenerate(t *testing.T) {
	for i := 0; i < count; i++ {
		t.Logf("%s\n", app.NewHandle(bytes))
	}
}
