package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"torre-test-gql/graph/generated"
	"torre-test-gql/graph/model"
	"torre-test-gql/pkg/jobs"
)

func (r *queryResolver) JobStatus(ctx context.Context) ([]*model.JobStatus, error) {
	query, err := jobs.CountPerStatus()
	if err != nil {
		return nil, err
	}
	return query, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
