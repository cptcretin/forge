package tagGenerator

import (
    "fmt"
    "testing"

    "forge/1.0/app"
)

const bytes = 11

func TestKeyGenerate(t *testing.T) {
    fmt.Printf("%s\n", app.NewHandle(bytes))
}