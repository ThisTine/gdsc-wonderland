package mng

import (
	"context"
	"time"

	"github.com/kamva/mgm/v3"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo/options"

	mod "backend/modules"
)

func Init() {
	// Setup command monitor
	monitor := &event.CommandMonitor{
		Started: (func() func(_ context.Context, evt *event.CommandStartedEvent) {
			if mod.Conf.LogLevel == 6 {
				return func(_ context.Context, evt *event.CommandStartedEvent) {
					logrus.WithField("command", evt.Command.String()).Trace("MONGO COMMAND STARTED")
				}
			}
			return nil
		})(),
	}

	// Initialize default MGM configuration
	if err := mgm.SetDefaultConfig(
		&mgm.Config{
			CtxTimeout: 5 * time.Second,
		},
		mod.Conf.MongoName,
		options.Client().ApplyURI(mod.Conf.MongoUrl).SetMonitor(monitor),
	); err != nil {
		logrus.WithError(err).Fatal("UNABLE TO CONFIGURATION MGM")
	}

	// Load MongoDB connection and database
	_, client, database, err := mgm.DefaultConfigs()
	if err != nil {
		logrus.WithError(err).Fatal("UNABLE TO LOAD MGM")
	}

	mod.Database = database
	mod.Client = client

	initCollection()
}
