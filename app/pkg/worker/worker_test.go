package worker_test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/tombull/teamdream/app/pkg/worker"

	. "github.com/tombull/teamdream/app/pkg/assert"
)

var dummyTask = worker.Task{
	Name: "Do Something",
	Job: func(c *worker.Context) error {
		return nil
	},
}

func TestBackgroundWorker(t *testing.T) {
	RegisterT(t)

	var finished bool
	mu := &sync.RWMutex{}

	w := worker.New()
	w.Enqueue(worker.Task{
		Name: "Do Something",
		Job: func(ctx *worker.Context) error {
			mu.Lock()
			defer mu.Unlock()
			finished = true
			return nil
		},
	})

	Expect(w.Length()).Equals(int64(1))
	go w.Run("worker-1")
	Expect(func() bool {
		mu.RLock()
		defer mu.RUnlock()
		return finished
	}).EventuallyEquals(true)
}

func TestBackgroundWorker_ShutdownWhenEmpty(t *testing.T) {
	RegisterT(t)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	w := worker.New()
	Expect(w.Shutdown(ctx)).IsNil()
}

func TestBackgroundWorker_ShutdownWithStuckTasks(t *testing.T) {
	RegisterT(t)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	w := worker.New()
	w.Enqueue(dummyTask)
	Expect(w.Shutdown(ctx)).IsNotNil()
}

func TestBackgroundWorker_ShutdownWithRunningTasks(t *testing.T) {
	RegisterT(t)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	w := worker.New()
	w.Enqueue(dummyTask)
	go w.Run("worker-1")
	Expect(w.Shutdown(ctx)).IsNil()
}
