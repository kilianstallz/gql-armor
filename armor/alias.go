package armor

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

const errAliasLimit = "ALIAS_LIMIT_EXCEEDED"
const aliasExtension = "AliasLimit"

func FixedAliasLimit(limit int) *AliasLimit {
	return &AliasLimit{
		Func: func(ctx context.Context, rc *graphql.OperationContext) int {
			return limit
		},
	}
}

// AliasLimit allows you to define a limit on alias usage in a query.
//
// If a query is submitted that exceeds the limit, a 422 status code will be returned.
type AliasLimit struct {
	Func func(ctx context.Context, rc *graphql.OperationContext) int
	es   graphql.ExecutableSchema
}

type AliasStats struct {
	Alias      int
	AliasLimit int
}

var _ interface {
	graphql.OperationContextMutator
	graphql.HandlerExtension
} = &AliasLimit{}

func (a *AliasLimit) ExtensionName() string {
	return aliasExtension
}

func (a *AliasLimit) Validate(schema graphql.ExecutableSchema) error {
	if a.Func == nil {
		return fmt.Errorf("AliasLimit func can not be nil")
	}
	a.es = schema
	return nil
}

func (c AliasLimit) MutateOperationContext(ctx context.Context, rc *graphql.OperationContext) *gqlerror.Error {
	op := rc.Doc.Operations.ForName(rc.OperationName)
	aliases := countAliases(op.SelectionSet)
	limit := c.Func(ctx, rc)

	rc.Stats.SetExtension(aliasExtension, &AliasStats{
		Alias:      aliases,
		AliasLimit: limit,
	})

	if aliases > limit {
		err := gqlerror.Errorf("Syntax Error: Aliases limit of %d exceeded, found %d", limit, aliases)
		errcode.Set(err, errAliasLimit)
		return err
	}

	return nil
}

func countAliases(selection ast.SelectionSet) int {
	aliases := 0
	for _, selection := range selection {
		switch selection := selection.(type) {
		case *ast.Field:
			if selection.Alias != "" {
				aliases++
			}
		case *ast.InlineFragment:
			aliases += countAliases(selection.SelectionSet)
		case *ast.FragmentSpread:
			break
		}
	}
	return aliases
}

func GetAliasStats(ctx context.Context) *AliasStats {
	rc := graphql.GetOperationContext(ctx)
	if rc == nil {
		return nil
	}

	s, _ := rc.Stats.GetExtension(aliasExtension).(*AliasStats)
	return s
}
