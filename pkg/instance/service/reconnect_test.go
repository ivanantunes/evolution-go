package instance_service

import (
	"errors"
	"testing"

	"github.com/evolution-foundation/evolution-go/pkg/config"
	instance_model "github.com/evolution-foundation/evolution-go/pkg/instance/model"
	logger_wrapper "github.com/evolution-foundation/evolution-go/pkg/logger"
	whatsmeow_service "github.com/evolution-foundation/evolution-go/pkg/whatsmeow/service"
	"go.mau.fi/whatsmeow"
)

type reconnectStub struct {
	whatsmeow_service.WhatsmeowService
	called string
	err    error
}

func (s *reconnectStub) ReconnectClient(id string) error {
	s.called = id
	return s.err
}

func TestReconnectDisconnectedInstance(t *testing.T) {
	wantErr := errors.New("reconnect failed")
	for _, reconnectErr := range []error{nil, wantErr} {
		stub := &reconnectStub{err: reconnectErr}
		service := instances{
			whatsmeowService: stub,
			clientPointer:    map[string]*whatsmeow.Client{"disconnected-instance": {}},
			loggerWrapper:    logger_wrapper.NewLoggerManager(&config.Config{LogDirectory: t.TempDir()}),
		}
		err := service.Reconnect(&instance_model.Instance{Id: "disconnected-instance"})
		if stub.called != "disconnected-instance" || !errors.Is(err, reconnectErr) {
			t.Fatalf("reconnect called=%q error=%v, want error=%v", stub.called, err, reconnectErr)
		}
	}
}
