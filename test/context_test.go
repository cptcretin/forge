package test

import (
	"testing"
	"time"

	"github.com/cptcretin/forge/context"
)

func TestContext(t *testing.T) {
	c := context.New("Test Context: %s", "Test Param")
	defer c.End()

	c.Start("This is a solo transaction")
	tx := c.Start("This is a parent transaction")

	tx.Start("This transaction has a parent")
	txx := tx.Start("This transaction happened at %v", time.Now())

	txx.
		NewThread("This is a background transaction").
		Run(func(tx *context.Tx) {
			<-time.After(time.Second * 10)
			tx.Start("Background transaction has ended")
		})

	txx.Start("Even deeper")
	tx.Start("Transaction has ended")

	<-time.After(time.Second * 5)

	c.Start("This transaction happened after a long wait.")
	c.Start("Complete 😊")
}
