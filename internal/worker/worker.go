package worker

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/rs/zerolog/log"
)

func Work(ctx context.Context, id int, err bool) error {
	log.Info().Int("worker", id).Msg("Worker starting")

	// 30% chance of error
	if err && rand.Float32() < 0.3 {
		err := fmt.Errorf("worker %d failed", id)
		log.Err(err).Msg("worker did not complete work")
		return err
	}

	// Sleep for a random duration less than 1000ms
	sleepDuration := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(sleepDuration)
	log.Info().Int("worker", id).Dur("slept_for", sleepDuration).Msg("Worker finished sleeping")

	log.Info().Int("worker", id).Msg("Worker completed successfully")
	return nil
}
