package tagGenerator

import (
    "testing"

    "forge/1.0/app"
)

const (
    bytes = 7
    count = 12
)

func TestKeyGenerate(t *testing.T) {
    for i := 0; i < count; i++ {
        t.Logf("%s\n", app.NewHandle(bytes))
    }
}
