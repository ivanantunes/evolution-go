package whatsmeow_service

import "go.mau.fi/whatsmeow/store"

func applyProtocolVersion(version clientVersion) {
	store.SetWAVersion(store.WAVersionContainer{
		uint32(version.Major), uint32(version.Minor), uint32(version.Patch),
	})
}
