package stack15

import (
	"fmt"

	"github.com/inconshreveable/log15"

	// Imports the Stackdriver Logging client package.
	"cloud.google.com/go/logging"
	"golang.org/x/net/context"
)

// StackdriverHandler is a log15.Handler
type StackdriverHandler struct {
	client *logging.Client
	logger *logging.Logger
}

// NewHandler returns a new log15 handler that sends logging to GCP.
func NewHandler(projectID, logName string) (log15.Handler, error) {
	client, err := logging.NewClient(context.Background(), projectID)
	if err != nil {
		return nil, fmt.Errorf("Failed to create client: %v", err)
	}
	return &StackdriverHandler{
		client: client,
		logger: client.Logger(logName),
	}, nil
}

// Log implements log15.Handler
func (h *StackdriverHandler) Log(r *log15.Record) error {
	// TODO copy other fields
	h.logger.Log(logging.Entry{Payload: r.Msg})
	return nil
}
