package test

import (
	"errors"
	"testing"
	"time"

	"forge/1.0/context"
)

func contextClosureA(title string, t *testing.T) func() {
	c := context.Create(title)

	return func() {
		c.StartTransaction("Testing TX 1 of " + title)
		testOutsideClosure(t)
		c.StartTransaction("Test TX 2 of " + title)

		c.End()

		t.Log(c.String())
	}
}

func contextClosureB(title string, t *testing.T) func() {
	c := context.Create(title)

	return func() {
		c.StartTransaction("Testing TX 1 of " + title)
		c.End()

		t.Log(c.String())
	}
}

func testOutsideClosure(t *testing.T) {
	if context.Current == nil {
		t.Error("No closured context")
	}

	context.Current.CurrentTransaction().StartTransaction("Testing TX 1:1")
}

func TestContextCreation(t *testing.T) {
	contextClosureA("First Run", t)()
	contextClosureB("Second Run", t)()
}

func TestContextError(t *testing.T) {
	c := context.Create("Error Test")

	c.StartTransaction("Testing error transation")
	time.Sleep(10 * time.Second)

	c.Error(errors.New("This is a bogus error"))
}
