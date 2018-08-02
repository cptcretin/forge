package test

import (
    "testing"
    "time"

    "domain"
)

func TestContext(t *testing.T) {
    c := domain.Create("Test Context")
    defer c.End()

    c.Start("This is a solo transaction")
    tx := c.Start("This is a parent transaction")

    tx.Start("This transaction has a parent")
    txx := tx.Startf("This transaction happened at %v", time.Now())
    txx.Start("Even deeper")
    tx.Start("Transaction has ended")

    <- time.After(time.Second * 5)

    c.Start("This transaction happened after a long wait.")
    c.Start("Complete 😊")
}