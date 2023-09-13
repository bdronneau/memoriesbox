package repositories

import (
	"context"
	"testing"
	"time"

	dbMemories "github.com/bdronneau/memoriesbox/pkg/db"
	"github.com/bdronneau/memoriesbox/pkg/logger"
	"github.com/bdronneau/memoriesbox/pkg/repositories/models"
	"github.com/bdronneau/memoriesbox/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

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

type Suite struct {
	suite.Suite

	integration *test.PostgresIntegration
	repoApp     App
}

func (r *Suite) SetupSuite() {
	r.integration = test.NewIntegration(r.T())
	r.integration.Bootstrap("memoriesbox_test")

	dbApp := dbMemories.App{DB: r.integration.DB()}
	loggerApp := logger.App{Sugar: zaptest.NewLogger(r.T()).Sugar()}
	r.repoApp = New(Config{}, loggerApp, dbApp)
}

func (r *Suite) TearDownSuite() {
	r.integration.Close()
}

func (r *Suite) TearDownTest() {
	r.integration.Reset(context.Background())
}

func TestMemoriesSuite(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	t.Parallel()

	suite.Run(t, new(Suite))
}

func (s *Suite) TestCreateAndRetrieve() {
	s.Run("Count", func() {
		count, err := s.repoApp.CountMemories()
		assert.NoError(s.T(), err)
		assert.Equal(s.T(), int64(0), count, "Expected to have a no memory in DB")
	})

	s.Run("complete workflow", func() {
		s.repoApp.AddMemory("Doe is not my lastname", "John", time.Date(2022, 12, 12, 0, 0, 0, 0, time.UTC))

		count, err := s.repoApp.CountMemories()
		assert.NoError(s.T(), err)
		assert.Equal(s.T(), int64(1), count, "Expected to have a new memory in DB")

		expected := models.Memory{
			ID:      1,
			Author:  "John",
			Content: "Doe is not my lastname",
			Append:  "2022-12-12",
		}

		actual, err := s.repoApp.GetRandomMemories()
		assert.NoError(s.T(), err)

		assert.Equal(s.T(), expected, actual, "Memory has not been insert has expected")
	})
}
