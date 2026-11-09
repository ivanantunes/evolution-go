package instance_service

import (
	instance_model "github.com/evolution-foundation/evolution-go/pkg/instance/model"
	whatsmeow_service "github.com/evolution-foundation/evolution-go/pkg/whatsmeow/service"
	"testing"
)

type pairingService struct {
	whatsmeow_service.WhatsmeowService
	paired bool
}

func (s pairingService) HasPairedSession(jid string) (bool, error) { return s.paired, nil }

func TestStatusRetainsPairingWithoutConnectedClient(t *testing.T) {
	service := instances{whatsmeowService: pairingService{paired: true}}
	status, err := service.Status(&instance_model.Instance{Id: "instance", Jid: "test:1@s.whatsapp.net"})
	if err != nil {
		t.Fatal(err)
	}
	if !status.Paired || status.Connected || status.LoggedIn {
		t.Fatalf("offline retained session: %+v", status)
	}
}
func TestStatusWithoutPairing(t *testing.T) {
	service := instances{whatsmeowService: pairingService{paired: false}}
	status, err := service.Status(&instance_model.Instance{Id: "instance"})
	if err != nil {
		t.Fatal(err)
	}
	if status.Paired {
		t.Fatal("missing auth store must not report paired")
	}
}
