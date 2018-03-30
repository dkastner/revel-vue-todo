package tests

import (
	"github.com/revel/revel/testing"
)

type IndexTest struct {
	testing.TestSuite
}

func (t *IndexTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}
