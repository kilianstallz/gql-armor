package testutils

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gavv/httpexpect/v2"
	"github.com/kilianstallz/gql-armor/testutils/testserver/graph"
	"net/http"
	"net/http/httptest"
	"testing"
)

type query struct {
	Query string `json:"query"`
}

func Query(querybody string) query {
	// create json query
	return query{Query: querybody}
}

type SetupOption struct {
	Gql      *handler.Server
	Teardown func(t *testing.T)
}

func BuildGqlServer() *handler.Server {
	gqlSrv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	return gqlSrv
}

func Setup(t *testing.T, option SetupOption) (expect *httpexpect.Expect, teardown func()) {
	t.Helper()
	// Create a new http server
	http := http.NewServeMux()
	http.Handle("/", option.Gql)

	srv := httptest.NewServer(http)

	return httpexpect.WithConfig(httpexpect.Config{
			BaseURL:  srv.URL,
			Reporter: httpexpect.NewAssertReporter(t),
			Printers: []httpexpect.Printer{
				httpexpect.NewDebugPrinter(t, true),
			},
		}), func() {
			option.Teardown(t)
			defer srv.Close()
		}
}

// GetData gets data from graphql response.
func GetData(e *httpexpect.Response) *httpexpect.Value {
	return e.JSON().Path("$.data")
}

// GetObject return data from path.
// Path returns a new Value object for child object(s) matching given
// JSONPath expression.
// Example 1:
//
//	json := `{"users": [{"name": "john"}, {"name": "bob"}]}`
//	value := NewValue(t, json)
//
//	value.Path("$.users[0].name").String().Equal("john")
//	value.Path("$.users[1].name").String().Equal("bob")
func GetObject(obj *httpexpect.Object, path string) *httpexpect.Object {
	return obj.Path("$." + path).Object()
}

// GetErrors return errors from graphql response.
func GetErrors(e *httpexpect.Response) *httpexpect.Value {
	return e.JSON().Path("$.errors")
}
