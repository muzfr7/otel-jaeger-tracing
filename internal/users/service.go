package users

import (
	"context"
	"fmt"

	"github.com/muzfr7/otel-jaeger-tracing/internal/pkg/storage"
	"github.com/muzfr7/otel-jaeger-tracing/internal/pkg/trace"
)

type service struct {
	storage storage.UserStorer
}

func (s service) create(ctx context.Context, req *createRequest) error {
	// Create a child span.
	ctx, span := trace.NewSpan(ctx, "service.create", nil)
	defer span.End()

	if err := s.storage.Insert(ctx, storage.User{Name: req.Name}); err != nil {
		return fmt.Errorf("create: unable to store: %w", err)
	}

	return nil
}
