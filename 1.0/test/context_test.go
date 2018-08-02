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

func TestContextFormat(t *testing.T) {
	f := struct {
		I int
		S string
	} { 10, "Twenty" }

	c1 := context.Createf("This is a formatted context: (%s)", "my special formatting")
	c1.StartTransaction("one line")
	t.Log(c1.String())

	c2 := context.Createf("This is a formatted context: (%v)", f)
	c2.StartTransaction("two line")
	t.Log(c2.String())

	c3 := context.Get("Can I Break In?")
	c3.StartTransaction("three line")
	t.Log(c3.String())
}

func TestContextLogging(t *testing.T) {
	d := struct {
		I int
		S string
	} { 33, "In the morning!" }

	context.Logf(context.Trace, "This is my first message: %s", "Hi!")
	context.Log(context.Trace, d)
	context.Logf(context.Trace, "message: %v", d)
	context.Logf(context.Trace, "Multiple strings: %v, %s, %s, %s", d, "one", "two", "three")
}