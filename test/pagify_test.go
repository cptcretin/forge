package test

import (
	"testing"

	"github.com/cptcretin/forge/app"
	"github.com/cptcretin/forge/context"
)

func TestPagify(t *testing.T) {
	c := context.New("Testing Pagify")

	defer c.End()

	l := []string{
		"Line #01",
		"Line #02",
		"Line #03",
		"Line #04",
		"Line #05",
		"Line #06",
		"Line #07",
		"Line #08",
		"Line #09",
		"Line #10",
		"Line #11",
		"Line #12",
		"Line #13",
		"Line #14",
		"Line #15",
		"Line #16",
		"Line #17",
		"Line #18",
		"Line #19",
		"Line #20",
		"Line #21",
		"Line #22",
		"Line #23",
		"Line #24",
		"Line #25",
		"Line #26",
		"Line #27",
		"Line #28",
		"Line #29",
		"Line #30",
		"Line #31",
		"Line #32",
		"Line #33",
		"Line #34",
		"Line #35",
		"Line #36",
		"Line #37",
		"Line #38",
		"Line #39",
		"Line #40",
		"Line #41",
		"Line #42",
		"Line #43",
		"Line #44",
		"Line #45",
		"Line #46",
		"Line #47",
		"Line #48",
		"Line #49",
		"Line #50",
		"Line #51",
		"Line #52",
	}

	job := func(i interface{}, tx *context.Tx) error {
		t.Logf("read string: %s", i)
		return nil
	}

	c.Pagify(app.ToInterfaces(l), job)
}
