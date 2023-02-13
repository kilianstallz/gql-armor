package armor_test

import (
	"context"
	"github.com/gavv/httpexpect/v2"
	"github.com/kilianstallz/gql-armor/armor"
	"github.com/kilianstallz/gql-armor/testutils"
	"testing"
)

func TestFixedAliasLimit1(t *testing.T) {
	// setup server
	gql := testutils.BuildGqlServer()
	gql.Use(armor.FixedAliasLimit(2))
	expect, td := testutils.Setup(t, testutils.SetupOption{
		Gql:      gql,
		Teardown: func(t *testing.T) {},
	})
	defer td()

	tests := []struct {
		name    string
		arrange func(t *testing.T)
		act     func(t *testing.T) *httpexpect.Response
		assert  func(t *testing.T, got *httpexpect.Response)
		args    struct {
			ctx context.Context
		}
		teardown func(t *testing.T)
	}{
		{
			name:    "It should return an error when the limit is exceeded",
			arrange: func(t *testing.T) {},
			assert: func(t *testing.T, got *httpexpect.Response) {
				// check code in extensions
				testutils.AssertCode(t, got, "ALIAS_LIMIT_EXCEEDED")
			},
			act: func(t *testing.T) *httpexpect.Response {
				query := `query {
						a: todos {id}
						b: todos {id}
						c: todos {id}
					}`
				return expect.POST("/").WithJSON(testutils.Query(query)).Expect()
			},
		},
		{
			name:    "It should return no error when the limit is equal",
			arrange: func(t *testing.T) {},
			assert: func(t *testing.T, got *httpexpect.Response) {
				got.JSON().Object().Value("data").Object().Value("a").Array().Length().IsEqual(1)
				got.JSON().Object().Value("data").Object().Value("b").Array().Length().IsEqual(1)
			},
			act: func(t *testing.T) *httpexpect.Response {
				query := `query {
						a: todos {id}
						b: todos {id}
					}`
				return expect.POST("/").WithJSON(testutils.Query(query)).Expect()
			},
		},
		{
			name:    "It should return no error when the limit is not exceeded",
			arrange: func(t *testing.T) {},
			assert: func(t *testing.T, got *httpexpect.Response) {
				got.JSON().Object().Value("data").Object().Value("a").Array().Length().IsEqual(1)
			},
			act: func(t *testing.T) *httpexpect.Response {
				query := `query {
						a: todos {id}
					}`
				return expect.POST("/").WithJSON(testutils.Query(query)).Expect()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.arrange(t)
			got := tt.act(t)
			tt.assert(t, got)
			if tt.teardown != nil {
				tt.teardown(t)
			}
		})
	}
}
