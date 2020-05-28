package resolvers

import (
	"context"

	"github.com/daesu/stoicadon/api/graphql/gen"
	"github.com/daesu/stoicadon/models"
	"github.com/daesu/stoicadon/services"
)

type Resolver struct{}

func (r *queryResolver) Health(ctx context.Context) (*models.Health, error) {
	return services.GetHealth(ctx)
}

// Query returns gen.QueryResolver implementation.
func (r *Resolver) Query() gen.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// NewResolver ...
func NewResolver() *Resolver {
	return &Resolver{}
}
