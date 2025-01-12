package channels

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func RunExamples() error {
	log := log.With().Str("program", "channels").Logger()
	basicChannelBuffered(log)
	return nil
}

func basicChannelBuffered(log zerolog.Logger) {

	log.Info().Msg("basic channels")

	// channels are like queues that you can add data to and receive data from
	// this is a buffered channel, values can be added to it without receiving
	queue := make(chan int, 3)
	queue <- 1
	queue <- 2

	log.Info().Int("q_len", len(queue)).Msg("length of channel")
	var x int
	for len(queue) > 0 {
		x = <-queue
		log.Info().Int("popped_val", x).Msg("popped value from queue")
	}

}

func basicChannelUnbuffered(log zerolog.Logger) {

}
