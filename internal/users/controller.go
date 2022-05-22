package users

import (
	"log"
	"net/http"

	"github.com/muzfr7/otel-jaeger-tracing/internal/pkg/storage"
	"github.com/muzfr7/otel-jaeger-tracing/internal/pkg/trace"
)

type Controller struct {
	service service
}

func New(storage storage.UserStorer) Controller {
	return Controller{
		service: service{
			storage: storage,
		},
	}
}

func (c Controller) Create(w http.ResponseWriter, r *http.Request) {
	// You could actually use `trace.SpanFromContext` here instead!

	// Create the parent span.
	ctx, span := trace.NewSpan(r.Context(), "Controller.Create", nil)
	defer span.End()

	// Some random informative tags.
	trace.AddSpanTags(span, map[string]string{"app.tag_1": "val_1", "app.tag_2": "val_2"})

	// Some random informative event.
	trace.AddSpanEvents(span, "test", map[string]string{"event_1": "val_1", "event_2": "val_2"})

	req := &createRequest{}
	if err := req.validate(ctx, r.Body); err != nil {
		// Logging error on span but not marking it as "failed".
		trace.AddSpanError(span, err)

		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := c.service.create(ctx, req); err != nil {
		// Logging error on span and marking it as "failed".
		trace.AddSpanError(span, err)
		trace.FailSpan(span, "internal error")

		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
