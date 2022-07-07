package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/aavsss/programado/graph/generated"
	"github.com/aavsss/programado/graph/model"
	"github.com/aavsss/programado/graph/repository"
)

// CreateNewRequest is the resolver for the createNewRequest field.
func (r *mutationResolver) CreateNewRequest(ctx context.Context, input model.NewRequest) (*model.Request, error) {
	request := model.Request{
		StartTime:   input.StartTime,
		EndTime:     input.StartTime,
		Description: input.Description,
		Requester:   &repository.MockUsers[0],
		Requestee:   &repository.MockUsers[1],
	}
	repository.ScheduleQueue = append(repository.ScheduleQueue, request)
	return &request, nil
}

// Periods is the resolver for the periods field.
func (r *queryResolver) Periods(ctx context.Context) ([]*model.Period, error) {
	var periods []*model.Period

	for _, period := range repository.MockData {
		period.User = &repository.MockUsers[0]
		periods = append(periods, &period)
	}

	return periods, nil
}

// Requests is the resolver for the requests field.
func (r *queryResolver) Requests(ctx context.Context) ([]*model.Request, error) {
	var requests []*model.Request

	for _, request := range repository.ScheduleQueue {
		requests = append(requests, &request)
	}

	return requests, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
