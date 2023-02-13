package testutils

import (
	"github.com/gavv/httpexpect/v2"
	"testing"
)

func AssertCode(t *testing.T, got *httpexpect.Response, code string) {
	t.Helper()
	got.JSON().Object().Value("errors").Array().Element(0).Object().Value("extensions").Object().Value("code").IsEqual(code)
}
