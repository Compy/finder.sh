package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/compy/finder.sh/pkg/datasources"
	"github.com/compy/finder.sh/pkg/tasks"
	"github.com/hibiken/asynq"
)

// ProcessTask is where the indexing logic lives.
// Each data source type in the system has to implement this function. This function will be called once
// when indexing is requested by the user or by the system as part of an automatic reindex.
func (d *JiraDataSource) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var pl datasources.DatasourceIndexPayload
	if err := json.Unmarshal(t.Payload(), &pl); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	dbDatasource, err := d.ORM.DataSource.Get(ctx, pl.ID)
	if err != nil {
		return fmt.Errorf("could not fetch datasource %d from database: %v %w", pl.ID, err, asynq.SkipRetry)
	}

	dbDatasource.Update().
		SetStatus("idle").
		SetLastIndexed(time.Now()).
		Exec(ctx)

	// List all projects
	// Fetch all issues in each project
	// Fetch all comments in each issue

	log.Printf("executing task: %s - %s", t.Type(), pl.Config)
	l := tasks.ExampleTaskPayload{Counter: 12}
	d.Tasks.New("example_task").Payload(l).Save()
	return nil
}
