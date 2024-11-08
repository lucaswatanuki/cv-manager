package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"
	"fmt"

	"github.com/lucaswatanuki/cv-manager/graph/model"
)

// CreatePerson is the resolver for the createPerson field.
func (r *mutationResolver) CreatePerson(ctx context.Context, input model.NewPerson) (*model.Person, error) {
	panic(fmt.Errorf("not implemented: CreatePerson - createPerson"))
}

// People is the resolver for the people field.
func (r *queryResolver) People(ctx context.Context) ([]*model.Person, error) {
	//mocking person for testing query
	person := &model.Person{
		ID: "id-01",
		Name: "Sasuke",
		Age: 12,
		IsMale: true,
	}

	people := []*model.Person{
		person,
	}

	return people, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
