package graph

import (
	"context"

	"github.com/aavsss/programado/graph/model"
	"github.com/aavsss/programado/graph/repository"
)

// Periods is the resolver for the periods field.
func (r *queryResolver) Periods(ctx context.Context) ([]*model.Period, error) {
	var periods []*model.Period

	for _, period := range repository.MockData {
		period.User = &repository.MockUsers[0]
		periods = append(periods, &period)
	}

	return periods, nil
}
