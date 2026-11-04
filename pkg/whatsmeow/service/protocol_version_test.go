package whatsmeow_service

import (
	"testing"

	"go.mau.fi/whatsmeow/store"
)

func TestSelectedVersionIsUsedForHandshake(t *testing.T) {
	previous := store.GetWAVersion()
	t.Cleanup(func() { store.SetWAVersion(previous) })
	selected := clientVersion{Major: 2, Minor: 3000, Patch: 1053049953}
	applyProtocolVersion(selected)
	want := store.WAVersionContainer{2, 3000, 1053049953}
	if got := store.GetWAVersion(); got != want {
		t.Fatalf("handshake version=%v, want selected version=%v", got, want)
	}
	if got := store.BaseClientPayload.GetUserAgent().GetAppVersion(); got.GetTertiary() != want[2] {
		t.Fatalf("login payload version=%v, want %v", got, want)
	}
}
