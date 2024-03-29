package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/kilianstallz/gql-armor/testutils/testserver/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	return &model.Todo{
		ID:   fmt.Sprintf("T%d", rand.Int()),
		Done: false,
		Text: input.Text,
		User: &model.User{
			ID:   input.UserID,
			Name: "A user",
		},
	}, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	// stub implementation
	return []*model.Todo{
		{
			ID:   "T1",
			Done: false,
			Text: "Todo 1",
			User: &model.User{
				ID:   "U1",
				Name: "A user",
			},
		},
	}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
