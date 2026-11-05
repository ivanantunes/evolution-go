package whatsmeow_service

import (
	"reflect"
	"testing"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store"
)

// A missing status handler leaves offline statuses unacknowledged: WhatsApp
// blocks later delivery and closes the stream with ack class=status errors.
func TestWhatsmeowHandlesStatusStanzas(t *testing.T) {
	client := whatsmeow.NewClient(&store.Device{}, nil)
	handlers := reflect.ValueOf(client).Elem().FieldByName("nodeHandlers")
	if !handlers.IsValid() || !handlers.MapIndex(reflect.ValueOf("status")).IsValid() {
		t.Fatal("Whatsmeow must handle status stanzas to acknowledge the offline queue")
	}
}
