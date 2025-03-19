package wgex

import (
	"context"
	"sync"

	"github.com/thenakulchawla/go-fundamentals/internal/worker"
	"github.com/thenakulchawla/parchment"
)

const NUM_THREADS = 5

func RunExamples(ctx context.Context) error {

	ctx = parchment.AddToLogger(ctx, []parchment.LoggerField{
		{Key: "program", Value: "wait_groups"},
	})

	// baseWG(ctx)
	wgWithErrors(ctx)
	return nil
}

func baseWG(ctx context.Context) {

	var wg sync.WaitGroup

	for i := range NUM_THREADS {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker.Work(ctx, id, false)
		}(i)
	}

	log := parchment.FromContext(ctx)
	log.Info().Msg("Waiting for all workers to complete...")
	wg.Wait()

}

func wgWithErrors(ctx context.Context) {
	var wg sync.WaitGroup
	errors := make([]error, NUM_THREADS)

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			if err := worker.Work(ctx, id, true); err != nil {
				errors[i] = err
			}
		}(i)
	}

	log := parchment.FromContext(ctx)
	log.Info().Msg("Waiting for all workers to complete...")
	wg.Wait()

	if len(errors) > 0 {
		log.Error().Msgf("Encountered %d errors:", len(errors))
		for _, err := range errors {
			log.Error().Msg(err.Error())
		}
	} else {
		log.Info().Msg("All workers completed successfully")
	}
}
