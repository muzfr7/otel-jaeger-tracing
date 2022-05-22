package users

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/muzfr7/otel-jaeger-tracing/internal/pkg/trace"
)

type createRequest struct {
	Name string `json:"name"`
}

func (c *createRequest) validate(ctx context.Context, body io.Reader) error {
	// Create a child span.
	_, span := trace.NewSpan(ctx, "createRequest.validate", nil)
	defer span.End()

	if err := json.NewDecoder(body).Decode(c); err != nil {
		return fmt.Errorf("validate: malformed body")
	}

	if c.Name == "" {
		return fmt.Errorf("validate: invalid request")
	}

	return nil
}
