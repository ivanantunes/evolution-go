package instance_service

import (
	"errors"
	"testing"

	"github.com/evolution-foundation/evolution-go/pkg/config"
	instance_model "github.com/evolution-foundation/evolution-go/pkg/instance/model"
	logger_wrapper "github.com/evolution-foundation/evolution-go/pkg/logger"
	whatsmeow_service "github.com/evolution-foundation/evolution-go/pkg/whatsmeow/service"
)

type qrStartStub struct {
	whatsmeow_service.WhatsmeowService
	starts int
}

func (s *qrStartStub) StartInstance(string) error {
	s.starts++
	return errors.New("unexpected start")
}

func TestQrPollingDoesNotStartAnotherClientDuringReconnect(t *testing.T) {
	stub := &qrStartStub{}
	service := instances{whatsmeowService: stub,
		loggerWrapper: logger_wrapper.NewLoggerManager(&config.Config{LogDirectory: t.TempDir()}),
	}
	_, err := service.GetQr(&instance_model.Instance{Id: "reconnecting"})
	if stub.starts != 0 {
		t.Fatalf("QR polling started %d competing clients", stub.starts)
	}
	if err == nil {
		t.Fatal("expected not-ready result while the client is being created")
	}
}
