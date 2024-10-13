package wgex

import (
	"sync"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/thenakulchawla/go-fundamentals/internal/worker"
)

const NUM_THREADS = 5

func RunExamples() error {

	log := log.With().Str("program", "wait_groups").Logger()
	// baseWG(log)
	wgWithErrors(log)
	return nil
}

func baseWG(log zerolog.Logger) {

	var wg sync.WaitGroup

	for i := range NUM_THREADS {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker.Work(id, false)
		}(i)
	}

	log.Info().Msg("Waiting for all workers to complete...")
	wg.Wait()

}

func wgWithErrors(log zerolog.Logger) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	errors := make([]error, 0)

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			if err := worker.Work(id, true); err != nil {
				mu.Lock()
				errors = append(errors, err)
				mu.Unlock()
			}
		}(i)
	}

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
