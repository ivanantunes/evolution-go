package whatsmeow_service

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

func TestHasPairedSession(t *testing.T) {
	for _, paired := range []bool{true, false} {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatal(err)
		}
		service := whatsmeowService{authDB: db}
		mock.ExpectQuery("SELECT EXISTS").WithArgs("test:1@s.whatsapp.net").WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(paired))
		result, err := service.HasPairedSession("test:1@s.whatsapp.net")
		if err != nil || result != paired {
			t.Fatalf("paired=%v, err=%v; want %v", result, err, paired)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Fatal(err)
		}
		db.Close()
	}
}

func TestHasPairedSessionWithoutJid(t *testing.T) {
	result, err := (whatsmeowService{}).HasPairedSession("")
	if result || err != nil {
		t.Fatalf("empty JID must be unpaired, got %v %v", result, err)
	}
}

func TestHasPairedSessionDatabaseUnavailable(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	service := whatsmeowService{authDB: db}
	mock.ExpectQuery("SELECT EXISTS").WithArgs("test:1@s.whatsapp.net").WillReturnError(errors.New("database unavailable"))
	result, err := service.HasPairedSession("test:1@s.whatsapp.net")
	if result || err == nil {
		t.Fatal("database failure must not imply a paired session")
	}
}
