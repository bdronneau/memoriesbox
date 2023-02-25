package repositories

import (
	dbMemories "memoriesbox/pkg/db"
	"memoriesbox/pkg/logger"
	"memoriesbox/pkg/repositories/models"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"go.uber.org/zap/zaptest"
)

func TestGetRandomMemories(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	// Set expectations for the mock database
	rows := sqlmock.NewRows([]string{"id", "author", "content", "append"}).AddRow(1, "John", "Doe is not my lastname", time.Date(2022, 12, 12, 0, 0, 0, 0, time.UTC))
	mock.ExpectQuery(`SELECT "mbox"\."memories".* FROM "mbox"\."memories" ORDER BY RANDOM\(\) LIMIT 1;`).WillReturnRows(rows)

	loggerApp := logger.App{Sugar: zaptest.NewLogger(t).Sugar()}
	dbApp := dbMemories.App{DB: db}
	repoApp := New(Config{}, loggerApp, dbApp)

	result, err := repoApp.GetRandomMemories()
	if err != nil {
		t.Errorf("GetRandomMemories: %s", err)
	}

	expected := models.Memory{
		ID:      1,
		Author:  "John",
		Content: "Doe is not my lastname",
		Append:  "2022-12-12",
	}

	if result != expected {
		t.Errorf("Expected %v but got %v", expected, result)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
