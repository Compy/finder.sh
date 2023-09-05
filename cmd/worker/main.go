package main

import (
	"fmt"
	"log"

	"github.com/compy/finder.sh/config"
	"github.com/compy/finder.sh/pkg/datasources/jira"
	"github.com/compy/finder.sh/pkg/tasks"
	"github.com/hibiken/asynq"
)

func main() {
	// Load the configuration
	cfg, err := config.GetConfig()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}

	// Build the worker server
	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     fmt.Sprintf("%s:%d", cfg.Cache.Hostname, cfg.Cache.Port),
			DB:       cfg.Cache.Database,
			Password: cfg.Cache.Password,
		},
		asynq.Config{
			// See asynq.Config for all available options and explanation
			Concurrency: 10,
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		},
	)

	// Map task types to the handlers
	mux := asynq.NewServeMux()
	mux.Handle(tasks.TypeExample, new(tasks.ExampleProcessor))
	mux.Handle("index_jira", jira.NewJiraDatasource())

	// Start the worker server
	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run worker server: %v", err)
	}
}
