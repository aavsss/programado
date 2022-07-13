package graph

import (
	"context"

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

// Requests is the resolver for the requests field.
func (r *queryResolver) Requests(ctx context.Context) ([]*model.Request, error) {
	var requests []*model.Request

	for _, request := range repository.ScheduleQueue {
		requests = append(requests, &request)
	}

	return requests, nil
}
