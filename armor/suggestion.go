package armor

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"regexp"
)

// BlockFieldSuggestionPresenter returns a new error presenter that masks field suggestions
func BlockFieldSuggestionPresenter() func(ctx context.Context, e error) *gqlerror.Error {
	return func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)
		// return the error
		var gqlErr *gqlerror.Error
		if errors.As(err, &gqlErr) {
			// replace error message string like
			//     error.message = error.message.replace(/Did you mean ".+"/g, mask);
			s := regexp.MustCompile(`Did you mean ".+"`).ReplaceAllString(gqlErr.Message, "Did you mean \"***\"")
			gqlErr.Message = s
			return gqlErr
		}
		return gqlerror.WrapPath(graphql.GetPath(ctx), err)
	}
}
