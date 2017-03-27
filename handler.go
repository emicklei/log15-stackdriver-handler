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
	// If context keys are  nil or empty (string) then this setting
	// tells whether the keys must be excluded
	OmitEmpty bool
}

// NewHandler returns a new log15 handler that sends logging to GCP.
func NewHandler(projectID, logName string) (*StackdriverHandler, error) {
	client, err := logging.NewClient(context.Background(), projectID)
	if err != nil {
		return nil, fmt.Errorf("Failed to create client: %v", err)
	}
	return &StackdriverHandler{
		client:    client,
		logger:    client.Logger(logName),
		OmitEmpty: true,
	}, nil
}

// Log implements log15.Handler
func (h *StackdriverHandler) Log(r *log15.Record) error {
	e := logging.Entry{
		Payload:   r.Msg,
		Timestamp: r.Time,
		Severity:  asSeverity(r.Lvl),
	}
	if r.Ctx != nil && len(r.Ctx) > 0 {
		labels := map[string]string{}
		mapContextToFields(r.Ctx, h.OmitEmpty, labels)
		e.Labels = labels
	}
	h.logger.Log(e)
	return nil
}

func asSeverity(l log15.Lvl) logging.Severity {
	switch l {
	case log15.LvlDebug:
		return logging.Debug
	case log15.LvlCrit:
		return logging.Emergency
	case log15.LvlError:
		return logging.Error
	case log15.LvlInfo:
		return logging.Info
	case log15.LvlWarn:
		return logging.Warning
	default:
		return logging.Notice
	}
}

// Close is delegated to the client (if any)
func (h *StackdriverHandler) Close() error {
	if h.client == nil {
		return nil
	}
	return h.client.Close()
}
