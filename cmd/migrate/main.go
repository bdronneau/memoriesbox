package main

import (
	"context"
	"flag"
	"log"
	"os"

	dbModels "github.com/bdronneau/memoriesbox/pkg/db/models"
	"github.com/bdronneau/memoriesbox/pkg/logger"
	"github.com/rs/xid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/bdronneau/memoriesbox/pkg/db"

	"github.com/peterbourgon/ff/v3"
)

func main() {
	fs := flag.NewFlagSet("memoriesbox", flag.ContinueOnError)

	dbConfig := db.GetConfig(fs)

	err := ff.Parse(fs, os.Args[1:],
		ff.WithEnvVarPrefix("MEMORIESBOX"),
	)
	if err != nil {
		log.Fatal(err)
	}

	loggerApp := logger.New(fs)
	dbApp, err := db.New(dbConfig, loggerApp)
	if err != nil {
		log.Fatal(err)
	}

	loggerApp.Sugar.Debug("Retrieve all memories")

	memories, err := dbModels.Memories(qm.OrderBy("RANDOM()")).All(context.Background(), dbApp.DB)
	if err != nil {
		loggerApp.Sugar.Fatal("Can not retrieve all memories")
	}

	for _, memory := range memories {
		guid := xid.New()
		memory.Xid = null.StringFrom(guid.String())
		memory.Update(context.Background(), dbApp.DB, boil.Infer())
		if err != nil {
			loggerApp.Sugar.Fatalf("Can not update row %v", err)
		}
	}
}
