package tagGenerator

import (
    "testing"

    "forge/1.0/app"
)

const bytes = 7

func TestKeyGenerate(t *testing.T) {
    t.Logf("%s\n", app.NewHandle(bytes))
}
