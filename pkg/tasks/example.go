package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

type ExampleTaskPayload struct {
	Counter int
}

// TypeExample is the type for the example task.
// This is what is passed in to TaskClient.New() when creating a new task
const TypeExample = "example_task"

// ExampleProcessor processes example tasks
type ExampleProcessor struct {
}

// ProcessTask handles the processing of the task
func (p *ExampleProcessor) ProcessTask(ctx context.Context, t *asynq.Task) error {

	var pl ExampleTaskPayload
	if err := json.Unmarshal(t.Payload(), &pl); err != nil {
		log.Printf("Fuck: %v", t.Payload())
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	log.Printf("executing task: %s - %d", t.Type(), pl.Counter)
	time.Sleep(250 * time.Millisecond)
	return nil
}
