package egex

import (
	"context"
	"fmt"
	"sync"

	"github.com/thenakulchawla/go-fundamentals/internal/worker"
	"github.com/thenakulchawla/parchment"
	"go.uber.org/multierr"
	"golang.org/x/sync/errgroup"
)

const NUM_THREADS = 5

func RunExamples(ctx context.Context) error {
	ctx = parchment.AddToLogger(ctx, []parchment.LoggerField{
		{Key: "program", Value: "error_groups"},
	})
	log := parchment.FromContext(ctx)

	// log.Info().Msg("Running waitForAll example")
	// if err := waitForAll(log); err != nil {
	// 	log.Error().Err(err).Msg("waitForAll encountered errors")
	// }

	// log.Info().Msg("Running cancelOnFirstError example")
	// if err := showFirstError(log); err != nil {
	// 	log.Error().Err(err).Msg("cancelOnFirstError encountered an error")
	// }

	log.Info().Msg("Running combineErrors example")
	if err := collectAllErrorsMultiErr(ctx); err != nil {
		log.Error().Err(err).Msg("combineErrors encountered errors")
		// Unwrap and log individual errors
		for i, err := range multierr.Errors(err) {
			log.Error().Int("error_index", i).Err(err).Msg("Individual error")
		}
	}

	return nil
}

func waitForAll(ctx context.Context) error {
	g := new(errgroup.Group)

	for i := 0; i < NUM_THREADS; i++ {
		id := i
		g.Go(func() error {
			return worker.Work(ctx, id, true)
		})
	}
	log := parchment.FromContext(ctx)

	log.Info().Msg("Waiting for all workers to complete...")
	if err := g.Wait(); err != nil {
		return fmt.Errorf("one or more errors occurred: %w", err)
	}

	log.Info().Msg("All workers completed successfully")
	return nil
}

func showFirstError(ctx context.Context) error {
	log := parchment.FromContext(ctx)
	var g errgroup.Group

	for i := 0; i < NUM_THREADS; i++ {
		id := i
		g.Go(func() error {
			err := worker.Work(ctx, id, true)
			if err != nil {
				log.Error().Int("worker", id).Err(err).Msg("Worker encountered an error")
				return err // This will signal the errgroup to stop waiting and return this error
			}
			log.Info().Int("worker", id).Msg("Worker completed successfully")
			return nil
		})
	}

	log.Info().Msg("Waiting for workers (will return on first error)...")
	if err := g.Wait(); err != nil {
		return fmt.Errorf("operation stopped due to an error: %w", err)
	}

	log.Info().Msg("All workers completed successfully")
	return nil
}

/***
t1    t2    t3    t4    t5    Time
|     |     |     |     |
X                             Thread 1 (Error)
|-----|                       Thread 2
|-----------|                 Thread 3
|-------------------|         Thread 4
|-------------------------|   Thread 5
^
g.Wait() returns here with Thread 1's error
***/

func collectAllErrorsMultiErr(ctx context.Context) error {
	log := parchment.FromContext(ctx)
	var g errgroup.Group
	var combinedErr error
	var errMutex sync.Mutex

	for i := 0; i < NUM_THREADS; i++ {
		id := i
		g.Go(func() error {
			err := worker.Work(ctx, id, true)
			if err != nil {
				log.Error().Int("worker", id).Err(err).Msg("Worker encountered an error")
				errMutex.Lock()
				combinedErr = multierr.Append(combinedErr, fmt.Errorf("worker %d: %w", id, err))
				errMutex.Unlock()
			} else {
				log.Info().Int("worker", id).Msg("Worker completed successfully")
			}
			return nil // Always return nil to ensure g.Wait() doesn't return early
		})
	}

	log.Info().Msg("Waiting for all workers to complete...")
	g.Wait() // We don't check the error here as we're managing errors manually

	if combinedErr != nil {
		log.Error().Msg("One or more workers encountered errors")
		return combinedErr
	}

	log.Info().Msg("All workers completed successfully")
	return nil
}

func cancelOnFirstErrorWithContext(ctx context.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Ensure all resources are cleaned up

	g, ctx := errgroup.WithContext(ctx)
	log := parchment.FromContext(ctx)

	for i := 0; i < NUM_THREADS; i++ {
		id := i
		g.Go(func() error {
			select {
			case <-ctx.Done():
				log.Info().Int("worker", id).Msg("Worker cancelled due to error in another goroutine")
				return ctx.Err()
			default:
				err := worker.Work(ctx, id, true)
				if err != nil {
					log.Error().Int("worker", id).Err(err).Msg("Worker encountered an error")
					return err // This will cancel the context for other goroutines
				}
				log.Info().Int("worker", id).Msg("Worker completed successfully")
				return nil
			}
		})
	}

	log.Info().Msg("Waiting for workers (will cancel on first error)...")
	if err := g.Wait(); err != nil {
		return fmt.Errorf("operation cancelled due to an error: %w", err)
	}

	log.Info().Msg("All workers completed successfully")
	return nil
}
