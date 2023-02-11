package armor

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

const characterLimitExtension = "CharacterLimit"

type CharacterLimitOptions struct {
	// Limit is the maximum number of characters allowed in a query.
	Limit int
}

var DefaultCharacterLimit = &CharacterLimitOptions{
	Limit: 15500,
}

func FixedCharacterLimit(options *CharacterLimitOptions) *CharacterLimit {
	limit := DefaultCharacterLimit.Limit
	if options != nil {
		if options.Limit > 0 {
			limit = options.Limit
		}
	}
	return &CharacterLimit{
		limit: limit,
	}
}

// CharacterLimit allows you to define a limit on character usage in a query.
//
// If a query is submitted that exceeds the limit, a 422 status code will be returned.
type CharacterLimit struct {
	limit int
}

type CharacterStats struct {
	Characters     int
	CharacterLimit int
}

var _ interface {
	graphql.OperationParameterMutator
	graphql.HandlerExtension
} = &CharacterLimit{}

func (c *CharacterLimit) ExtensionName() string {
	return characterLimitExtension
}

func (c *CharacterLimit) Validate(schema graphql.ExecutableSchema) error {
	if c.limit == 0 {
		return fmt.Errorf("CharacterLimit limit can not be 0")
	}
	return nil
}

func (c CharacterLimit) MutateOperationParameters(ctx context.Context, request *graphql.RawParams) *gqlerror.Error {
	characters := len(request.Query)
	limit := c.limit

	if characters > limit {
		return gqlerror.Errorf("Syntax Error: Character limit of %d exceeded, found %d", limit, characters)
	}

	return nil
}
