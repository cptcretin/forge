package tagGenerator

import (
	"testing"

	"app"
)

const (
	bytes = 9
	count = 20
)

func TestKeyGenerate(t *testing.T) {
	for i := 0; i < count; i++ {
		t.Logf("%s\n", app.NewHandle(bytes))
	}
}
